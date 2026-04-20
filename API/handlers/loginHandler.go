package handlers

import (
	"database/sql"
	"API/repositories"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var body struct {
			Email      string `json:"email"`
			Password   string `json:"password"`
		}

		c.ShouldBindJSON(&body)

		particulier, err := repositories.GetParticulierByEmail(db, body.Email)
		if err != nil || particulier == nil {
			c.JSON(404, gin.H{"error": "Utilisateur introuvable"})
			return
		}

		err = bcrypt.CompareHashAndPassword(
			[]byte(particulier.MotDePasse),
			[]byte(body.Password),
		)

		if err != nil {
			c.JSON(401, gin.H{"error": "Mot de passe incorrect"})
			return
		}

		c.JSON(200, gin.H{
			"id": particulier.ID,
			"nom": particulier.Nom,
			"email": particulier.Email,
		})
	}
}