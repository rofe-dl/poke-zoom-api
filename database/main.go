package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type Sprites struct {
	FrontDefault string `json:"front_default"`
}

type TypeObj struct {
	Name string `json:"name"`
}

type TypeSlot struct {
	Slot int     `json:"slot"`
	Type TypeObj `json:"type"`
}

type PokemonJSON struct {
	Name    string     `json:"name"`
	ID      int        `json:"id"`
	Sprites Sprites    `json:"sprites"`
	Types   []TypeSlot `json:"types"`
}

func main() {
	host := getEnv("DB_HOST_WITHOUT_DOCKER", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "poke_zoom")
	password := getEnv("DB_PASSWORD", "poke_zoom")
	dbname := getEnv("DB_NAME", "poke_zoom_db")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Database unreachable: %v", err)
	}

	for i := 1; i <= 5; i++ {
		filename := fmt.Sprintf("database/seeds/gen%d.json", i)
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			fmt.Printf("File %s not found, skipping...\n", filename)
			continue
		}

		fmt.Printf("Processing %s...\n", filename)
		fileBytes, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalf("Failed to read file %s: %v", filename, err)
		}

		var pokemonMap map[string]PokemonJSON
		if err := json.Unmarshal(fileBytes, &pokemonMap); err != nil {
			log.Fatalf("Failed to parse JSON in %s: %v", filename, err)
		}

		tx, err := db.Begin()
		if err != nil {
			log.Fatalf("Failed to begin transaction: %v", err)
		}

		stmt, err := tx.Prepare(`
			INSERT INTO pokemon (name, official_artwork, id, primary_type, secondary_type, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			ON CONFLICT (id) DO UPDATE SET
				name = EXCLUDED.name,
				official_artwork = EXCLUDED.official_artwork,
				primary_type = EXCLUDED.primary_type,
				secondary_type = EXCLUDED.secondary_type,
				updated_at = EXCLUDED.updated_at
		`)
		if err != nil {
			tx.Rollback()
			log.Fatalf("Failed to prepare statement: %v", err)
		}

		now := time.Now()

		for _, p := range pokemonMap {
			var primaryType, secondaryType string

			for _, t := range p.Types {
				if t.Slot == 1 {
					primaryType = strings.ToUpper(t.Type.Name[:1]) + t.Type.Name[1:]
				} else if t.Slot == 2 {
					secondaryType = strings.ToUpper(t.Type.Name[:1]) + t.Type.Name[1:]
				}
			}

			if primaryType == "" && len(p.Types) > 0 {
				primaryType = p.Types[0].Type.Name
			}
			var secArg interface{}
			if secondaryType != "" {
				secArg = secondaryType
			} else {
				secArg = nil
			}

			_, err := stmt.Exec(
				strings.ToUpper(p.Name[:1])+p.Name[1:],
				p.Sprites.FrontDefault,
				p.ID,
				primaryType,
				secArg,
				now,
				now,
			)
			if err != nil {
				tx.Rollback()
				log.Fatalf("Failed to insert pokemon ID %d (%s): %v", p.ID, p.Name, err)
			}
		}

		stmt.Close()
		if err := tx.Commit(); err != nil {
			log.Fatalf("Failed to commit transaction for %s: %v", filename, err)
		}

		fmt.Printf("Successfully seeded %s\n", filename)
	}

	fmt.Println("All generation files seeded successfully!")
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
