package repositories

import (
	"database/sql"
	"log"
	"API/model"
)

func GetAllSalaries(db *sql.DB) ([]model.Salarie, error) {
	rows, err := db.Query("SELECT id, nom, prenom, email, mot_de_passe, intitule_poste, type_contrat, salaire, date_embauche, created_at, updated_at FROM salarie")
	if err != nil {
		log.Printf("Erreur lors de la récupération des salariés: %v", err)
		return nil, err
	}
	defer rows.Close()

	var salaries []model.Salarie
	for rows.Next() {
		var s model.Salarie
		err := rows.Scan(&s.ID, &s.Nom, &s.Prenom, &s.Email, &s.MotDePasse, &s.IntitulePoste, &s.TypeContrat, &s.Salaire, &s.DateEmbauche, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			log.Printf("Erreur lors du scan du salarié: %v", err)
			return nil, err
		}
		salaries = append(salaries, s)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur après l'itération des salariés: %v", err)
		return nil, err
	}

	return salaries, nil
}

func GetSalarieByID(db *sql.DB, id int) (*model.Salarie, error) {
	row := db.QueryRow("SELECT id, nom, prenom, email, mot_de_passe, intitule_poste, type_contrat, salaire, date_embauche, created_at, updated_at FROM salarie WHERE id = ?", id)
	
	var s model.Salarie
	err := row.Scan(&s.ID, &s.Nom, &s.Prenom, &s.Email, &s.MotDePasse, &s.IntitulePoste, &s.TypeContrat, &s.Salaire, &s.DateEmbauche, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Salarié avec l'ID %d non trouvé", id)
			return nil, nil // Salarié non trouvé
		}
		log.Printf("Erreur lors de la récupération du salarié par ID: %v", err)
		return nil, err
	}

	return &s, nil
}

func GetSalarieByEmail(db *sql.DB, email string) (*model.Salarie, error) {
	row := db.QueryRow("SELECT id, nom, prenom, email, mot_de_passe, intitule_poste, type_contrat, salaire, date_embauche, created_at, updated_at FROM salarie WHERE email = ?", email)
	
	var s model.Salarie
	err := row.Scan(&s.ID, &s.Nom, &s.Prenom, &s.Email, &s.MotDePasse, &s.IntitulePoste, &s.TypeContrat, &s.Salaire, &s.DateEmbauche, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Salarié avec l'email %s non trouvé", email)
			return nil, nil // Salarié non trouvé
		}
		log.Printf("Erreur lors de la récupération du salarié par email: %v", err)
		return nil, err
	}

	return &s, nil
}

func CreateSalarie(db *sql.DB, s *model.Salarie) error {
	result, err := db.Exec("INSERT INTO salarie (nom, prenom, email, mot_de_passe, intitule_poste, type_contrat, salaire, date_embauche) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", s.Nom, s.Prenom, s.Email, s.MotDePasse, s.IntitulePoste, s.TypeContrat, s.Salaire, s.DateEmbauche)
	if err != nil {
		log.Printf("Erreur lors de la création du salarié: %v", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Erreur lors de la récupération de l'ID du salarié créé: %v", err)
		return err
	}
	s.ID = int(id)

	return nil
}

func UpdateSalarie(db *sql.DB, s *model.Salarie) error {
	_, err := db.Exec("UPDATE salarie SET nom = ?, prenom = ?, email = ?, mot_de_passe = ?, intitule_poste = ?, type_contrat = ?, salaire = ?, date_embauche = ? WHERE id = ?", s.Nom, s.Prenom, s.Email, s.MotDePasse, s.IntitulePoste, s.TypeContrat, s.Salaire, s.DateEmbauche, s.ID)
	if err != nil {
		log.Printf("Erreur lors de la mise à jour du salarié: %v", err)
		return err
	}
	return nil
}

func DeleteSalarie(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM salarie WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de la suppression du salarié: %v", err)
		return err
	}
	return nil
}