package middleware

import (
	"database/sql"
	"net/http"
	"API/repositories"
)

func GetChiffreAffairesMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chiffreAffaire, err := repositories.GetChiffreAffaire(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du chiffre d'affaires", http.StatusInternalServerError)
			return
		}
		if chiffreAffaire == 0 {
			http.Error(w, "Aucun chiffre d'affaires trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetBeneficeMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		benefice, err := repositories.GetBenefice(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du bénéfice", http.StatusInternalServerError)
			return
		}
		if benefice == 0 {
			http.Error(w, "Aucun bénéfice trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetTauxBeneficeMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tauxBenefice, err := repositories.GetTauxBenefice(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du taux de bénéfice", http.StatusInternalServerError)
			return
		}
		if tauxBenefice == 0 {
			http.Error(w, "Aucun taux de bénéfice trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetChargeMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		charges, err := repositories.GetCharges(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des charges", http.StatusInternalServerError)
			return
		}
		if charges == 0 {
			http.Error(w, "Aucune charge trouvée", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}