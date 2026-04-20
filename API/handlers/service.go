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

func GetServicesHandler(db *sql.DB) http.Handler {
	return middleware.GetServiceMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		services, err := repositories.GetAllServices(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des services", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(services)
	}))
}

func GetServiceByIDHandler(db *sql.DB) http.Handler {
	return middleware.GetServiceByIDMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
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
		json.NewEncoder(w).Encode(service)
	}))
}

func GetServicesByTypeHandler(db *sql.DB) http.Handler {
	return middleware.GetServicesByTypeMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		json.NewEncoder(w).Encode(services)
	}))
}

func CreateServiceHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var service model.Service
		if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			return
		}
		if err := repositories.CreateService(db, &service); err != nil {
			http.Error(w, "Erreur lors de la création du service", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(service)
	})
}

func UpdateServiceHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		var service model.Service
		if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			return
		}
		service.ID = id
		if err := repositories.UpdateService(db, &service); err != nil {
			http.Error(w, "Erreur lors de la mise à jour du service", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(service)
	})
}

func DeleteServiceHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
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
		if err := repositories.DeleteService(db, id); err != nil {
			http.Error(w, "Erreur lors de la suppression du service", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(service)
	})
}