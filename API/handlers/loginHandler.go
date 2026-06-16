package handlers

import (
	"database/sql"
	"API/repositories"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginParticulierHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var body struct {
			Email      string `json:"email"`
			MotDePasse   string `json:"mot_de_passe"`
		}

		c.ShouldBindJSON(&body)

		particulier, err := repositories.GetParticulierByEmail(db, body.Email)
		if err != nil || particulier == nil {
			c.JSON(404, gin.H{"error": "Utilisateur introuvable"})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(particulier.MotDePasse), []byte(body.MotDePasse))

		if err != nil {
			c.JSON(401, gin.H{"error": "Mot de passe incorrect ! " + body.MotDePasse + " et " + particulier.MotDePasse})
			return
		}

		c.JSON(200, gin.H{
			"id": particulier.ID,
			"nom": particulier.Nom,
			"email": particulier.Email,
		})
	}
}

func LoginSalarieHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var body struct {
			Email      string `json:"email"`
			MotDePasse   string `json:"mot_de_passe"`
		}

		c.ShouldBindJSON(&body)

		salarie, err := repositories.GetSalarieByEmail(db, body.Email)
		if err != nil || salarie == nil {
			c.JSON(404, gin.H{"error": "Utilisateur introuvable"})
			return
		}

		mdpHash, err := bcrypt.GenerateFromPassword([]byte(body.MotDePasse), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "Erreur lors du hachage du mot de passe"})
			return
		}

		if salarie.MotDePasse != string(mdpHash) {
			c.JSON(401, gin.H{"error": "Mot de passe incorrect ! L'email étant la suivante : " + body.Email})
			return
		}

		c.JSON(200, gin.H{
			"id": salarie.ID,
			"nom": salarie.Nom,
			"email": salarie.Email,
		})
	}
}