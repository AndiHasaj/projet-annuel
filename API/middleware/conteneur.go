package middleware

import (
	"database/sql"
	"net/http"
	"API/repositories"
	"strconv"
	"math/rand"
)

func GetConteneurMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conteneurs, err := repositories.GetAllConteneurs(db)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des conteneurs", http.StatusInternalServerError)
			return
		}
		if len(conteneurs) == 0 {
			http.Error(w, "Aucun conteneur trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetConteneurByIDMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID du conteneur manquant", http.StatusBadRequest)
			return
		}
		conteneur, err := repositories.GetConteneurByID(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du conteneur", http.StatusInternalServerError)
			return
		}
		if conteneur == nil {
			http.Error(w, "Conteneur non trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetConteneurByMatriculeMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		matricule := r.URL.Query().Get("matricule")
		if matricule == "" {
			http.Error(w, "Matricule du conteneur manquant", http.StatusBadRequest)
			return
		}
		conteneur, err := repositories.GetConteneurByMatricule(db, matricule)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du conteneur", http.StatusInternalServerError)
			return
		}
		if conteneur == nil {
			http.Error(w, "Conteneur non trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetConteneursByLocalisationIDMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		localisation:= r.URL.Query().Get("localisation")
		if localisation == "" {
			http.Error(w, "Localisation du conteneur manquante", http.StatusBadRequest)
			return
		}
		conteneurs, err := repositories.GetConteneursByLocalisation(db, localisation)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des conteneurs", http.StatusInternalServerError)
			return
		}
		if len(conteneurs) == 0 {
			http.Error(w, "Aucun conteneur trouvé pour cette localisation", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetConteneurByCodeMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "Code du conteneur manquant", http.StatusBadRequest)
			return
		}
		conteneur, err := repositories.GetConteneurByCode(db, code)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du conteneur", http.StatusInternalServerError)
			return
		}
		if conteneur == nil {
			http.Error(w, "Aucun conteneur trouvé pour ce code", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}


func CreateConteneurMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		matricule := r.FormValue("matricule")
		if matricule == "" {
			http.Error(w, "Matricule du conteneur manquant", http.StatusBadRequest)
			return
		}
		r.Form.Set("matricule", matricule)
		localisation := r.FormValue("localisation")
		if localisation == "" {
			http.Error(w, "Localisation du conteneur manquante", http.StatusBadRequest)
			return
		}
		r.Form.Set("localisation", localisation)
		occupation := 0
		r.Form.Set("occupation", strconv.Itoa(occupation))
		code := strconv.Itoa(rand.Intn(1000000))
		existingConteneur, err := repositories.GetConteneurByCode(db, code)
		if err != nil {
			http.Error(w, "Erreur lors de la vérification du code du conteneur", http.StatusInternalServerError)
			return
		}
		for existingConteneur != nil {
			code = strconv.Itoa(rand.Intn(1000000))
			existingConteneur, err = repositories.GetConteneurByCode(db, code)
			if err != nil {
				http.Error(w, "Erreur lors de la vérification du code du conteneur", http.StatusInternalServerError)
				return
			}
		}
		r.Form.Set("code", code)
		next.ServeHTTP(w, r)
	})
}

func UpdateConteneurMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID du conteneur manquant", http.StatusBadRequest)
			return
		}
		conteneur, err := repositories.GetConteneurByID(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du conteneur", http.StatusInternalServerError)
			return
		}
		if conteneur == nil {
			http.Error(w, "Conteneur non trouvé", http.StatusNotFound)
			return
		}
		localisation := r.FormValue("localisation")
		if localisation != "" {
			r.Form.Set("localisation", localisation)
		}
		next.ServeHTTP(w, r)
	})
}

func DeleteConteneurMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "ID du conteneur manquant", http.StatusBadRequest)
			return
		}
		conteneur, err := repositories.GetConteneurByID(db, id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du conteneur", http.StatusInternalServerError)
			return
		}
		if conteneur == nil {
			http.Error(w, "Conteneur non trouvé", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}