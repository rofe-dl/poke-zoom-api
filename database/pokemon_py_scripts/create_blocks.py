import io
import json
import argparse
from pathlib import Path

import cv2
import numpy as np
import requests
from PIL import Image

"""
AI-generated script to extract zoomed in blocks
of each Pokemon's image and storing the coordinates.
"""


def download_image(url: str) -> np.ndarray:
    """Download an image URL and return it as an RGBA numpy array."""
    response = requests.get(url, timeout=30)
    response.raise_for_status()

    image = Image.open(io.BytesIO(response.content)).convert("RGBA")
    return np.array(image)


def create_foreground_mask(image: np.ndarray) -> np.ndarray:
    """
    Create a binary foreground mask.

    For images with transparency, alpha is used directly.
    For images without transparency, everything is considered foreground.
    """
    alpha = image[:, :, 3]

    # Anything with some opacity is foreground.
    # You can change 10 to 1 if you want to be more permissive.
    foreground = alpha > 10

    return foreground.astype(np.uint8)


def create_silhouette_edges(foreground: np.ndarray) -> np.ndarray:
    """
    Find the outer boundary of the foreground object.
    """
    kernel = np.ones((3, 3), np.uint8)

    eroded = cv2.erode(foreground, kernel, iterations=1)

    # Pixels that disappear after erosion are boundary pixels.
    boundary = foreground - eroded

    return boundary


def create_image_edges(image: np.ndarray) -> np.ndarray:
    """
    Create a regular image edge map using Canny.
    """
    rgb = image[:, :, :3]

    gray = cv2.cvtColor(rgb, cv2.COLOR_RGB2GRAY)

    edges = cv2.Canny(
        gray,
        threshold1=50,
        threshold2=150,
    )

    return edges > 0


def calculate_patch_scores(
    image: np.ndarray,
    foreground: np.ndarray,
    silhouette_edges: np.ndarray,
    image_edges: np.ndarray,
    section_size: int,
    stride: int,
):
    """
    Generate and score all candidate patches.
    """

    height, width = image.shape[:2]

    results = []

    for y in range(0, height - section_size + 1, stride):
        for x in range(0, width - section_size + 1, stride):
            foreground_patch = foreground[
                y : y + section_size,
                x : x + section_size,
            ]

            silhouette_patch = silhouette_edges[
                y : y + section_size,
                x : x + section_size,
            ]

            image_edge_patch = image_edges[
                y : y + section_size,
                x : x + section_size,
            ]

            total_pixels = section_size * section_size

            foreground_density = np.count_nonzero(foreground_patch) / total_pixels

            silhouette_density = np.count_nonzero(silhouette_patch) / total_pixels

            image_edge_density = np.count_nonzero(image_edge_patch) / total_pixels

            # Weighted score.
            #
            # Silhouette edges are particularly valuable because they
            # allow patches containing some background.
            score = (
                0.45 * image_edge_density
                + 0.40 * silhouette_density
                + 0.15 * foreground_density
            )

            results.append(
                {
                    "x": x,
                    "y": y,
                    "size": section_size,
                    "score": round(float(score), 6),
                    "foreground_density": round(
                        float(foreground_density),
                        6,
                    ),
                    "silhouette_density": round(
                        float(silhouette_density),
                        6,
                    ),
                    "edge_density": round(
                        float(image_edge_density),
                        6,
                    ),
                }
            )

    return results


def filter_patches(
    patches,
    minimum_score: float,
):
    """
    Remove patches that contain essentially no useful information.

    No duplicate removal is performed.
    No diversity selection is performed.
    """

    return [patch for patch in patches if patch["score"] >= minimum_score]


def process_image(
    url: str,
    section_size: int,
    stride: int,
    minimum_score: float,
):
    print(f"Processing: {url}")

    image = download_image(url)

    height, width = image.shape[:2]

    print(f"  Image size: {width}x{height}")

    # Foreground / transparent-background mask
    foreground = create_foreground_mask(image)

    # Outer boundary of the Pokémon
    silhouette_edges = create_silhouette_edges(foreground)

    # Normal visual edges
    image_edges = create_image_edges(image)

    patches = calculate_patch_scores(
        image=image,
        foreground=foreground,
        silhouette_edges=silhouette_edges,
        image_edges=image_edges,
        section_size=section_size,
        stride=stride,
    )

    useful_patches = filter_patches(
        patches,
        minimum_score=minimum_score,
    )

    # Highest scoring first.
    useful_patches.sort(
        key=lambda patch: patch["score"],
        reverse=True,
    )

    print(f"  Candidate patches: {len(patches)}")
    print(f"  Useful patches:    {len(useful_patches)}")

    return {
        "url": url,
        "width": width,
        "height": height,
        "section_size": section_size,
        "stride": stride,
        "patches": useful_patches,
    }


def main():
    parser = argparse.ArgumentParser(description="Extract informative image sections.")

    parser.add_argument(
        "--urls",
        default="urls.json",
        help="JSON file containing a list of image URLs.",
    )

    parser.add_argument(
        "--size",
        type=int,
        required=True,
        help="Section width and height.",
    )

    parser.add_argument(
        "--stride",
        type=int,
        default=16,
        help="Distance between candidate sections. Default: 4.",
    )

    parser.add_argument(
        "--threshold",
        type=float,
        default=0.11,
        help="Minimum usefulness score. Default: 0.03.",
    )

    parser.add_argument(
        "--output",
        default="patches.json",
        help="Output JSON file.",
    )

    args = parser.parse_args()

    with open(args.urls, "r", encoding="utf-8") as f:
        urls = json.load(f)

    if not isinstance(urls, list):
        raise ValueError("URLs file must contain a JSON list.")

    results = []

    for url in urls:
        try:
            result = process_image(
                url=url,
                section_size=args.size,
                stride=args.stride,
                minimum_score=args.threshold,
            )

            results.append(result)

        except Exception as error:
            print(f"  ERROR: {error}")

            results.append(
                {
                    "url": url,
                    "error": str(error),
                }
            )

    output_path = Path(args.output)

    with open(output_path, "w", encoding="utf-8") as f:
        json.dump(
            results,
            f,
            indent=2,
        )

    print()
    print(f"Saved results to: {output_path}")


if __name__ == "__main__":
    main()
