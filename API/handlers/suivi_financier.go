package handlers

import (
	"API/repositories"
	"database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetChiffreAffairesHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chiffreAffaires, err := repositories.GetChiffreAffaire(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du chiffre d'affaires"})
			return
		}
		c.JSON(http.StatusOK, chiffreAffaires)
	}
}

func GetChargeHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		charges, err := repositories.GetCharges(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des charges"})
			return
		}
		c.JSON(http.StatusOK, charges)
	}
}

func GetBeneficeHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		benefice, err := repositories.GetBenefice(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du bénéfice"})
			return
		}
		c.JSON(http.StatusOK, benefice)
	}
}

func GetTauxBeneficeHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tauxBenefice, err := repositories.GetTauxBenefice(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du taux de bénéfice"})
			return
		}
		c.JSON(http.StatusOK, tauxBenefice)
	}
}