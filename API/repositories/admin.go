package repositories

import (
	"database/sql"
	"log"
	"API/model"
)

func GetAllAdmins(db *sql.DB) ([]model.Admin, error) {
	rows, err := db.Query("SELECT id, mdp FROM admin")
	if err != nil {
		log.Printf("Erreur lors de la récupération des admins: %v", err)
		return nil, err
	}
	defer rows.Close()

	var admins []model.Admin
	for rows.Next() {
		var a model.Admin
		err := rows.Scan(&a.ID, &a.MotDePasse)
		if err != nil {
			log.Printf("Erreur lors du scan de l'admin: %v", err)
			return nil, err
		}
		admins = append(admins, a)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur après l'itération des admins: %v", err)
		return nil, err
	}

	return admins, nil
}

func GetAdminByID(db *sql.DB, id int) (*model.Admin, error) {
	row := db.QueryRow("SELECT id, mdp FROM admin WHERE id = ?", id)
	
	var a model.Admin
	err := row.Scan(&a.ID, &a.MotDePasse)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Admin avec l'ID %d non trouvé", id)
			return nil, nil // Admin non trouvé
		}
		log.Printf("Erreur lors de la récupération de l'admin par ID: %v", err)
		return nil, err
	}

	return &a, nil
}

func CreateAdmin(db *sql.DB, a *model.Admin) error {
	result, err := db.Exec("INSERT INTO admin (mdp) VALUES (?)", a.MotDePasse)
	if err != nil {
		log.Printf("Erreur lors de la création de l'admin: %v", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Erreur lors de la récupération de l'ID de l'admin créé: %v", err)
		return err
	}
	a.ID = int(id)

	return nil
}

func DeleteAdmin(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM admin WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de la suppression de l'admin: %v", err)
		return err
	}
	return nil
}