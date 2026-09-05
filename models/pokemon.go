package models

import (
	"time"

	"github.com/rofe-dl/poke-zoom-api/utils"
)

type Pokemon struct {
	ID              uint      `json:"id" example:"1" gorm:"primaryKey;not null"`
	Name            string    `json:"name" example:"Pikachu" gorm:"not null"`
	OfficialArtwork string    `json:"official_artwork" gorm:"not null" example:"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/1.png"`
	PrimaryType     string    `json:"primary_type" example:"fire" gorm:"not null"`
	SecondaryType   string    `json:"secondary_type" example:"dragon" required:"false"`
	CreatedAt       time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"not null"`
}

type PokemonListResponse struct {
	Body struct {
		Data       []Pokemon `json:"data"`
		TotalCount int64     `json:"total_count"`
	}
}

type PokemonIDInput struct {
	ID uint `path:"id" doc:"ID of the Pokemon" example:"1"`
}

type PokemonListQueryParams struct {
	utils.PaginationQueryParams
	utils.SearchQueryParams
	Sort string `query:"sort" enum:"name,created_at,id" example:"name" required:"false"`
}

func (Pokemon) TableName() string {
	return "pokemon"
}
