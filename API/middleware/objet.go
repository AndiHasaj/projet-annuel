package middleware

import (
	"database/sql"
	"net/http"
	"API/repositories"
	"strconv"
)

func GetObjetMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		objets, err := repositories.GetAllObjets(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des objets", http.StatusInternalServerError)
			return
		}
		if len(objets) == 0 {
			http.Error(w, "Aucun objet trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetObjetByIDMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID de l'objet manquant", http.StatusBadRequest)
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
		next.ServeHTTP(w, r)
	})
}

func GetObjetsByConteneurIDMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conteneurID, err := strconv.Atoi(r.URL.Query().Get("conteneur_id"))
		if err != nil {
			http.Error(w, "ID du conteneur manquant", http.StatusBadRequest)
			return
		}
		objets, err := repositories.GetObjetsByIDConteneur(db, conteneurID)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des objets", http.StatusInternalServerError)
			return
		}
		if len(objets) == 0 {
			http.Error(w, "Aucun objet trouvé pour ce conteneur", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func CreateObjetMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nom:= r.FormValue("nom")
		if nom == "" {
			http.Error(w, "Nom de l'objet manquant", http.StatusBadRequest)
			return
		}
		r.Form.Set("nom", nom)
		type_objet := r.FormValue("type_objet")
		if type_objet == "" {
			http.Error(w, "Type de l'objet manquant", http.StatusBadRequest)
			return
		}
		r.Form.Set("type_objet", type_objet)
		localisation := r.FormValue("localisation")
		if localisation == "" {
			http.Error(w, "Localisation de l'objet manquante", http.StatusBadRequest)
			return
		}
		r.Form.Set("localisation", localisation)
		next.ServeHTTP(w, r)
	})
}

func UpdateObjetMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID de l'objet manquant", http.StatusBadRequest)
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
		next.ServeHTTP(w, r)
	})
}

func DeleteObjetMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID de l'objet manquant", http.StatusBadRequest)
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
		next.ServeHTTP(w, r)
	})
}