package service

import (
	"API/repositories"
	"API/model"
	"database/sql"
)

type AnnonceService struct {
}

func NewAnnonceService(database *sql.DB) *AnnonceService {
	db = database
	return &AnnonceService{}
}

func (s *AnnonceService) GetAllAnnonces() ([]model.Annonce, error) {
	return repositories.GetAllAnnonces(db)
}

func (s *AnnonceService) GetAnnonceByID(id int) (*model.Annonce, error) {
	return repositories.GetAnnonceByID(db, id)
}

func (s *AnnonceService) GetAnnoncesByIDAnnonceur(idAnnonceur int) ([]model.Annonce, error) {
	return repositories.GetAnnoncesByIDAnnonceur(db, idAnnonceur)
}

func (s *AnnonceService) CreateAnnonce(a *model.Annonce) error {
	return repositories.CreateAnnonce(db, a)
}

func (s *AnnonceService) UpdateAnnonce(a *model.Annonce) error {
	return repositories.UpdateAnnonce(db, a)
}

func (s *AnnonceService) DeleteAnnonce(id int) error {
	return repositories.DeleteAnnonce(db, id)
}