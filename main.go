package main

import (
	"fmt"
	"log"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rofe-dl/poke-zoom-api/routes"
)

func main() {
	envError := godotenv.Load(".env")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		getSSLMode(),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Connected to Postgres!")

	if envError != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	routerGroup := r.Group("/api/v1")

	humaConfig := huma.DefaultConfig("Poke-zoom API", "1.0.0")
	humaConfig.Servers = []*huma.Server{
		{URL: "/api/v1"},
	}
	humaConfig.CreateHooks = nil // disable $schema field in docs

	humaApi := humagin.NewWithGroup(r, routerGroup, humaConfig)

	routes.RegisterAll(humaApi, db)

	port := os.Getenv("PORT")

	r.Run(":" + port)
}

func getSSLMode() string {
	env := os.Getenv("ENV")

	if env == "production" {
		return "require"
	}

	return "disable"
}
