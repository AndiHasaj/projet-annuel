package repositories

import (
	"database/sql"
	"log"
	"API/model"
)

func GetAllServices(db *sql.DB) ([]model.Service, error) {
	rows, err := db.Query("SELECT id, type_service, titre, description, prix, date_service, created_at, updated_at FROM service")
	if err != nil {
		log.Printf("Erreur lors de la récupération des services: %v", err)
		return nil, err
	}
	defer rows.Close()

	var services []model.Service
	for rows.Next() {
		var s model.Service
		err := rows.Scan(&s.ID, &s.TypeService, &s.Titre, &s.Description, &s.Prix, &s.DateService, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			log.Printf("Erreur lors du scan du service: %v", err)
			return nil, err
		}
		services = append(services, s)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur après l'itération des services: %v", err)
		return nil, err
	}

	return services, nil
}

func GetServiceByID(db *sql.DB, id int) (*model.Service, error) {
	row := db.QueryRow("SELECT id, type_service, titre, description, prix, date_service, created_at, updated_at FROM service WHERE id = ?", id)
	
	var s model.Service
	err := row.Scan(&s.ID, &s.TypeService, &s.Titre, &s.Description, &s.Prix, &s.DateService, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Service avec l'ID %d non trouvé", id)
			return nil, nil // Service non trouvé
		}
		log.Printf("Erreur lors de la récupération du service par ID: %v", err)
		return nil, err
	}

	return &s, nil
}

func GetServicesByType(db *sql.DB, typeService string) ([]model.Service, error) {
	rows, err := db.Query("SELECT id, type_service, titre, description, prix, date_service, created_at, updated_at FROM service WHERE type_service = ?", typeService)
	if err != nil {
		log.Printf("Erreur lors de la récupération des services par type: %v", err)
		return nil, err
	}
	defer rows.Close()

	var services []model.Service
	for rows.Next() {
		var s model.Service
		err := rows.Scan(&s.ID, &s.TypeService, &s.Titre, &s.Description, &s.Prix, &s.DateService, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			log.Printf("Erreur lors du scan du service: %v", err)
			return nil, err
		}
		services = append(services, s)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur après l'itération des services: %v", err)
		return nil, err
	}

	return services, nil
}

func CreateService(db *sql.DB, s *model.Service) error {
	result, err := db.Exec("INSERT INTO service (type_service, titre, description, prix, date_service) VALUES (?, ?, ?, ?, ?)",
		s.TypeService, s.Titre, s.Description, s.Prix, s.DateService)
	if err != nil {
		log.Printf("Erreur lors de la création du service: %v", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Erreur lors de la récupération de l'ID du service créé: %v", err)
		return err
	}
	s.ID = int(id)

	return nil
}

func UpdateService(db *sql.DB, s *model.Service) error {
	_, err := db.Exec("UPDATE service SET type_service = ?, titre = ?, description = ?, prix = ?, date_service = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?",
		s.TypeService, s.Titre, s.Description, s.Prix, s.DateService, s.ID)
	if err != nil {
		log.Printf("Erreur lors de la mise à jour du service: %v", err)
		return err
	}
	return nil
}

func DeleteService(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM service WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de la suppression du service: %v", err)
		return err
	}
	return nil
}