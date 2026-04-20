package service

import (
	"API/repositories"
	"API/model"
	"database/sql"
)

type ConteneurService struct {
}

func NewConteneurService(database *sql.DB) *ConteneurService {
	db = database
	return &ConteneurService{}
}

func (s *ConteneurService) GetAllConteneurs() ([]model.Conteneur, error) {
	return repositories.GetAllConteneurs(db)
}

func (s *ConteneurService) GetConteneurByID(id int) (*model.Conteneur, error) {
	return repositories.GetConteneurByID(db, id)
}

func (s *ConteneurService) GetConteneurByMatricule(matricule string) (*model.Conteneur, error) {
	return repositories.GetConteneurByMatricule(db, matricule)
}

func (s *ConteneurService) GetConteneursByLocalisation(localisation string) ([]model.Conteneur, error) {
	return repositories.GetConteneursByLocalisation(db, localisation)
}

func (s *ConteneurService) GetConteneurByCode(code string) (*model.Conteneur, error) {
	return repositories.GetConteneurByCode(db, code)
}

func (s *ConteneurService) CreateConteneur(c *model.Conteneur) error {
	return repositories.CreateConteneur(db, c)
}

func (s *ConteneurService) UpdateConteneur(c *model.Conteneur) error {
	return repositories.UpdateConteneur(db, c)
}

func (s *ConteneurService) DeleteConteneur(id int) error {
	return repositories.DeleteConteneur(db, id)
}