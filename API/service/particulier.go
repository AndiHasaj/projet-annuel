package service

import (
	"API/model"
	"API/repositories"
	"database/sql"
)

type ParticulierService struct {
}

var db *sql.DB
func NewParticulierService(database *sql.DB) *ParticulierService {
	db = database
	return &ParticulierService{}
}

func (s *ParticulierService) GetAllParticuliers() ([]model.Particulier, error) {
	return repositories.GetAllParticuliers(db)
}

func (s *ParticulierService) GetParticulierByID(id int) (*model.Particulier, error) {
	return repositories.GetParticulierByID(db, id)
}

func (s *ParticulierService) GetParticulierByEmail(email string) (*model.Particulier, error) {
	return repositories.GetParticulierByEmail(db, email)
}

func (s *ParticulierService) CreateParticulier(p *model.Particulier) error {
	id, err := repositories.CreateParticulier(db, p)
	if err != nil {
		return err
	}
	p.ID = int(id)
	return nil
}

func (s *ParticulierService) UpdateParticulier(p *model.Particulier) error {
	return repositories.UpdateParticulier(db, p)
}

func (s *ParticulierService) DeleteParticulier(id int) error {
	return repositories.DeleteParticulier(db, id)
}

func (s *ParticulierService) UpdateScoreParticulier(id int, score int) error {
	return repositories.UpdateScoreParticulier(db, id, score)
}