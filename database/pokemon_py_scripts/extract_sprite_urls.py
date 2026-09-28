import json

urls = []

# Loop through gen1.json up to gen5.json
for i in range(1, 6):
    filename = f"pokemon_data/gen{i}.json"
    try:
        with open(filename, "r") as f:
            data = json.load(f)
            # Iterate through each entry in the generation file
            for entry in data.values():
                # Extract only the front_default sprite URL
                sprite_url = entry.get("sprites", {}).get("front_default")
                if sprite_url:
                    urls.append(sprite_url)
    except FileNotFoundError:
        print(f"Warning: {filename} not found, skipping.")

# Save the array of URLs to urls.json
with open("urls.json", "w") as f:
    json.dump(urls, f, indent=4)

print(f"Successfully extracted {len(urls)} URLs into urls.json")
