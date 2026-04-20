package middleware

import (
	"database/sql"
	"net/http"
	"API/repositories"
	"strconv"
	"golang.org/x/crypto/bcrypt"
)

func GetAdminMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		admins, err := repositories.GetAllAdmins(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des administrateurs", http.StatusInternalServerError)
			return
		}
		if len(admins) == 0 {
			http.Error(w, "Aucun administrateur trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetAdminByIDMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID de l'administrateur manquant", http.StatusBadRequest)
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
		next.ServeHTTP(w, r)
	})
}

func CreateAdminMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mdp, err := r.FormValue("mot_de_passe"), error(nil)
		if mdp == "" {
			http.Error(w, "Mot de passe de l'administrateur manquant", http.StatusBadRequest)
			return
		}
		if len(mdp) < 8 {
			http.Error(w, "Mot de passe de l'administrateur invalide", http.StatusBadRequest)
			return
		}
		hasUpper, hasLower, hasDigit, hasSpecial := false, false, false, false
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
			http.Error(w, "Mot de passe de l'administrateur invalide", http.StatusBadRequest)
			return
		}
		mdpHash, err := bcrypt.GenerateFromPassword([]byte(mdp), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Erreur lors du hachage du mot de passe de l'administrateur", http.StatusInternalServerError)
			return
		}
		r.Form.Set("mot_de_passe", string(mdpHash))
		next.ServeHTTP(w, r)
	})
}

func DeleteAdminMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID de l'administrateur manquant", http.StatusBadRequest)
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
		next.ServeHTTP(w, r)
	})
}