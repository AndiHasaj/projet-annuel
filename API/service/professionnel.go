package service

import (
	"API/model"
	"API/repositories"
	"database/sql"
)

type ProfessionnelService struct {
}

func NewProfessionnelService(database *sql.DB) *ProfessionnelService {
	db = database
	return &ProfessionnelService{}
}

func (s *ProfessionnelService) GetAllProfessionnels() ([]model.Professionnel, error) {
	return repositories.GetAllProfessionnels(db)
}

func (s *ProfessionnelService) GetProfessionnelByID(id int) (*model.Professionnel, error) {
	return repositories.GetProfessionnelByID(db, id)
}

func (s *ProfessionnelService) GetProfessionnelByEmail(email string) (*model.Professionnel, error) {
	return repositories.GetProfessionnelByEmail(db, email)
}

func (s *ProfessionnelService) CreateProfessionnel(p *model.Professionnel) error {
	return repositories.CreateProfessionnel(db, p)
}

func (s *ProfessionnelService) UpdateProfessionnel(p *model.Professionnel) error {
	return repositories.UpdateProfessionnel(db, p)
}

func (s *ProfessionnelService) DeleteProfessionnel(id int) error {
	return repositories.DeleteProfessionnel(db, id)
}