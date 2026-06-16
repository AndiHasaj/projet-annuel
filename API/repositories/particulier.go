package repositories

import (
	"database/sql"
	"log"
	"API/model"
)

func GetAllParticuliers(db *sql.DB) ([]model.Particulier, error) {
	rows, err := db.Query("SELECT id, nom, prenom, email, mot_de_passe, score, derniere_connexion, created_at, updated_at FROM particulier")
	if err != nil {
		log.Printf("Erreur lors de la récupération des particuliers: %v", err)
		return nil, err
	}
	defer rows.Close()

	var particuliers []model.Particulier
	for rows.Next() {
		var p model.Particulier
		err := rows.Scan(&p.ID, &p.Nom, &p.Prenom, &p.Email, &p.MotDePasse, &p.Score, &p.DerniereConnexion, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			log.Printf("Erreur lors du scan du particulier: %v", err)
			return nil, err
		}
		particuliers = append(particuliers, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur après l'itération des particuliers: %v", err)
		return nil, err
	}

	return particuliers, nil
}

func GetParticulierByID(db *sql.DB, id int) (*model.Particulier, error) {
	row := db.QueryRow("SELECT id, nom, prenom, email, mot_de_passe, score, derniere_connexion, created_at, updated_at FROM particulier WHERE id = ?", id)
	
	var p model.Particulier
	err := row.Scan(&p.ID, &p.Nom, &p.Prenom, &p.Email, &p.MotDePasse, &p.Score, &p.DerniereConnexion, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Particulier avec l'ID %d non trouvé", id)
			return nil, nil // Particulier non trouvé
		}
		log.Printf("Erreur lors de la récupération du particulier par ID: %v", err)
		return nil, err
	}

	return &p, nil
}

func GetParticulierByEmail(db *sql.DB, email string) (*model.Particulier, error) {
	row := db.QueryRow("SELECT id, nom, prenom, email, mot_de_passe, score, derniere_connexion, created_at, updated_at FROM particulier WHERE email = ?", email)
	
	var p model.Particulier
	err := row.Scan(&p.ID, &p.Nom, &p.Prenom, &p.Email, &p.MotDePasse, &p.Score, &p.DerniereConnexion, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Particulier avec l'email %s non trouvé", email)
			return nil, nil // Particulier non trouvé
		}
		log.Printf("Erreur lors de la récupération du particulier par email: %v", err)
		return nil, err
	}

	return &p, nil
}

func CreateParticulier(db *sql.DB, p *model.Particulier) (int64, error) {
	result, err := db.Exec("INSERT INTO particulier (nom, prenom, email, mot_de_passe) VALUES (?, ?, ?, ?)", p.Nom, p.Prenom, p.Email, p.MotDePasse)
	if err != nil {
		log.Printf("Erreur lors de la création du particulier: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Erreur lors de la récupération de l'ID du particulier créé: %v", err)
		return 0, err
	}

	return id, nil
}

func UpdateParticulier(db *sql.DB, p *model.Particulier) error {
	_, err := db.Exec("UPDATE particulier SET nom = ?, prenom = ?, email = ?, mot_de_passe = ?,  derniere_connexion = ? WHERE id = ?", p.Nom, p.Prenom, p.Email, p.MotDePasse, p.DerniereConnexion, p.ID)
	if err != nil {
		log.Printf("Erreur lors de la mise à jour du particulier: %v", err)
		return err
	}
	return nil
}

func DeleteParticulier(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM particulier WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de la suppression du particulier: %v", err)
		return err
	}
	return nil
}

func UpdateScoreParticulier(db *sql.DB, id int, score int) error {
	_, err := db.Exec("UPDATE particulier SET score = ? WHERE id = ?", score, id)
	if err != nil {
		log.Printf("Erreur lors de la mise à jour du score du particulier: %v", err)
		return err
	}
	return nil
}

