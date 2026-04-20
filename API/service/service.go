package service

import (
	"API/repositories"
	"API/model"
	"database/sql"
)

type ServiceService struct {
}

func NewServiceService(database *sql.DB) *ServiceService {
	db = database
	return &ServiceService{}
}

func (s *ServiceService) GetAllServices() ([]model.Service, error) {
	return repositories.GetAllServices(db)
}

func (s *ServiceService) GetServiceByID(id int) (*model.Service, error) {
	return repositories.GetServiceByID(db, id)
}

func (s *ServiceService) GetServicesByType(type_service string) ([]model.Service, error) {
	return repositories.GetServicesByType(db, type_service)
}

func (s *ServiceService) CreateService(srv *model.Service) error {
	return repositories.CreateService(db, srv)
}

func (s *ServiceService) UpdateService(srv *model.Service) error {
	return repositories.UpdateService(db, srv)
}

func (s *ServiceService) DeleteService(id int) error {
	return repositories.DeleteService(db, id)
}