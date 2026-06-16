package model

import (
	"time"
)

type Objet struct {
	ID           int       `json:"id"`
	Nom          string    `json:"nom"`
	TypeObjet    string    `json:"type_objet"`
	Description  string    `json:"description"`
	Localisation string    `json:"localisation"`
	IDConteneur  int       `json:"id_conteneur"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
