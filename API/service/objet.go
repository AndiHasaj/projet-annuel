package service

import (
	"API/repositories"
	"API/model"
	"database/sql"
)

type ObjetService struct {
}

func NewObjetService(database *sql.DB) *ObjetService {
	db = database
	return &ObjetService{}
}

func (s *ObjetService) GetAllObjets() ([]model.Objet, error) {
	return repositories.GetAllObjets(db)
}

func (s *ObjetService) GetObjetByID(id int) (*model.Objet, error) {
	return repositories.GetObjetByID(db, id)
}

func (s *ObjetService) GetObjetsByIDConteneur(id_conteneur int) ([]model.Objet, error) {
	return repositories.GetObjetsByIDConteneur(db, id_conteneur)
}

func (s *ObjetService) GetObjetsByType(type_objet string) ([]model.Objet, error) {
	return repositories.GetObjetsByType(db, type_objet)
}

func (s *ObjetService) GetObjetsByLocalisation(localisation string) ([]model.Objet, error) {
	return repositories.GetObjetsByLocalisation(db, localisation)
}

func (s *ObjetService) CreateObjet(o *model.Objet) error {
	return repositories.CreateObjet(db, o)
}

func (s *ObjetService) UpdateObjet(o *model.Objet) error {
	return repositories.UpdateObjet(db, o)
}

func (s *ObjetService) DeleteObjet(id int) error {
	return repositories.DeleteObjet(db, id)
}