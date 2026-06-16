package model

import (
	"time"
)

type Professionnel struct {
	ID                int       `json:"id"`
	NomEntreprise     string    `json:"nom_entreprise"`
	Email             string    `json:"email"`
	MotDePasse        string    `json:"mot_de_passe"`
	NumeroSiret       string    `json:"numero_siret"`
	Telephone         string    `json:"telephone"`
	SiteWeb           string    `json:"site_web"`
	Adresse           string    `json:"adresse"`
	CodePostal        string    `json:"code_postal"`
	Ville             string    `json:"ville"`
	DerniereConnexion time.Time `json:"derniere_connexion"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
