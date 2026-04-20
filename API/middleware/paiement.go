package middleware

import (
	"database/sql"
	"net/http"
	"API/repositories"
	"strconv"
)

func GetPaiementMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paiements, err := repositories.GetAllPaiements(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des paiements", http.StatusInternalServerError)
			return
		}
		if len(paiements) == 0 {
			http.Error(w, "Aucun paiement trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetPaiementByIDMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID du paiement manquant", http.StatusBadRequest)
			return
		}
		paiement, err := repositories.GetPaiementByID(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du paiement", http.StatusInternalServerError)
			return
		}
		if paiement == nil {
			http.Error(w, "Paiement non trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetPaiementsByIDPayeurMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id_payeur, err := strconv.Atoi(r.URL.Query().Get("id_payeur"))
		if err != nil {
			http.Error(w, "ID du payeur manquant", http.StatusBadRequest)
			return
		}
		paiements, err := repositories.GetPaiementsByIDPayeur(db, id_payeur)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des paiements", http.StatusInternalServerError)
			return
		}
		if len(paiements) == 0 {
			http.Error(w, "Aucun paiement trouvé pour ce payeur", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetPaiementsByIDReceveurMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id_receveur, err := strconv.Atoi(r.URL.Query().Get("id_receveur"))
		if err != nil {
			http.Error(w, "ID du receveur manquant", http.StatusBadRequest)
			return
		}
		paiements, err := repositories.GetPaiementsByIDReceveur(db, id_receveur)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des paiements", http.StatusInternalServerError)
			return
		}
		if len(paiements) == 0 {
			http.Error(w, "Aucun paiement trouvé pour ce receveur", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func CreatePaiementMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idAnnonce, err := strconv.Atoi(r.FormValue("id_annonce"))
		if err != nil {
			idAnnonce = 0
		}
		idService, err := strconv.Atoi(r.FormValue("id_service"))
		if err != nil {
			idService = 0
		}
		if idAnnonce == 0 && idService == 0 {
			http.Error(w, "ID de l'annonce ou du service manquant", http.StatusBadRequest)
			return
		}
		if idAnnonce != 0 {
			r.Form.Set("id_annonce", strconv.Itoa(idAnnonce))
		}
		if idService != 0 {
			r.Form.Set("id_service", strconv.Itoa(idService))
		}
		id_payeur, err := strconv.Atoi(r.FormValue("id_payeur"))
		if err != nil {
			http.Error(w, "ID du payeur manquant", http.StatusBadRequest)
			return
		}
		r.Form.Set("id_payeur", strconv.Itoa(id_payeur))
		id_receveur, err := strconv.Atoi(r.FormValue("id_receveur"))
		if err != nil {
			http.Error(w, "ID du receveur manquant", http.StatusBadRequest)
			return
		}
		r.Form.Set("id_receveur", strconv.Itoa(id_receveur))
		montant, err := strconv.ParseFloat(r.FormValue("montant"), 64)
		if err != nil {
			http.Error(w, "Montant invalide", http.StatusBadRequest)
			return
		}
		r.Form.Set("montant", strconv.FormatFloat(montant, 'f', 2, 64))
		date_paiement := r.FormValue("date_paiement")
		if date_paiement == "" {
			http.Error(w, "Date de paiement manquante", http.StatusBadRequest)
			return
		}
		r.Form.Set("date_paiement", date_paiement)
		type_paiement := r.FormValue("type_paiement")
		if type_paiement == "" {
			http.Error(w, "Type de paiement manquant", http.StatusBadRequest)
			return
		}
		r.Form.Set("type_paiement", type_paiement)
		next.ServeHTTP(w, r)
	})
}