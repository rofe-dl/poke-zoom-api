package main

import (
	"log"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/rofe-dl/poke-zoom-api/routes"
)

func main() {
	envError := godotenv.Load(".env")

	if envError != nil {
		log.Fatal("Error loading .env file")
	}

	db := connectDatabase()

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
