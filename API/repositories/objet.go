package repositories

import (
	"database/sql"
	"log"
	"API/model"
)

func GetAllObjets(db *sql.DB) ([]model.Objet, error) {
	rows, err := db.Query("SELECT id, nom, type_objet, description, localisation, id_conteneur, created_at, updated_at FROM objet")
	if err != nil {
		log.Printf("Erreur lors de la récupération des objets: %v", err)
		return nil, err
	}
	defer rows.Close()

	var objets []model.Objet
	for rows.Next() {
		var o model.Objet
		err := rows.Scan(&o.ID, &o.Nom, &o.TypeObjet, &o.Description, &o.Localisation, &o.IDConteneur, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			log.Printf("Erreur lors du scan de l'objet: %v", err)
			return nil, err
		}
		objets = append(objets, o)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur après l'itération des objets: %v", err)
		return nil, err
	}

	return objets, nil
}

func GetObjetByID(db *sql.DB, id int) (*model.Objet, error) {
	row := db.QueryRow("SELECT id, nom, type_objet, description, localisation, id_conteneur, created_at, updated_at FROM objet WHERE id = ?", id)
	
	var o model.Objet
	err := row.Scan(&o.ID, &o.Nom, &o.TypeObjet, &o.Description, &o.Localisation, &o.IDConteneur, &o.CreatedAt, &o.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Objet avec l'ID %d non trouvé", id)
			return nil, nil // Objet non trouvé
		}
		log.Printf("Erreur lors de la récupération de l'objet par ID: %v", err)
		return nil, err
	}

	return &o, nil
}

func GetObjetsByIDConteneur(db *sql.DB, idConteneur int) ([]model.Objet, error) {
	rows, err := db.Query("SELECT id, nom, type_objet, description, localisation, id_conteneur, created_at, updated_at FROM objet WHERE id_conteneur = ?", idConteneur)
	if err != nil {
		log.Printf("Erreur lors de la récupération des objets par ID conteneur: %v", err)
		return nil, err
	}
	defer rows.Close()

	var objets []model.Objet
	for rows.Next() {
		var o model.Objet
		err := rows.Scan(&o.ID, &o.Nom, &o.TypeObjet, &o.Description, &o.Localisation, &o.IDConteneur, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			log.Printf("Erreur lors du scan de l'objet: %v", err)
			return nil, err
		}
		objets = append(objets, o)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur après l'itération des objets: %v", err)
		return nil, err
	}

	return objets, nil
}

func GetObjetsByType(db *sql.DB, typeObjet string) ([]model.Objet, error) {
	rows, err := db.Query("SELECT id, nom, type_objet, description, localisation, id_conteneur, created_at, updated_at FROM objet WHERE type_objet = ?", typeObjet)
	if err != nil {
		log.Printf("Erreur lors de la récupération des objets par type: %v", err)
		return nil, err
	}
	defer rows.Close()

	var objets []model.Objet
	for rows.Next() {
		var o model.Objet
		err := rows.Scan(&o.ID, &o.Nom, &o.TypeObjet, &o.Description, &o.Localisation, &o.IDConteneur, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			log.Printf("Erreur lors du scan de l'objet: %v", err)
			return nil, err
		}
		objets = append(objets, o)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur après l'itération des objets: %v", err)
		return nil, err
	}

	return objets, nil
}

func GetObjetsByLocalisation(db *sql.DB, localisation string) ([]model.Objet, error) {
	rows, err := db.Query("SELECT id, nom, type_objet, description, localisation, id_conteneur, created_at, updated_at FROM objet WHERE localisation = ?", localisation)
	if err != nil {
		log.Printf("Erreur lors de la récupération des objets par localisation: %v", err)
		return nil, err
	}
	defer rows.Close()

	var objets []model.Objet
	for rows.Next() {
		var o model.Objet
		err := rows.Scan(&o.ID, &o.Nom, &o.TypeObjet, &o.Description, &o.Localisation, &o.IDConteneur, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			log.Printf("Erreur lors du scan de l'objet: %v", err)
			return nil, err
		}
		objets = append(objets, o)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur après l'itération des objets: %v", err)
		return nil, err
	}

	return objets, nil
}

func CreateObjet(db *sql.DB, o *model.Objet) error {
	result, err := db.Exec("INSERT INTO objet (nom, type_objet, description, localisation, id_conteneur) VALUES (?, ?, ?, ?, ?)",
		o.Nom, o.TypeObjet, o.Description, o.Localisation, o.IDConteneur)
	if err != nil {
		log.Printf("Erreur lors de la création de l'objet: %v", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Erreur lors de la récupération de l'ID de l'objet créé: %v", err)
		return err
	}
	o.ID = int(id)

	return nil
}

func UpdateObjet(db *sql.DB, o *model.Objet) error {
	_, err := db.Exec("UPDATE objet SET nom = ?, type_objet = ?, description = ?, localisation = ?, id_conteneur = ?, updated_at = ? WHERE id = ?",
		o.Nom, o.TypeObjet, o.Description, o.Localisation, o.IDConteneur, o.UpdatedAt, o.ID)
	if err != nil {
		log.Printf("Erreur lors de la mise à jour de l'objet: %v", err)
		return err
	}

	return nil
}

func DeleteObjet(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM objet WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de la suppression de l'objet: %v", err)
		return err
	}

	return nil
}