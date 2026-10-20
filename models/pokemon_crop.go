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

type CropCoordinates struct {
	X1 int `json:"x1" example:"10" gorm:"not null"`
	Y1 int `json:"y1" example:"20" gorm:"not null"`
	X2 int `json:"x2" example:"50" gorm:"not null"`
	Y2 int `json:"y2" example:"60" gorm:"not null"`
	X3 int `json:"x3" example:"90" gorm:"not null"`
	Y3 int `json:"y3" example:"100" gorm:"not null"`
}

func (PokemonCrop) TableName() string {
	return "pokemon_crops"
}
