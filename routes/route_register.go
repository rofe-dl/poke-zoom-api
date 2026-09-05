package routes

import (
	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func RegisterAll(api huma.API, db *gorm.DB) {
	registerPokemonRoutes(api, db)
}
