import json
import os
import time
import requests

# Generation cutoffs (Inclusive ranges for National Dex IDs up to Gen 5)
GENERATIONS = {
    "gen1": (1, 151),
    "gen2": (152, 251),
    "gen3": (252, 386),
    "gen4": (387, 493),
    "gen5": (494, 649),
}

DELAY_SECONDS = 0.25  # Rate limit delay between API calls
OUTPUT_DIR = "pokemon_data"


def fetch_pokemon_data():
  os.makedirs(OUTPUT_DIR, exist_ok=True)

  for gen_name, (start_id, end_id) in GENERATIONS.items():
    print(
        f"Fetching {gen_name.upper()} (Pokémon #{start_id} to #{end_id})..."
    )
    gen_data = {}

    for pokemon_id in range(start_id, end_id + 1):
      url = f"https://pokeapi.co/api/v2/pokemon/{pokemon_id}"

      try:
        response = requests.get(url)

        if response.status_code == 200:
          full_data = response.json()
          sprite = full_data.get("sprites").get("other").get("official-artwork")

          # Extract only the required fields
          filtered_data = {
              "name": full_data.get("name"),
              "id": full_data.get("id"),
              "sprites": sprite,
              "types": full_data.get("types"),
          }

          # Use string ID as the key in the JSON dictionary
          gen_data[str(pokemon_id)] = filtered_data
          print(
              f"  [Success] Fetched #{pokemon_id}: {filtered_data['name']}"
          )

        else:
          print(
              f"  [Error] Failed to fetch ID {pokemon_id}. Status:"
              f" {response.status_code}"
          )

      except requests.exceptions.RequestException as e:
        print(f"  [Exception] Network error on ID {pokemon_id}: {e}")

      # Enforce the rate limit delay between calls
      time.sleep(DELAY_SECONDS)

    # Save the generation's dictionary to a JSON file
    file_path = os.path.join(OUTPUT_DIR, f"{gen_name}.json")
    with open(file_path, "w", encoding="utf-8") as f:
      json.dump(gen_data, f, indent=4, ensure_ascii=False)

    print(f" Saved {gen_name}.json successfully!\n")

  print("All generations fetched and saved successfully.")


if __name__ == "__main__":
  fetch_pokemon_data()
