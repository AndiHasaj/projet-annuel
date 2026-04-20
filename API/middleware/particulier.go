package middleware

import (
	"API/repositories"
	"database/sql"
	"net/mail"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetParticulierMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		particuliers, err := repositories.GetAllParticuliers(db)
		if err != nil {
			c.JSON(500, gin.H{"error": "Erreur lors de la récupération des particuliers"})
			c.Abort()
			return
		}

		if len(particuliers) == 0 {
			c.JSON(404, gin.H{"error": "Aucun particulier trouvé"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func GetParticulierByIDMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "ID du particulier manquant"})
			c.Abort()
			return
		}

		particulier, err := repositories.GetParticulierByID(db, id)
		if err != nil {
			c.JSON(500, gin.H{"error": "Erreur lors de la récupération du particulier"})
			c.Abort()
			return
		}

		if particulier == nil {
			c.JSON(404, gin.H{"error": "Particulier non trouvé"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func GetParticulierByEmailMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		email := c.Query("email")
		if email == "" {
			c.JSON(400, gin.H{"error": "Email du particulier manquant"})
			c.Abort()
			return
		}

		particulier, err := repositories.GetParticulierByEmail(db, email)
		if err != nil {
			c.JSON(500, gin.H{"error": "Erreur lors de la récupération du particulier"})
			c.Abort()
			return
		}

		if particulier == nil {
			c.JSON(404, gin.H{"error": "Particulier non trouvé"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func CreateParticulierMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var body map[string]string

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "Données invalides"})
			c.Abort()
			return
		}

		nom := body["nom"]
		prenom := body["prenom"]
		email := body["email"]
		mdp := body["mot_de_passe"]

		if nom == "" {
			c.JSON(400, gin.H{"error": "Nom du particulier manquant"})
			c.Abort()
			return
		}

		if prenom == "" {
			c.JSON(400, gin.H{"error": "Prénom du particulier manquant"})
			c.Abort()
			return
		}

		for _, c2 := range nom {
			if (c2 >= '0' && c2 <= '9') || (c2 >= 33 && c2 <= 47) || (c2 >= 58 && c2 <= 64) || (c2 >= 91 && c2 <= 96) || (c2 >= 123 && c2 <= 126) {
				c.JSON(400, gin.H{"error": "Nom du particulier invalide"})
				c.Abort()
				return
			}
		}

		for _, c2 := range prenom {
			if (c2 >= '0' && c2 <= '9') || (c2 >= 33 && c2 <= 47) || (c2 >= 58 && c2 <= 64) || (c2 >= 91 && c2 <= 96) || (c2 >= 123 && c2 <= 126) {
				c.JSON(400, gin.H{"error": "Prénom du particulier invalide"})
				c.Abort()
				return
			}
		}

		if email == "" {
			c.JSON(400, gin.H{"error": "Email du particulier manquant"})
			c.Abort()
			return
		}

		_, err := mail.ParseAddress(email)
		if err != nil {
			c.JSON(400, gin.H{"error": "Email invalide"})
			c.Abort()
			return
		}

		existing, err := repositories.GetParticulierByEmail(db, email)
		if err != nil {
			c.JSON(500, gin.H{"error": "Erreur vérification email"})
			c.Abort()
			return
		}

		if existing != nil {
			c.JSON(409, gin.H{"error": "Email déjà utilisé"})
			c.Abort()
			return
		}

		if mdp == "" {
			c.JSON(400, gin.H{"error": "Mot de passe manquant"})
			c.Abort()
			return
		}

		if len(mdp) < 8 {
			c.JSON(400, gin.H{"error": "Mot de passe invalide"})
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
			c.JSON(400, gin.H{"error": "Mot de passe invalide"})
			c.Abort()
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(mdp), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "Erreur hash password"})
			c.Abort()
			return
		}

		body["mot_de_passe"] = string(hash)

		c.Set("body", body)
		c.Next()
	}
}

func UpdateParticulierMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "ID invalide"})
			c.Abort()
			return
		}

		particulier, err := repositories.GetParticulierByID(db, id)
		if err != nil {
			c.JSON(500, gin.H{"error": "Erreur serveur"})
			c.Abort()
			return
		}

		if particulier == nil {
			c.JSON(404, gin.H{"error": "Particulier non trouvé"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func DeleteParticulierMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "ID invalide"})
			c.Abort()
			return
		}

		particulier, err := repositories.GetParticulierByID(db, id)
		if err != nil {
			c.JSON(500, gin.H{"error": "Erreur serveur"})
			c.Abort()
			return
		}

		if particulier == nil {
			c.JSON(404, gin.H{"error": "Particulier non trouvé"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func UpdateScoreParticulierMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "ID invalide"})
			c.Abort()
			return
		}

		particulier, err := repositories.GetParticulierByID(db, id)
		if err != nil {
			c.JSON(500, gin.H{"error": "Erreur serveur"})
			c.Abort()
			return
		}

		if particulier == nil {
			c.JSON(404, gin.H{"error": "Particulier non trouvé"})
			c.Abort()
			return
		}

		c.Next()
	}
}
