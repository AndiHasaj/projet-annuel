package handlers

import (
	"API/middleware"
	"API/model"
	"API/repositories"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetPaiementsHandler(db *sql.DB) http.Handler {
	return middleware.GetPaiementMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paiements, err := repositories.GetAllPaiements(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des paiements", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(paiements)
	}))
}

func GetPaiementByIDHandler(db *sql.DB) http.Handler {
	return middleware.GetPaiementByIDMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		paiement, err := repositories.GetPaiementByID(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du paiement", http.StatusInternalServerError)
			return
		}
		if paiement == nil {
			http.Error(w, "Paiement non trouvé", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(paiement)
	}))
}

func GetPaiementsByIDPayeurHandler(db *sql.DB) http.Handler {
	return middleware.GetPaiementsByIDPayeurMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id_payeur, err := strconv.Atoi(r.URL.Query().Get("id_payeur"))
		if err != nil {
			http.Error(w, "ID du payeur invalide", http.StatusBadRequest)
			return
		}
		paiements, err := repositories.GetPaiementsByIDPayeur(db, id_payeur)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des paiements", http.StatusInternalServerError)
			return
		}
		if len(paiements) == 0 {
			http.Error(w, "Aucun paiement trouvé pour ce payeur", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(paiements)
	}))
}

func GetPaiementsByIDReceveurHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id_receveur, err := strconv.Atoi(r.URL.Query().Get("id_receveur"))
		if err != nil {
			http.Error(w, "ID du receveur invalide", http.StatusBadRequest)
			return
		}
		paiements, err := repositories.GetPaiementsByIDReceveur(db, id_receveur)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des paiements", http.StatusInternalServerError)
			return
		}
		if len(paiements) == 0 {
			http.Error(w, "Aucun paiement trouvé pour ce receveur", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(paiements)
	})
}

func CreatePaiementHandler(db *sql.DB) http.Handler {
	return middleware.CreatePaiementMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var paiement model.Paiement
		if err := json.NewDecoder(r.Body).Decode(&paiement); err != nil {
			http.Error(w, "Données de paiement invalides", http.StatusBadRequest)
			return
		}
		if err := repositories.CreatePaiement(db, &paiement); err != nil {
			http.Error(w, "Erreur lors de la création du paiement", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(paiement)
	}))
}