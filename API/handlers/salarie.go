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

func GetSalariesHandler(db *sql.DB) http.Handler {
	return middleware.GetSalarieMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		salaries, err := repositories.GetAllSalaries(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des salariés", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(salaries)
	}))
}

func GetSalarieByIDHandler(db *sql.DB) http.Handler {
	return middleware.GetSalarieByIDMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		salarie, err := repositories.GetSalarieByID(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du salarié", http.StatusInternalServerError)
			return
		}
		if salarie == nil {
			http.Error(w, "Salarié non trouvé", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(salarie)
	}))
}

func GetSalarieByEmailHandler(db *sql.DB) http.Handler {
	return middleware.GetSalarieByEmailMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		salarie, err := repositories.GetSalarieByEmail(db, email)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du salarié", http.StatusInternalServerError)
			return
		}
		if salarie == nil {
			http.Error(w, "Salarié non trouvé", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(salarie)
	}))
}

func CreateSalarieHandler(db *sql.DB) http.Handler {
	return middleware.CreateSalarieMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var salarie model.Salarie
		if err := json.NewDecoder(r.Body).Decode(&salarie); err != nil {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			return
		}
		if err := repositories.CreateSalarie(db, &salarie); err != nil {
			http.Error(w, "Erreur lors de la création du salarié", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(salarie)
	}))
}

func UpdateSalarieHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		var salarie model.Salarie
		if err := json.NewDecoder(r.Body).Decode(&salarie); err != nil {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			return
		}
		salarie.ID = id
		if err := repositories.UpdateSalarie(db, &salarie); err != nil {
			http.Error(w, "Erreur lors de la mise à jour du salarié", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(salarie)
	})
}

func DeleteSalarieHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		if err := repositories.DeleteSalarie(db, id); err != nil {
			http.Error(w, "Erreur lors de la suppression du salarié", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}