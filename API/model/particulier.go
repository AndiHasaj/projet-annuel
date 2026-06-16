package model

import (
	"time"
)

type Particulier struct {
	ID                int       `json:"id"`
	Nom               string    `json:"nom"`
	Prenom            string    `json:"prenom"`
	Email             string    `json:"email"`
	MotDePasse        string    `json:"mot_de_passe"`
	Score             int       `json:"score"`
	DerniereConnexion time.Time `json:"derniere_connexion"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
