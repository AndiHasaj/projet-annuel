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

func GetAdminsHandler(db *sql.DB) http.Handler {
	return middleware.GetAdminMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		admins, err := repositories.GetAllAdmins(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des administrateurs", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(admins)
	}))
}

func GetAdminByIDHandler(db *sql.DB) http.Handler {
	return middleware.GetAdminByIDMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		admin, err := repositories.GetAdminByID(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération de l'administrateur", http.StatusInternalServerError)
			return
		}
		if admin == nil {
			http.Error(w, "Administrateur non trouvé", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(admin)
	}))
}

func CreateAdminHandler(db *sql.DB) http.Handler {
	return middleware.CreateAdminMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var admin model.Admin
		if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			return
		}
		if err := repositories.CreateAdmin(db, &admin); err != nil {
			http.Error(w, "Erreur lors de la création de l'administrateur", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(admin)
	}))
}

func DeleteAdminHandler(db *sql.DB) http.Handler {
	return middleware.DeleteAdminMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		admin, err := repositories.GetAdminByID(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération de l'administrateur", http.StatusInternalServerError)
			return
		}
		if admin == nil {
			http.Error(w, "Administrateur non trouvé", http.StatusNotFound)
			return
		}
		if err := repositories.DeleteAdmin(db, id); err != nil {
			http.Error(w, "Erreur lors de la suppression de l'administrateur", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(admin)
	}))
}