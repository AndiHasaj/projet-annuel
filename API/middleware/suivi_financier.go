package middleware

import (
	"database/sql"
	"net/http"
	"API/repositories"
	"github.com/gin-gonic/gin"
)

func GetChiffreAffairesMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chiffreAffaire, err := repositories.GetChiffreAffaire(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du chiffre d'affaires"})
			return
		}
		if chiffreAffaire == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucun chiffre d'affaires trouvé"})
			return
		}
		c.Next()
	}
}

func GetBeneficeMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		benefice, err := repositories.GetBenefice(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du bénéfice"})
			return
		}
		if benefice == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucun bénéfice trouvé"})
			return
		}
		c.Next()
	}
}

func GetTauxBeneficeMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tauxBenefice, err := repositories.GetTauxBenefice(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du taux de bénéfice"})
			return
		}
		if tauxBenefice == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucun taux de bénéfice trouvé"})
			return
		}
		c.Next()
	}
}

func GetChargeMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		charges, err := repositories.GetCharges(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des charges"})
			return
		}
		if charges == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucune charge trouvée"})
			return
		}
		c.Next()
	}
}