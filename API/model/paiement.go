package model

import (
	"time"
)

type Paiement struct {
	ID           int       `json:"id"`
	IDAnnonce    int       `json:"id_annonce"`
	IDService    int       `json:"id_service"`
	IDPayeur     int       `json:"id_payeur"`
	IDReceveur   int       `json:"id_receveur"`
	Montant      float32   `json:"montant"`
	DatePaiement time.Time `json:"date_paiement"`
	TypePaiement string    `json:"type_paiement"`
	Statut       string    `json:"statut"`
	CreatedAt    time.Time `json:"created_at"`
}
