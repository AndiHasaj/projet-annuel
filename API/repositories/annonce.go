package repositories

import (
	"database/sql"
	"log"
	"API/model"
)

func GetAllAnnonces(db *sql.DB) ([]model.Annonce, error) {
	rows, err := db.Query("SELECT id, id_annonceur, titre, description, prix, date_publication, created_at, updated_at FROM annonce")
	if err != nil {
		log.Printf("Erreur lors de la récupération des annonces: %v", err)
		return nil, err
	}
	defer rows.Close()

	var annonces []model.Annonce
	for rows.Next() {
		var a model.Annonce
		err := rows.Scan(&a.ID, &a.IDAnnonceur, &a.Titre, &a.Description, &a.Prix, &a.DatePublication, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			log.Printf("Erreur lors du scan de l'annonce: %v", err)
			return nil, err
		}
		annonces = append(annonces, a)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur après l'itération des annonces: %v", err)
		return nil, err
	}

	return annonces, nil
}

func GetAnnonceByID(db *sql.DB, id int) (*model.Annonce, error) {
	row := db.QueryRow("SELECT id, id_annonceur, titre, description, prix, date_publication, created_at, updated_at FROM annonce WHERE id = ?", id)
	
	var a model.Annonce
	err := row.Scan(&a.ID, &a.IDAnnonceur, &a.Titre, &a.Description, &a.Prix, &a.DatePublication, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Annonce avec l'ID %d non trouvée", id)
			return nil, nil // Annonce non trouvée
		}
		log.Printf("Erreur lors de la récupération de l'annonce par ID: %v", err)
		return nil, err
	}

	return &a, nil
}

func GetAnnoncesByIDAnnonceur(db *sql.DB, idAnnonceur int) ([]model.Annonce, error) {
	rows, err := db.Query("SELECT id, id_annonceur, titre, description, prix, date_publication, created_at, updated_at FROM annonce WHERE id_annonceur = ?", idAnnonceur)
	if err != nil {
		log.Printf("Erreur lors de la récupération des annonces par ID annonceur: %v", err)
		return nil, err
	}
	defer rows.Close()

	var annonces []model.Annonce
	for rows.Next() {
		var a model.Annonce
		err := rows.Scan(&a.ID, &a.IDAnnonceur, &a.Titre, &a.Description, &a.Prix, &a.DatePublication, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			log.Printf("Erreur lors du scan de l'annonce: %v", err)
			return nil, err
		}
		annonces = append(annonces, a)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur après l'itération des annonces: %v", err)
		return nil, err
	}

	return annonces, nil
}

func CreateAnnonce(db *sql.DB, a *model.Annonce) error {
	result, err := db.Exec("INSERT INTO annonce (id_annonceur, titre, description, prix, date_publication) VALUES (?, ?, ?, ?, ?)",
		a.IDAnnonceur, a.Titre, a.Description, a.Prix, a.DatePublication)
	if err != nil {
		log.Printf("Erreur lors de la création de l'annonce: %v", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Erreur lors de la récupération de l'ID de l'annonce créée: %v", err)
		return err
	}
	a.ID = int(id)

	return nil
}

func UpdateAnnonce(db *sql.DB, a *model.Annonce) error {
	_, err := db.Exec("UPDATE annonce SET id_annonceur = ?, titre = ?, description = ?, prix = ?, date_publication = ? WHERE id = ?",
		a.IDAnnonceur, a.Titre, a.Description, a.Prix, a.DatePublication, a.ID)
	if err != nil {
		log.Printf("Erreur lors de la mise à jour de l'annonce: %v", err)
		return err
	}
	return nil
}

func DeleteAnnonce(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM annonce WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de la suppression de l'annonce: %v", err)
		return err
	}
	return nil
}
