package handlers

import (
	"API/middleware"
	"API/repositories"
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetChiffreAffairesHandler(db *sql.DB) http.Handler {
	return middleware.GetChiffreAffairesMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chiffreAffaires, err := repositories.GetChiffreAffaire(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du chiffre d'affaires", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(chiffreAffaires)
	}))
}

func GetChargeHandler(db *sql.DB) http.Handler {
	return middleware.GetChargeMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		charges, err := repositories.GetCharges(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des charges", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(charges)
	}))
}

func GetBeneficeHandler(db *sql.DB) http.Handler {
	return middleware.GetBeneficeMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		benefice, err := repositories.GetBenefice(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du bénéfice", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(benefice)
	}))
}

func GetTauxBeneficeHandler(db *sql.DB) http.Handler {
	return middleware.GetTauxBeneficeMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tauxBenefice, err := repositories.GetTauxBenefice(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du taux de bénéfice", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(tauxBenefice)
	}))
}