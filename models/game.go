package models

import "time"

type Game struct {
	ID           string        `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Score        int           `json:"score" example:"150" gorm:"not null;default:0"`
	StartTime    time.Time     `json:"start_time" gorm:"not null"`
	EndTime      time.Time     `json:"end_time"`
	PokemonCrops []PokemonCrop `json:"pokemon_crops"`
	MaxDuration  int           `json:"max_duration" example:"300" gorm:"not null"`
	CreatedAt    time.Time     `json:"created_at" gorm:"not null"`
	UpdatedAt    time.Time     `json:"updated_at" gorm:"not null"`
}

func (Game) TableName() string {
	return "games"
}
