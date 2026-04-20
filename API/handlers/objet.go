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

func GetObjetsHandler(db *sql.DB) http.Handler {
	return middleware.GetObjetMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		objets, err := repositories.GetAllObjets(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des objets", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(objets)
	}))
}

func GetObjetByIDHandler(db *sql.DB) http.Handler {
	return middleware.GetObjetByIDMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		objet, err := repositories.GetObjetByID(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération de l'objet", http.StatusInternalServerError)
			return
		}
		if objet == nil {
			http.Error(w, "Objet non trouvé", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(objet)
	}))
}

func GetObjetsByIDConteneurHandler(db *sql.DB) http.Handler {
	return middleware.GetObjetsByConteneurIDMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idConteneur, err := strconv.Atoi(r.URL.Query().Get("id_conteneur"))
		if err != nil {
			http.Error(w, "ID du conteneur invalide", http.StatusBadRequest)
			return
		}
		objets, err := repositories.GetObjetsByIDConteneur(db, idConteneur)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des objets du conteneur", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(objets)
	}))
}

func CreateObjetHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var objet model.Objet
		if err := json.NewDecoder(r.Body).Decode(&objet); err != nil {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			return
		}
		if err := repositories.CreateObjet(db, &objet); err != nil {
			http.Error(w, "Erreur lors de la création de l'objet", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(objet)
	})
}

func UpdateObjetHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		var objet model.Objet
		if err := json.NewDecoder(r.Body).Decode(&objet); err != nil {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			return
		}
		objet.ID = id
		if err := repositories.UpdateObjet(db, &objet); err != nil {
			http.Error(w, "Erreur lors de la mise à jour de l'objet", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(objet)
	})
}

func DeleteObjetHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		objet, err := repositories.GetObjetByID(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération de l'objet", http.StatusInternalServerError)
			return
		}
		if objet == nil {
			http.Error(w, "Objet non trouvé", http.StatusNotFound)
			return
		}
		if err := repositories.DeleteObjet(db, id); err != nil {
			http.Error(w, "Erreur lors de la suppression de l'objet", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}