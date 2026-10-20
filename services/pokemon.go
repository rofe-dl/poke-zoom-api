package services

import (
	"math/rand/v2"
	"net/http"

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
		return nil, 0, utils.NewHTTPError(http.StatusInternalServerError, "Error fetching count of Pokemon", countResult.Error)
	}

	result := service.DB.
		Scopes(
			utils.Paginate(queryParams.Offset, queryParams.Limit),
			utils.Sort(queryParams.Sort, queryParams.SortOrder)).
		Where(SEARCH_QUERY, search).
		Find(&pokemon)

	if result.Error != nil {
		// TODO: add some logger
		return nil, 0, utils.NewHTTPError(http.StatusInternalServerError, "Error fetching all Pokemon", result.Error)
	}

	return pokemon, totalCount, nil
}

func (service PokemonService) GetRandomCroppedPokemon() (string, error) {
	const TOTAL_COUNT = 649

	// TODO: Keep track of duplicates
	randomIndex := rand.IntN(TOTAL_COUNT + 1)

	var pokemon models.Pokemon

	result := service.DB.First(&pokemon, randomIndex)

	if result.RowsAffected == 0 {
		return "", utils.NewHTTPError(http.StatusNotFound, "Pokemon not found")
	} else if result.Error != nil {
		return "", utils.NewHTTPError(http.StatusInternalServerError, "Error fetching random Pokemon", result.Error)
	}

	base64EncodedImage, err := utils.GetRandomImageCrop(pokemon.OfficialArtwork, 128)

	if err != nil {
		return "", err
	}

	return base64EncodedImage, nil

}
