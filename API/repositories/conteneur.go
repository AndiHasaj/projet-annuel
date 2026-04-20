package repositories

import (
	"database/sql"
	"log"
	"API/model"
)

func GetAllConteneurs(db *sql.DB) ([]model.Conteneur, error) {
	rows, err := db.Query("SELECT id, matricule, localisation, statut, occupation, code, created_at, updated_at FROM conteneur")
	if err != nil {
		log.Printf("Erreur lors de la récupération des conteneurs: %v", err)
		return nil, err
	}
	defer rows.Close()

	var conteneurs []model.Conteneur
	for rows.Next() {
		var c model.Conteneur
		err := rows.Scan(&c.ID, &c.Matricule, &c.Localisation, &c.Statut, &c.Occupation, &c.Code, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			log.Printf("Erreur lors du scan du conteneur: %v", err)
			return nil, err
		}
		conteneurs = append(conteneurs, c)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur après l'itération des conteneurs: %v", err)
		return nil, err
	}

	return conteneurs, nil
}

func GetConteneurByID(db *sql.DB, id int) (*model.Conteneur, error) {
	row := db.QueryRow("SELECT id, matricule, localisation, statut, occupation, code, created_at, updated_at FROM conteneur WHERE id = ?", id)
	
	var c model.Conteneur
	err := row.Scan(&c.ID, &c.Matricule, &c.Localisation, &c.Statut, &c.Occupation, &c.Code, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Conteneur avec l'ID %d non trouvé", id)
			return nil, nil // Conteneur non trouvé
		}
		log.Printf("Erreur lors de la récupération du conteneur par ID: %v", err)
		return nil, err
	}

	return &c, nil
}

func GetConteneurByMatricule(db *sql.DB, matricule string) (*model.Conteneur, error) {
	row := db.QueryRow("SELECT id, matricule, localisation, statut, occupation, code, created_at, updated_at FROM conteneur WHERE matricule = ?", matricule)
	
	var c model.Conteneur
	err := row.Scan(&c.ID, &c.Matricule, &c.Localisation, &c.Statut, &c.Occupation, &c.Code, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Conteneur avec le matricule %s non trouvé", matricule)
			return nil, nil // Conteneur non trouvé
		}
		log.Printf("Erreur lors de la récupération du conteneur par matricule: %v", err)
		return nil, err
	}

	return &c, nil
}

func GetConteneursByLocalisation(db *sql.DB, localisation string) ([]model.Conteneur, error) {
	rows, err := db.Query("SELECT id, matricule, localisation, statut, occupation, code, created_at, updated_at FROM conteneur WHERE localisation = ?", localisation)
	if err != nil {
		log.Printf("Erreur lors de la récupération des conteneurs par localisation: %v", err)
		return nil, err
	}
	defer rows.Close()

	var conteneurs []model.Conteneur
	for rows.Next() {
		var c model.Conteneur
		err := rows.Scan(&c.ID, &c.Matricule, &c.Localisation, &c.Statut, &c.Occupation, &c.Code, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			log.Printf("Erreur lors du scan du conteneur: %v", err)
			return nil, err
		}
		conteneurs = append(conteneurs, c)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur après l'itération des conteneurs: %v", err)
		return nil, err
	}

	return conteneurs, nil
}

func GetConteneurByCode(db *sql.DB, code string) (*model.Conteneur, error) {
	row := db.QueryRow("SELECT id, matricule, localisation, statut, occupation, code, created_at, updated_at FROM conteneur WHERE code = ?", code)
	var c model.Conteneur
	err := row.Scan(&c.ID, &c.Matricule, &c.Localisation, &c.Statut, &c.Occupation, &c.Code, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Conteneur avec le code %s non trouvé", code)
			return nil, nil // Conteneur non trouvé
		}
		log.Printf("Erreur lors de la récupération du conteneur par code: %v", err)
		return nil, err
	}

	return &c, nil
}

func CreateConteneur(db *sql.DB, c *model.Conteneur) error {
	result, err := db.Exec("INSERT INTO conteneur (matricule, localisation, statut, occupation, code, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		c.Matricule, c.Localisation, c.Statut, c.Occupation, c.Code, c.CreatedAt, c.UpdatedAt)
	if err != nil {
		log.Printf("Erreur lors de la création du conteneur: %v", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Erreur lors de la récupération de l'ID du conteneur créé: %v", err)
		return err
	}
	c.ID = int(id)

	return nil
}

func UpdateConteneur(db *sql.DB, c *model.Conteneur) error {
	_, err := db.Exec("UPDATE conteneur SET matricule = ?, localisation = ?, statut = ?, occupation = ?, code = ?, updated_at = ? WHERE id = ?",
		c.Matricule, c.Localisation, c.Statut, c.Occupation, c.Code, c.UpdatedAt, c.ID)
	if err != nil {
		log.Printf("Erreur lors de la mise à jour du conteneur: %v", err)
		return err
	}
	return nil
}

func DeleteConteneur(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM conteneur WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de la suppression du conteneur: %v", err)
		return err
	}
	return nil
}