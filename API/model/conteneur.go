package model

import (
	"time"
)

type Conteneur struct {
	ID           int       `json:"id"`
	Matricule    string    `json:"matricule"`
	Localisation string    `json:"localisation"`
	Statut       string    `json:"statut"`
	Occupation   int       `json:"occupation"`
	Code         string    `json:"code"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
