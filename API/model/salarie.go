package model

import (
	"time"
)

type Salarie struct {
	ID                int       `json:"id"`
	Nom               string    `json:"nom"`
	Prenom            string    `json:"prenom"`
	Email             string    `json:"email"`
	MotDePasse        string    `json:"mot_de_passe"`
	IntitulePoste     string    `json:"intitule_poste"`
	TypeContrat       string    `json:"type_contrat"`
	Salaire           int       `json:"salaire"`
	DateEmbauche      time.Time `json:"date_embauche"`
	DerniereConnexion time.Time `json:"derniere_connexion"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
