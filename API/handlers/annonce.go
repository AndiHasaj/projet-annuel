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

func GetAnnoncesHandler(db *sql.DB) http.Handler {
	return middleware.GetAnnonceMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		annonces, err := repositories.GetAllAnnonces(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des annonces", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(annonces)
	}))
}

func GetAnnonceByIDHandler(db *sql.DB) http.Handler {
	return middleware.GetAnnonceByIDMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		annonce, err := repositories.GetAnnonceByID(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération de l'annonce", http.StatusInternalServerError)
			return
		}
		if annonce == nil {
			http.Error(w, "Annonce non trouvée", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(annonce)
	}))
}

func GetAnnoncesByIDAnnonceurHandler(db *sql.DB) http.Handler {
	return middleware.GetAnnoncesByIDAnnonceurMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idAnnonceur, err := strconv.Atoi(r.URL.Query().Get("id_annonceur"))
		if err != nil {
			http.Error(w, "ID de l'annonceur invalide", http.StatusBadRequest)
			return
		}
		annonces, err := repositories.GetAnnoncesByIDAnnonceur(db, idAnnonceur)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des annonces de l'annonceur", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(annonces)
	}))
}

func CreateAnnonceHandler(db *sql.DB) http.Handler {
	return middleware.CreateAnnonceMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var annonce model.Annonce
		if err := json.NewDecoder(r.Body).Decode(&annonce); err != nil {
			http.Error(w, "Données d'annonce invalides", http.StatusBadRequest)
			return
		}
		if err := repositories.CreateAnnonce(db, &annonce); err != nil {
			http.Error(w, "Erreur lors de la création de l'annonce", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(annonce)
	}))
}

func UpdateAnnonceHandler(db *sql.DB) http.Handler {
	return middleware.UpdateAnnonceMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var annonce model.Annonce
		if err := json.NewDecoder(r.Body).Decode(&annonce); err != nil {
			http.Error(w, "Données d'annonce invalides", http.StatusBadRequest)
			return
		}
		if err := repositories.UpdateAnnonce(db, &annonce); err != nil {
			http.Error(w, "Erreur lors de la mise à jour de l'annonce", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(annonce)
	}))
}

func DeleteAnnonceHandler(db *sql.DB) http.Handler {
	return middleware.DeleteAnnonceMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		err = repositories.DeleteAnnonce(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la suppression de l'annonce", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}))
}