package routes

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/rofe-dl/poke-zoom-api/models"
	"github.com/rofe-dl/poke-zoom-api/services"
	"github.com/rofe-dl/poke-zoom-api/utils"
	"gorm.io/gorm"
)

func registerPokemonRoutes(api huma.API, db *gorm.DB) {
	service := services.PokemonService{DB: db}

	const PREFIX string = "/pokemon"

	huma.Register(api, huma.Operation{
		OperationID: "GetAllPokemon",
		Method:      http.MethodGet,
		Path:        PREFIX,
		Summary:     "Get all pokemon",
		Tags:        []string{"Pokemon"},
		Description: "Retrieves all Pokemon from the database",
	}, func(ctx context.Context, input *models.PokemonListQueryParams) (*models.PokemonListResponse, error) {
		pokemon, totalCount, err := service.FetchAllPokemon(input)

		if err != nil {
			return nil, utils.HandleHttpError(err)
		}

		result := &models.PokemonListResponse{}
		result.Body.Data = pokemon
		result.Body.TotalCount = totalCount

		return result, nil
	})
}
