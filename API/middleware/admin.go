package middleware

import (
	"database/sql"
	"net/http"
	"API/repositories"
	"strconv"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
)

func GetAdminMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		admins, err := repositories.GetAllAdmins(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des administrateurs"})
			return
		}
		if len(admins) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucun administrateur trouvé"})
			return
		}
		c.Next()
	}
}

func GetAdminByIDMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de l'administrateur manquant"})
			return
		}
		admin, err := repositories.GetAdminByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération de l'administrateur"})
			return
		}
		if admin == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Administrateur non trouvé"})
			return
		}
		c.Next()
	}
}

func CreateAdminMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body map[string]string
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
			return
		}

		mdp := body["mot_de_passe"]
		
		if mdp == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Mot de passe manquant"})
			c.Abort()
			return
		}

		if len(mdp) < 8 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Mot de passe invalide"})
			c.Abort()
			return
		}

		hasUpper, hasLower, hasDigit, hasSpecial := false, false, false, false

		for _, c2 := range mdp {
			switch {
			case c2 >= 'A' && c2 <= 'Z':
				hasUpper = true
			case c2 >= 'a' && c2 <= 'z':
				hasLower = true
			case c2 >= '0' && c2 <= '9':
				hasDigit = true
			case (c2 >= 33 && c2 <= 47) || (c2 >= 58 && c2 <= 64) || (c2 >= 91 && c2 <= 96) || (c2 >= 123 && c2 <= 126):
				hasSpecial = true
			}
		}

		if !hasUpper || !hasLower || !hasDigit || !hasSpecial {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Mot de passe invalide"})
			c.Abort()
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(mdp), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur hash password"})
			c.Abort()
			return
		}

		body["mot_de_passe"] = string(hash)

		c.Set("body", body)
		c.Next()
	}
}

func DeleteAdminMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de l'administrateur manquant"})
			return
		}
		admin, err := repositories.GetAdminByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération de l'administrateur"})
			return
		}
		if admin == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Administrateur non trouvé"})
			return
		}
		c.Next()
	}
}