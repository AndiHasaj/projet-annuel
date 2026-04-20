package service

import (
	"API/repositories"
	"API/model"
	"database/sql"
)

type PaiementService struct {
}

func NewPaiementService(database *sql.DB) *PaiementService {
	db = database
	return &PaiementService{}
}

func (s *PaiementService) GetAllPaiements() ([]model.Paiement, error) {
	return repositories.GetAllPaiements(db)
}

func (s *PaiementService) GetPaiementByID(id int) (*model.Paiement, error) {
	return repositories.GetPaiementByID(db, id)
}

func (s *PaiementService) GetPaiementsByIDPayeur(idPayeur int) ([]model.Paiement, error) {
	return repositories.GetPaiementsByIDPayeur(db, idPayeur)
}

func (s *PaiementService) GetPaiementsByIDReceveur(idReceveur int) ([]model.Paiement, error) {
	return repositories.GetPaiementsByIDReceveur(db, idReceveur)
}

func (s *PaiementService) CreatePaiement(p *model.Paiement) error {
	return repositories.CreatePaiement(db, p)
}