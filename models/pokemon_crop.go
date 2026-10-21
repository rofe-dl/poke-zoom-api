package models

import "time"

type PokemonCrop struct {
	ID        string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" gorm:"type:uuid;primaryKey;not null;default:gen_random_uuid()"`
	PokemonID uint   `json:"pokemon_id" example:"25" gorm:"not null"`
	GameID    string `json:"game_id" example:"123e4567-e89b-12d3-a456-426614174000" gorm:"type:uuid;not null"`
	CropCoordinates
	IsAnswered bool      `json:"is_answered" gorm:"not null;default:false"`
	IsCorrect  bool      `json:"is_correct"`
	TimeTaken  int       `json:"time_taken" example:"12000"`
	TriesTaken int       `json:"tries_taken" example:"2" gorm:"not null;default:0"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"not null"`

	Pokemon Pokemon `json:"pokemon" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Game    Game    `json:"game" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// coordinates are space separated "x1 y1 x2 y2" to represented top left and bottom right points of rectangle
type CropCoordinates struct {
	EasyCoordinates   string `json:"easy_coordinates" gorm:"not null"`
	MediumCoordinates string `json:"medium_coordinates" gorm:"not null"`
	HardCoordinates   string `json:"hard_coordinates" gorm:"not null"`
}

func (PokemonCrop) TableName() string {
	return "pokemon_crops"
}
