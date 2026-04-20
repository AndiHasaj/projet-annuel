package config

import (
	"API/service"
	"database/sql"
)

type AppConfig struct {
	AdminService          *service.AdminService
	AnnonceService        *service.AnnonceService
	ConteneurService      *service.ConteneurService
	ObjetService          *service.ObjetService
	PaiementService       *service.PaiementService
	ParticulierService    *service.ParticulierService
	ProfessionnelService  *service.ProfessionnelService
	SalarieService        *service.SalarieService
	ServiceService        *service.ServiceService
	SuiviFinancierService *service.SuiviFinancierService
}

func NewAppConfig(db *sql.DB) *AppConfig {
	return &AppConfig{
		AdminService:          service.NewAdminService(db),
		AnnonceService:        service.NewAnnonceService(db),
		ConteneurService:      service.NewConteneurService(db),
		ObjetService:          service.NewObjetService(db),
		PaiementService:       service.NewPaiementService(db),
		ParticulierService:    service.NewParticulierService(db),
		ProfessionnelService:  service.NewProfessionnelService(db),
		SalarieService:        service.NewSalarieService(db),
		ServiceService:        service.NewServiceService(db),
		SuiviFinancierService: service.NewSuiviFinancierService(db),
	}
}