package repositories

import (
	"database/sql"
	"log"
	"API/model"
)

func GetAllPaiements(db *sql.DB) ([]model.Paiement, error) {
	rows, err := db.Query("SELECT id, id_annonce, id_service, id_payeur, id_receveur, montant, date_paiement, created_at FROM paiement")
	if err != nil {
		log.Printf("Erreur lors de la récupération des paiements: %v", err)
		return nil, err
	}
	defer rows.Close()

	var paiements []model.Paiement
	for rows.Next() {
		var p model.Paiement
		err := rows.Scan(&p.ID, &p.IDAnnonce, &p.IDService, &p.IDPayeur, &p.IDReceveur, &p.Montant, &p.DatePaiement, &p.CreatedAt)
		if err != nil {
			log.Printf("Erreur lors du scan du paiement: %v", err)
			return nil, err
		}
		paiements = append(paiements, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur après l'itération des paiements: %v", err)
		return nil, err
	}

	return paiements, nil
}

func GetPaiementByID(db *sql.DB, id int) (*model.Paiement, error) {
	row := db.QueryRow("SELECT id, id_annonce, id_service, id_payeur, id_receveur, montant, date_paiement, created_at FROM paiement WHERE id = ?", id)
	
	var p model.Paiement
	err := row.Scan(&p.ID, &p.IDAnnonce, &p.IDService, &p.IDPayeur, &p.IDReceveur, &p.Montant, &p.DatePaiement, &p.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Paiement avec l'ID %d non trouvé", id)
			return nil, nil // Paiement non trouvé
		}
		log.Printf("Erreur lors de la récupération du paiement par ID: %v", err)
		return nil, err
	}

	return &p, nil
}

func GetPaiementsByIDPayeur(db *sql.DB, idPayeur int) ([]model.Paiement, error) {
	rows, err := db.Query("SELECT id, id_annonce, id_service, id_payeur, id_receveur, montant, date_paiement, created_at FROM paiement WHERE id_payeur = ?", idPayeur)
	if err != nil {
		log.Printf("Erreur lors de la récupération des paiements par ID payeur: %v", err)
		return nil, err
	}
	defer rows.Close()

	var paiements []model.Paiement
	for rows.Next() {
		var p model.Paiement
		err := rows.Scan(&p.ID, &p.IDAnnonce, &p.IDService, &p.IDPayeur, &p.IDReceveur, &p.Montant, &p.DatePaiement, &p.CreatedAt)
		if err != nil {
			log.Printf("Erreur lors du scan du paiement: %v", err)
			return nil, err
		}
		paiements = append(paiements, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur après l'itération des paiements: %v", err)
		return nil, err
	}

	return paiements, nil
}

func GetPaiementsByIDReceveur(db *sql.DB, idReceveur int) ([]model.Paiement, error) {
	rows, err := db.Query("SELECT id, id_annonce, id_service, id_payeur, id_receveur, montant, date_paiement, created_at FROM paiement WHERE id_receveur = ?", idReceveur)
	if err != nil {
		log.Printf("Erreur lors de la récupération des paiements par ID receveur: %v", err)
		return nil, err
	}
	defer rows.Close()

	var paiements []model.Paiement
	for rows.Next() {
		var p model.Paiement
		err := rows.Scan(&p.ID, &p.IDAnnonce, &p.IDService, &p.IDPayeur, &p.IDReceveur, &p.Montant, &p.DatePaiement, &p.CreatedAt)
		if err != nil {
			log.Printf("Erreur lors du scan du paiement: %v", err)
			return nil, err
		}
		paiements = append(paiements, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur après l'itération des paiements: %v", err)
		return nil, err
	}

	return paiements, nil
}

func CreatePaiement(db *sql.DB, p *model.Paiement) error {
	_, err := db.Exec("INSERT INTO paiement (id_annonce, id_service, id_payeur, id_receveur, montant, date_paiement) VALUES (?, ?, ?, ?, ?, ?)", p.IDAnnonce, p.IDService, p.IDPayeur, p.IDReceveur, p.Montant, p.DatePaiement)
	if err != nil {
		log.Printf("Erreur lors de la création du paiement: %v", err)
		return err
	}
	return nil
}