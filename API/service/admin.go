package service

import (
	"API/repositories"
	"API/model"
	"database/sql"
)

type AdminService struct {
}

func NewAdminService(database *sql.DB) *AdminService {
	db = database
	return &AdminService{}
}

func (s *AdminService) GetAllAdmins() ([]model.Admin, error) {
	return repositories.GetAllAdmins(db)
}

func (s *AdminService) GetAdminByID(id int) (*model.Admin, error) {
	return repositories.GetAdminByID(db, id)
}

func (s *AdminService) CreateAdmin(a *model.Admin) error {
	return repositories.CreateAdmin(db, a)
}

func (s *AdminService) DeleteAdmin(id int) error {
	return repositories.DeleteAdmin(db, id)
}