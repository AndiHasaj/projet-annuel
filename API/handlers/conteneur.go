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

func GetConteneursHandler(db *sql.DB) http.Handler {
	return middleware.GetConteneurMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conteneurs, err := repositories.GetAllConteneurs(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des conteneurs", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(conteneurs)
	}))
}

func GetConteneurByIDHandler(db *sql.DB) http.Handler {
	return middleware.GetConteneurByIDMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		conteneur, err := repositories.GetConteneurByID(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du conteneur", http.StatusInternalServerError)
			return
		}
		if conteneur == nil {
			http.Error(w, "Conteneur non trouvé", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(conteneur)
	}))
}

func GetConteneurByMatriculeHandler(db *sql.DB) http.Handler {
	return middleware.GetConteneurByMatriculeMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		matricule := r.URL.Query().Get("matricule")
		conteneur, err := repositories.GetConteneurByMatricule(db, matricule)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du conteneur", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(conteneur)
	}))
}

func GetConteneursByLocalisationHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		localisation := r.URL.Query().Get("localisation")
		conteneurs, err := repositories.GetConteneursByLocalisation(db, localisation)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des conteneurs", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(conteneurs)
	})
}

func GetConteneurByCodeHandler(db *sql.DB) http.Handler {
	return middleware.GetConteneurByMatriculeMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		conteneur, err := repositories.GetConteneurByCode(db, code)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du conteneur", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(conteneur)
	}))
}

func CreateConteneurHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var conteneur model.Conteneur
		if err := json.NewDecoder(r.Body).Decode(&conteneur); err != nil {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			return
		}
		if conteneur.Matricule == ""{
			http.Error(w, "Matricule du conteneur est requis", http.StatusBadRequest)
			return
		}
		if err := repositories.CreateConteneur(db, &conteneur); err != nil {
			http.Error(w, "Erreur lors de la création du conteneur", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(conteneur)
	})
}

func UpdateConteneurHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		var conteneur model.Conteneur
		if err := json.NewDecoder(r.Body).Decode(&conteneur); err != nil {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			return
		}
		if conteneur.Matricule == ""{
			http.Error(w, "Matricule du conteneur est requis", http.StatusBadRequest)
			return
		}
		conteneur.ID = id
		if err := repositories.UpdateConteneur(db, &conteneur); err != nil {
			http.Error(w, "Erreur lors de la mise à jour du conteneur", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(conteneur)
	})
}

func DeleteConteneurHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		if err := repositories.DeleteConteneur(db, id); err != nil {
			http.Error(w, "Erreur lors de la suppression du conteneur", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}