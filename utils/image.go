package utils

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/draw"
	"image/jpeg"
	"math/rand"
	"net/http"
	"time"

	_ "golang.org/x/image/webp"
	_ "image/png"
)

func GetRandomImageCrop(imageURL string, cropSize int) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(imageURL)

	if err != nil || resp.StatusCode != http.StatusOK {
		return "", NewHTTPError(http.StatusBadGateway, "Could not fetch Pokemon sprite!", err)
	}

	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)

	if err != nil {
		return "", NewHTTPError(http.StatusInternalServerError, "Error decoding Pokemon sprite.", err)
	}

	bounds := img.Bounds()

	cropW := cropSize
	if cropW > bounds.Dx() {
		cropW = bounds.Dx()
	}

	cropH := cropSize
	if cropH > bounds.Dy() {
		cropH = bounds.Dy()
	}

	// find random starting point of crop
	x := bounds.Min.X + rand.Intn(bounds.Dx()-cropW+1)
	y := bounds.Min.Y + rand.Intn(bounds.Dy()-cropH+1)

	crop := image.Rectangle{
		Min: image.Pt(x, y),
		Max: image.Pt(x+cropW, y+cropH),
	}

	cropped := image.NewRGBA(crop)
	draw.Draw(cropped, cropped.Bounds(), img, crop.Min, draw.Over)

	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, cropped, nil); err != nil {
		return "", NewHTTPError(http.StatusInternalServerError, "Error encoding cropped Pokemon sprite.", err)
	}

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())

	return encoded, nil
}
