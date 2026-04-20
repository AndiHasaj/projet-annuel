package middleware

import (
	"database/sql"
	"net/http"
	"API/repositories"
	"strconv"
	"golang.org/x/crypto/bcrypt"
)

func GetProfessionnelMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		professionnels, err := repositories.GetAllProfessionnels(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des professionnels", http.StatusInternalServerError)
			return
		}
		if len(professionnels) == 0 {
			http.Error(w, "Aucun professionnel trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetProfessionnelByIDMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID du professionnel manquant", http.StatusBadRequest)
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
		next.ServeHTTP(w, r)
	})
}

func GetProfessionnelByEmailMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		if email == "" {
			http.Error(w, "Email du professionnel manquant", http.StatusBadRequest)
			return
		}
		professionnel, err := repositories.GetProfessionnelByEmail(db, email)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du professionnel", http.StatusInternalServerError)
			return
		}
		if professionnel == nil {
			http.Error(w, "Professionnel non trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func CreateProfessionnelMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		if email == "" {
			http.Error(w, "Email du professionnel manquant", http.StatusBadRequest)
			return
		}
		professionnel, err := repositories.GetProfessionnelByEmail(db, email)
		if err != nil {
			http.Error(w, "Erreur lors de la vérification de l'email du professionnel", http.StatusInternalServerError)
			return
		}
		if professionnel != nil {
			http.Error(w, "Un professionnel avec cet email existe déjà", http.StatusConflict)
			return
		}
		mdp, err := r.FormValue("mot_de_passe"), error(nil)
		if mdp == "" {
			http.Error(w, "Mot de passe du professionnel manquant", http.StatusBadRequest)
			return
		}

		
		if len(mdp) < 8 {
			http.Error(w, "Mot de passe du professionnel invalide", http.StatusBadRequest)
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
			http.Error(w, "Mot de passe du professionnel invalide", http.StatusBadRequest)
			return
		}
		mdpHash, err := bcrypt.GenerateFromPassword([]byte(mdp), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Erreur lors du hachage du mot de passe du professionnel", http.StatusInternalServerError)
			return
		}
		r.Form.Set("mot_de_passe", string(mdpHash))

		numeroSiret := r.FormValue("numero_siret")
		if numeroSiret == "" {
			http.Error(w, "Numéro de SIRET du professionnel manquant", http.StatusBadRequest)
			return
		}
		if len(numeroSiret) != 14 {
			http.Error(w, "Numéro de SIRET du professionnel invalide", http.StatusBadRequest)
			return
		}
		for _, c := range numeroSiret {
			if c < '0' || c > '9' {
				http.Error(w, "Numéro de SIRET du professionnel invalide", http.StatusBadRequest)
				return
			}
		}
		// Possibilité de faire la vérif par api.insee.fr mais pas de clé d'api pour le moment
		r.Form.Set("numero_siret", numeroSiret)

		telephone, err := r.FormValue("telephone"), error(nil)
		if telephone == "" {
			http.Error(w, "Téléphone du professionnel manquant", http.StatusBadRequest)
			return
		}
		if len(telephone) != 10 {
			http.Error(w, "Téléphone du professionnel invalide", http.StatusBadRequest)
			return
		}
		for _, c := range telephone {
			if c < '0' || c > '9' {
				http.Error(w, "Téléphone du professionnel invalide", http.StatusBadRequest)
				return
			}
		}
		r.Form.Set("telephone", telephone)

		site_web := r.FormValue("site_web")
		if site_web == "" {
			http.Error(w, "Site web du professionnel manquant", http.StatusBadRequest)
			return
		}
		r.Form.Set("site_web", site_web)

		adresse := r.FormValue("adresse")
		if adresse == "" {
			http.Error(w, "Adresse du professionnel manquante", http.StatusBadRequest)
			return
		}
		r.Form.Set("adresse", adresse)

		code_postal := r.FormValue("code_postal")
		if code_postal == "" {
			http.Error(w, "Code postal du professionnel manquant", http.StatusBadRequest)
			return
		}
		if len(code_postal) != 5 {
			http.Error(w, "Code postal du professionnel invalide", http.StatusBadRequest)
			return
		}
		for _, c := range code_postal {
			if c < '0' || c > '9' {
				http.Error(w, "Code postal du professionnel invalide", http.StatusBadRequest)
				return
			}
		}
		r.Form.Set("code_postal", code_postal)

		ville, err := r.FormValue("ville"), error(nil)
		if ville == "" {
			http.Error(w, "Ville du professionnel manquante", http.StatusBadRequest)
			return
		}
		r.Form.Set("ville", ville)
		next.ServeHTTP(w, r)
	})
}

func UpdateProfessionnelMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID du professionnel manquant", http.StatusBadRequest)
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
		next.ServeHTTP(w, r)
	})
}

func DeleteProfessionnelMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID du professionnel manquant", http.StatusBadRequest)
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
		next.ServeHTTP(w, r)
	})
}