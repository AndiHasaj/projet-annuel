package model

import (
	"time"
)

type Annonce struct {
	ID              int       `json:"id"`
	IDAnnonceur     int       `json:"id_annonceur"`
	Titre           string    `json:"titre"`
	Description     string    `json:"description"`
	Prix            float32   `json:"prix"`
	DatePublication time.Time `json:"date_publication"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
