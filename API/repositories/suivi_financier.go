package repositories

import (
	"database/sql"
	"log"
)

func GetChiffreAffaire(db *sql.DB) (float64, error) {
	row := db.QueryRow("SELECT chiffre_affaire FROM suivifinancier")
	var chiffreAffaire float64
	err := row.Scan(&chiffreAffaire)
	if err != nil {
		log.Printf("Erreur lors de la récupération du chiffre d'affaires: %v", err)
		return 0, err
	}
	return chiffreAffaire, nil
}

func GetCharges(db *sql.DB) (float64, error) {
	row := db.QueryRow("SELECT charges FROM suivifinancier")
	var charges float64
	err := row.Scan(&charges)
	if err != nil {
		log.Printf("Erreur lors de la récupération des charges: %v", err)
		return 0, err
	}
	return charges, nil
}

func GetBenefice(db *sql.DB) (float64, error) {
	row := db.QueryRow("SELECT benefice FROM suivifinancier")
	var benefice float64
	err := row.Scan(&benefice)
	if err != nil {
		log.Printf("Erreur lors de la récupération du bénéfice: %v", err)
		return 0, err
	}
	return benefice, nil
}

func GetTauxBenefice(db *sql.DB) (float64, error) {
	chiffreAffaire, err := GetChiffreAffaire(db)
	if err != nil {
		return 0, err
	}
	charges, err := GetCharges(db)
	if err != nil {
		return 0, err
	}
	if chiffreAffaire == 0 {
		return 0, nil // Éviter la division par zéro
	}
	tauxBenefice := (chiffreAffaire - charges) / chiffreAffaire * 100
	return tauxBenefice, nil
}

func UpdateSuiviFinancier(db *sql.DB, chiffreAffaire, charges, benefice float64) error {
	_, err := db.Exec("UPDATE suivifinancier SET chiffre_affaire = ?, charges = ?, benefice = ?", chiffreAffaire, charges, benefice)
	if err != nil {
		log.Printf("Erreur lors de la mise à jour du suivi financier: %v", err)
		return err
	}
	return nil
}