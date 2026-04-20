package middleware

import (
	"database/sql"
	"net/http"
	"API/repositories"
	"strconv"
	"time"
)

func GetAnnonceMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		annonces, err := repositories.GetAllAnnonces(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des annonces", http.StatusInternalServerError)
			return
		}
		if len(annonces) == 0 {
			http.Error(w, "Aucune annonce trouvée", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetAnnonceByIDMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID de l'annonce manquant", http.StatusBadRequest)
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
		next.ServeHTTP(w, r)
	})
}

func GetAnnoncesByIDAnnonceurMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idAnnonceur, err := strconv.Atoi(r.URL.Query().Get("id_annonceur"))
		if err != nil {
			http.Error(w, "ID de l'annonceur manquant", http.StatusBadRequest)
			return
		}
		annonces, err := repositories.GetAnnoncesByIDAnnonceur(db, idAnnonceur)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des annonces", http.StatusInternalServerError)
			return
		}
		if len(annonces) == 0 {
			http.Error(w, "Aucune annonce trouvée pour ce particulier", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func CreateAnnonceMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		titre := r.FormValue("titre")
		if titre == "" {
			http.Error(w, "Titre de l'annonce manquant", http.StatusBadRequest)
			return
		}
		r.Form.Set("titre", titre)
		description := r.FormValue("description")
		if description == "" {
			http.Error(w, "Description de l'annonce manquante", http.StatusBadRequest)
			return
		}
		r.Form.Set("description", description)
		prix, err := strconv.ParseFloat(r.FormValue("prix"), 64)
		if err != nil {
			http.Error(w, "Prix de l'annonce manquant ou invalide", http.StatusBadRequest)
			return
		}
		r.Form.Set("prix", strconv.FormatFloat(prix, 'f', -1, 64))
		particulierID, err := strconv.Atoi(r.FormValue("particulier_id"))
		if err != nil {
			http.Error(w, "ID du particulier manquant ou invalide", http.StatusBadRequest)
			return
		}
		r.Form.Set("particulier_id", strconv.Itoa(particulierID))
		r.Form.Set("date_publication", time.Now().Format(time.RFC3339))
		next.ServeHTTP(w, r)
	})
}

func UpdateAnnonceMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID de l'annonce manquant", http.StatusBadRequest)
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
		next.ServeHTTP(w, r)
	})
}

func DeleteAnnonceMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID de l'annonce manquant", http.StatusBadRequest)
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
		next.ServeHTTP(w, r)
	})
}