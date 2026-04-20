package repositories

import (
	"database/sql"
	"log"
	"API/model"
)

func GetAllProfessionnels(db *sql.DB) ([]model.Professionnel, error) {
	rows, err := db.Query("SELECT id, nom_entreprise, email, mot_de_passe, numero_siret, telephone, site_web, adresse, code_postal, ville, derniere_connexion, created_at, updated_at FROM professionnel")
	if err != nil {
		log.Printf("Erreur lors de la récupération des professionnels: %v", err)
		return nil, err
	}
	defer rows.Close()

	var professionnels []model.Professionnel
	for rows.Next() {
		var p model.Professionnel
		err := rows.Scan(&p.ID, &p.NomEntreprise, &p.Email, &p.MotDePasse, &p.NumeroSiret, &p.Telephone, &p.SiteWeb, &p.Adresse, &p.CodePostal, &p.Ville, &p.DerniereConnexion, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			log.Printf("Erreur lors du scan du professionnel: %v", err)
			return nil, err
		}
		professionnels = append(professionnels, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur après l'itération des professionnels: %v", err)
		return nil, err
	}

	return professionnels, nil
}

func GetProfessionnelByID(db *sql.DB, id int) (*model.Professionnel, error) {
	row := db.QueryRow("SELECT id, nom_entreprise, email, mot_de_passe, numero_siret, telephone, site_web, adresse, code_postal, ville, derniere_connexion, created_at, updated_at FROM professionnel WHERE id = ?", id)
	
	var p model.Professionnel
	err := row.Scan(&p.ID, &p.NomEntreprise, &p.Email, &p.MotDePasse, &p.NumeroSiret, &p.Telephone, &p.SiteWeb, &p.Adresse, &p.CodePostal, &p.Ville, &p.DerniereConnexion, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Professionnel avec l'ID %d non trouvé", id)
			return nil, nil
		}
		log.Printf("Erreur lors de la récupération du professionnel par ID: %v", err)
		return nil, err
	}

	return &p, nil
}

func GetProfessionnelByEmail(db *sql.DB, email string) (*model.Professionnel, error) {
	row := db.QueryRow("SELECT id, nom_entreprise, email, mot_de_passe, numero_siret, telephone, site_web, adresse, code_postal, ville, derniere_connexion, created_at, updated_at FROM professionnel WHERE email = ?", email)
	
	var p model.Professionnel
	err := row.Scan(&p.ID, &p.NomEntreprise, &p.Email, &p.MotDePasse, &p.NumeroSiret, &p.Telephone, &p.SiteWeb, &p.Adresse, &p.CodePostal, &p.Ville, &p.DerniereConnexion, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Professionnel avec l'email %s non trouvé", email)
			return nil, nil // Professionnel non trouvé
		}
		log.Printf("Erreur lors de la récupération du professionnel par email: %v", err)
		return nil, err
	}

	return &p, nil
}

func CreateProfessionnel(db *sql.DB, p *model.Professionnel) error {
	result, err := db.Exec("INSERT INTO professionnel (nom_entreprise, email, mot_de_passe, numero_siret, telephone, site_web, adresse, code_postal, ville, derniere_connexion) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		p.NomEntreprise, p.Email, p.MotDePasse, p.NumeroSiret, p.Telephone, p.SiteWeb, p.Adresse, p.CodePostal, p.Ville, p.DerniereConnexion)
	if err != nil {
		log.Printf("Erreur lors de la création du professionnel: %v", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Erreur lors de la récupération de l'ID du professionnel créé: %v", err)
		return err
	}
	p.ID = int(id)

	return nil
}

func UpdateProfessionnel(db *sql.DB, p *model.Professionnel) error {
	_, err := db.Exec("UPDATE professionnel SET nom_entreprise = ?, email = ?, mot_de_passe = ?, numero_siret = ?, telephone = ?, site_web = ?, adresse = ?, code_postal = ?, ville = ?, derniere_connexion = ? WHERE id = ?",
		p.NomEntreprise, p.Email, p.MotDePasse, p.NumeroSiret, p.Telephone, p.SiteWeb, p.Adresse, p.CodePostal, p.Ville, p.DerniereConnexion, p.ID)
	if err != nil {
		log.Printf("Erreur lors de la mise à jour du professionnel: %v", err)
		return err
	}
	return nil
}

func DeleteProfessionnel(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM professionnel WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de la suppression du professionnel: %v", err)
		return err
	}
	return nil
}