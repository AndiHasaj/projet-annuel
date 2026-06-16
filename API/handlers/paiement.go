package handlers

import (
	"API/model"
	"API/repositories"
	"database/sql"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetPaiementsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		paiements, err := repositories.GetAllPaiements(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des paiements"})
			return
		}
		c.JSON(http.StatusOK, paiements)
	}
}

func GetPaiementByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
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

func GetPaiementsByIDPayeurHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id_payeur, err := strconv.Atoi(c.Query("id_payeur"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du payeur invalide"})
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

func GetPaiementsByIDReceveurHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id_receveur, err := strconv.Atoi(c.Query("id_receveur"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du receveur invalide"})
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

func CreatePaiementHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p model.Paiement
		
		body, exists := c.Get("body")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Corps de la requête manquant"})
			return
		}

		data := body.(map[string]string)
		id_annonce, err := strconv.Atoi(data["id_annonce"])
		if err != nil {
			id_annonce = -1
		}

		id_service, err := strconv.Atoi(data["id_service"])
		if err != nil {
			id_service = -1
		}

		id_payeur, err := strconv.Atoi(data["id_payeur"])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du payeur invalide"})
			return
		}

		id_receveur, err := strconv.Atoi(data["id_receveur"])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du receveur invalide"})
			return
		}

		montant, err := strconv.ParseFloat(data["montant"], 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Montant invalide"})
			return
		}

		p = model.Paiement{
			IDAnnonce:    id_annonce,
			IDService:    id_service,
			IDPayeur:     id_payeur,
			IDReceveur:   id_receveur,
			Montant:      montant,
			TypePaiement: data["type_paiement"],
			Statut:       data["statut"],
		}

		_, err = repositories.CreatePaiement(db, &p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création du paiement"})
			return
		}
		c.JSON(http.StatusCreated, p)
	}
}