package middleware

import (
	"database/sql"
	"net/http"
	"API/repositories"
	"strconv"
	"golang.org/x/crypto/bcrypt"
)

func GetSalarieMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		salaries, err := repositories.GetAllSalaries(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des salariés", http.StatusInternalServerError)
			return
		}
		if len(salaries) == 0 {
			http.Error(w, "Aucun salarié trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetSalarieByIDMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID du salarié manquant", http.StatusBadRequest)
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
		next.ServeHTTP(w, r)
	})
}

func GetSalarieByEmailMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		if email == "" {
			http.Error(w, "Email du salarié manquant", http.StatusBadRequest)
			return
		}
		salarie, err := repositories.GetSalarieByEmail(db, email)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du salarié", http.StatusInternalServerError)
			return
		}
		if salarie == nil {
			http.Error(w, "Salarié non trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func CreateSalarieMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nom, err := r.FormValue("nom"), error(nil)
		if nom == "" {
			http.Error(w, "Nom du salarié manquant", http.StatusBadRequest)
			return
		}
		r.Form.Set("nom", nom)
		prenom, err := r.FormValue("prenom"), error(nil)
		if prenom == "" {
			http.Error(w, "Prénom du salarié manquant", http.StatusBadRequest)
			return
		}
		r.Form.Set("prenom", prenom)
		email := prenom + "." + nom + "@upcycle.fr"
		salarie, err := repositories.GetSalarieByEmail(db, email)
		if err != nil {
			http.Error(w, "Erreur lors de la vérification de l'email du salarié", http.StatusInternalServerError)
			return
		}
		if salarie != nil {
			http.Error(w, "Un salarié avec cet email existe déjà", http.StatusConflict)
			return
		}
		r.Form.Set("email", email)
		mdp, err := r.FormValue("mot_de_passe"), error(nil)
		if mdp == "" {
			http.Error(w, "Mot de passe du particulier manquant", http.StatusBadRequest)
			return
		}

		
		if len(mdp) < 8 {
			http.Error(w, "Mot de passe du particulier invalide", http.StatusBadRequest)
			return
		}

		hasUpper := false
		hasLower := false
		hasDigit := false
		hasSpecial := false
		for _, c := range mdp {
			switch {
			case 'A' <= c && c <= 'Z':
				hasUpper = true
			case 'a' <= c && c <= 'z':
				hasLower = true
			case '0' <= c && c <= '9':
				hasDigit = true
			case (c >= 33 && c <= 47) || (c >= 58 && c <= 64) || (c >= 91 && c <= 96) || (c >= 123 && c <= 126):
				hasSpecial = true
			}
		}
		if !hasUpper || !hasLower || !hasDigit || !hasSpecial {
			http.Error(w, "Mot de passe du particulier invalide", http.StatusBadRequest)
			return
		}
		mdpHash, err := bcrypt.GenerateFromPassword([]byte(mdp), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Erreur lors du hachage du mot de passe du particulier", http.StatusInternalServerError)
			return
		}
		r.Form.Set("mot_de_passe", string(mdpHash))

		intitulePoste, err := r.FormValue("intitule_poste"), error(nil)
		if intitulePoste == "" {
			http.Error(w, "Intitulé du poste du salarié manquant", http.StatusBadRequest)
			return
		}
		r.Form.Set("intitule_poste", intitulePoste)

		typeContrat, err := r.FormValue("type_contrat"), error(nil)
		if typeContrat == "" {
			http.Error(w, "Type de contrat du salarié manquant", http.StatusBadRequest)
			return
		}
		r.Form.Set("type_contrat", typeContrat)

		salaire, err := r.FormValue("salaire"), error(nil)
		if salaire == "" {
			http.Error(w, "Salaire du salarié manquant", http.StatusBadRequest)
			return
		}
		r.Form.Set("salaire", salaire)

		dateEmbauche, err := r.FormValue("date_embauche"), error(nil)
		if dateEmbauche == "" {
			http.Error(w, "Date d'embauche du salarié manquante", http.StatusBadRequest)
			return
		}
		r.Form.Set("date_embauche", dateEmbauche)
		next.ServeHTTP(w, r)
	})
}

func UpdateSalarieMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID du salarié manquant", http.StatusBadRequest)
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
		next.ServeHTTP(w, r)
	})
}

func DeleteSalarieMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID du salarié manquant", http.StatusBadRequest)
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
		next.ServeHTTP(w, r)
	})
}
