package middleware

import (
	"database/sql"
	"net/http"
	"API/repositories"
	"strconv"
)

func GetServiceMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		services, err := repositories.GetAllServices(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des services", http.StatusInternalServerError)
			return
		}
		if len(services) == 0 {
			http.Error(w, "Aucun service trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetServiceByIDMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID du service manquant", http.StatusBadRequest)
			return
		}
		service, err := repositories.GetServiceByID(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du service", http.StatusInternalServerError)
			return
		}
		if service == nil {
			http.Error(w, "Service non trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetServicesByTypeMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type_service := r.URL.Query().Get("type")
		if type_service == "" {
			http.Error(w, "Type de service manquant", http.StatusBadRequest)
			return
		}
		services, err := repositories.GetServicesByType(db, type_service)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des services", http.StatusInternalServerError)
			return
		}
		if len(services) == 0 {
			http.Error(w, "Aucun service trouvé pour ce type", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func CreateServiceMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		titre := r.FormValue("titre")
		if titre == "" {
			http.Error(w, "Titre du service manquant", http.StatusBadRequest)
			return
		}
		r.Form.Set("titre", titre)
		type_service := r.FormValue("type")
		if type_service == "" {
			http.Error(w, "Type de service manquant", http.StatusBadRequest)
			return
		}
		r.Form.Set("type", type_service)
		description := r.FormValue("description")
		if description == "" {
			http.Error(w, "Description du service manquante", http.StatusBadRequest)
			return
		}
		r.Form.Set("description", description)
		prix := r.FormValue("prix")
		if prix == "" {
			http.Error(w, "Prix du service manquant", http.StatusBadRequest)
			return
		}
		r.Form.Set("prix", prix)
		next.ServeHTTP(w, r)
	})
}

func UpdateServiceMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID du service manquant", http.StatusBadRequest)
			return
		}
		service, err := repositories.GetServiceByID(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du service", http.StatusInternalServerError)
			return
		}
		if service == nil {
			http.Error(w, "Service non trouvé", http.StatusNotFound)
			return
		}
	})
}

func DeleteServiceMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID du service manquant", http.StatusBadRequest)
			return
		}
		service, err := repositories.GetServiceByID(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du service", http.StatusInternalServerError)
			return
		}
		if service == nil {
			http.Error(w, "Service non trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}