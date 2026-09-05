package services

import (
	"fmt"

	"github.com/rofe-dl/poke-zoom-api/models"
	"github.com/rofe-dl/poke-zoom-api/utils"
	"gorm.io/gorm"
)

type PokemonService struct {
	DB *gorm.DB
}

func (service PokemonService) FetchAllPokemon(queryParams *models.PokemonListQueryParams) ([]models.Pokemon, int64, error) {
	var pokemon []models.Pokemon
	var totalCount int64

	search := "%" + queryParams.Search + "%"

	const SEARCH_QUERY = `
		name ILIKE ?
	`

	// TODO: do both queries in goroutines

	countResult := service.DB.
		Model(&models.Pokemon{}).
		Where(SEARCH_QUERY, search).
		Count(&totalCount)

	if countResult.Error != nil {
		return nil, 0, countResult.Error
	}

	result := service.DB.
		Scopes(
			utils.Paginate(queryParams.Offset, queryParams.Limit),
			utils.Sort(queryParams.Sort, queryParams.SortOrder)).
		Where(SEARCH_QUERY, search).
		Find(&pokemon)

	if result.Error != nil {
		// TODO: add some logger
		fmt.Printf("Error fetching pokemon: %v", result.Error)

		return nil, 0, result.Error
	}

	return pokemon, totalCount, nil
}
