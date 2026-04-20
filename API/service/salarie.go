package service

import (
	"API/model"
	"API/repositories"
	"database/sql"
)

type SalarieService struct {
}

func NewSalarieService(database *sql.DB) *SalarieService {
	db = database
	return &SalarieService{}
}

func (s *SalarieService) GetAllSalaries() ([]model.Salarie, error) {
	return repositories.GetAllSalaries(db)
}

func (s *SalarieService) GetSalarieByID(id int) (*model.Salarie, error) {
	return repositories.GetSalarieByID(db, id)
}

func (s *SalarieService) GetSalarieByEmail(email string) (*model.Salarie, error) {
	return repositories.GetSalarieByEmail(db, email)
}

func (s *SalarieService) CreateSalarie(sal *model.Salarie) error {
	return repositories.CreateSalarie(db, sal)
}

func (s *SalarieService) UpdateSalarie(sal *model.Salarie) error {
	return repositories.UpdateSalarie(db, sal)
}

func (s *SalarieService) DeleteSalarie(id int) error {
	return repositories.DeleteSalarie(db, id)
}