package middleware

import (
	"database/sql"
	"net/http"
	"API/repositories"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetPaiementMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c*gin.Context) {
		paiements, err := repositories.GetAllPaiements(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des paiements"})
			return
		}
		if len(paiements) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucun paiement trouvé"})
			return
		}
		c.JSON(http.StatusOK, paiements)
	}
}

func GetPaiementByIDMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c*gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du paiement manquant"})
			return
		}
		paiement, err := repositories.GetPaiementByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du paiement"})
			return
		}
		if paiement == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Paiement non trouvé"})
			return
		}
		c.JSON(http.StatusOK, paiement)
	}
}

func GetPaiementsByIDPayeurMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c*gin.Context) {
		id_payeur, err := strconv.Atoi(c.Query("id_payeur"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du payeur manquant"})
			return
		}
		paiements, err := repositories.GetPaiementsByIDPayeur(db, id_payeur)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des paiements"})
			return
		}
		if len(paiements) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucun paiement trouvé pour ce payeur"})
			return
		}
		c.JSON(http.StatusOK, paiements)
	}
}

func GetPaiementsByIDReceveurMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c*gin.Context) {
		id_receveur, err := strconv.Atoi(c.Query("id_receveur"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du receveur manquant"})
			return
		}
		paiements, err := repositories.GetPaiementsByIDReceveur(db, id_receveur)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des paiements"})
			return
		}
		if len(paiements) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucun paiement trouvé pour ce receveur"})
			return
		}
		c.JSON(http.StatusOK, paiements)
	}
}

func CreatePaiementMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c*gin.Context) {
		var body map[string]string
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Body invalide"})
			return
		}

		id_annonce, err := strconv.Atoi(body["id_annonce"])
		if err != nil {
			id_annonce = 0
		}
		id_service, err := strconv.Atoi(body["id_service"])
		if err != nil {
			id_service = 0
		}
		if id_annonce == 0 && id_service == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de l'annonce ou du service manquant"})
			return
		}
		if id_annonce != 0 {
			body["id_annonce"] = strconv.Itoa(id_annonce)
		}
		if id_service != 0 {
			body["id_service"] = strconv.Itoa(id_service)
		}
		_ , err = strconv.Atoi(body["id_payeur"])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du payeur manquant"})
			return
		}
		_ , err = strconv.Atoi(body["id_receveur"])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du receveur manquant"})
			return
		}
		montant, err := strconv.ParseFloat(body["montant"], 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Montant invalide"})
			return
		}
		body["montant"] = strconv.FormatFloat(montant, 'f', 2, 64)
		date_paiement := body["date_paiement"]
		if date_paiement == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Date de paiement manquante"})
			return
		}
		type_paiement := body["type_paiement"]
		if type_paiement == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Type de paiement manquant"})
			return
		}
		
		c.Set("body", body)
		c.Next()
	}
}