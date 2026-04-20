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

func GetProfessionnelsHandler(db *sql.DB) http.Handler {
	return middleware.GetProfessionnelMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		professionnels, err := repositories.GetAllProfessionnels(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des professionnels", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(professionnels)
	}))
}

func GetProfessionnelByIDHandler(db *sql.DB) http.Handler {
	return middleware.GetProfessionnelByIDMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		professionnel, err := repositories.GetProfessionnelByID(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du professionnel", http.StatusInternalServerError)
			return
		}
		if professionnel == nil {
			http.Error(w, "Professionnel non trouvé", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(professionnel)
	}))
}

func GetProfessionnelByEmailHandler(db *sql.DB) http.Handler {
	return middleware.GetProfessionnelByEmailMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		professionnel, err := repositories.GetProfessionnelByEmail(db, email)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du professionnel", http.StatusInternalServerError)
			return
		}
		if professionnel == nil {
			http.Error(w, "Professionnel non trouvé", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(professionnel)
	}))
}

func CreateProfessionnelHandler(db *sql.DB) http.Handler {
	return middleware.CreateProfessionnelMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var p model.Professionnel
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			return
		}
		if err := repositories.CreateProfessionnel(db, &p); err != nil {
			http.Error(w, "Erreur lors de la création du professionnel", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(p)
	}))
}

func UpdateProfessionnelHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		var p model.Professionnel
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			return
		}
		p.ID = id
		if err := repositories.UpdateProfessionnel(db, &p); err != nil {
			http.Error(w, "Erreur lors de la mise à jour du professionnel", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(p)
	})
}

func DeleteProfessionnelHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		professionnel, err := repositories.GetProfessionnelByID(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du professionnel", http.StatusInternalServerError)
			return
		}
		if professionnel == nil {
			http.Error(w, "Professionnel non trouvé", http.StatusNotFound)
			return
		}
		if err := repositories.DeleteProfessionnel(db, id); err != nil {
			http.Error(w, "Erreur lors de la suppression du professionnel", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}