package service

import (
	"API/repositories"
	"database/sql"
)

type SuiviFinancierService struct {
}

func NewSuiviFinancierService(database *sql.DB) *SuiviFinancierService {
	db = database
	return &SuiviFinancierService{}
}

func (s *SuiviFinancierService) GetChiffreAffaire(database *sql.DB) (float64, error) {
	return repositories.GetChiffreAffaire(db)
}

func (s *SuiviFinancierService) GetBenefice(database *sql.DB) (float64, error) {
	return repositories.GetBenefice(db)
}

func (s *SuiviFinancierService) GetTauxBenefice(database *sql.DB) (float64, error) {
	return repositories.GetTauxBenefice(db)
}

func (s *SuiviFinancierService) GetCharges(database *sql.DB) (float64, error) {
	return repositories.GetCharges(db)
}