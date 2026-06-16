package middleware

import (
	"database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
	"API/repositories"
	"strconv"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
)

func GetProfessionnelMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		professionnels, err := repositories.GetAllProfessionnels(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des professionnels"})
			return
		}
		if len(professionnels) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucun professionnel trouvé"})
			return
		}
		c.Next()
	}
}

func GetProfessionnelByIDMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du professionnel manquant"})
			return
		}
		professionnel, err := repositories.GetProfessionnelByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du professionnel"})
			return
		}
		if professionnel == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Professionnel non trouvé"})
			return
		}
		c.JSON(http.StatusOK, professionnel)
	}
}

func GetProfessionnelByEmailMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Query("email")
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email du professionnel manquant"})
			return
		}
		professionnel, err := repositories.GetProfessionnelByEmail(db, email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du professionnel"})
			return
		}
		if professionnel == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Professionnel non trouvé"})
			return
		}
		c.JSON(http.StatusOK, professionnel)
	}
}

func CreateProfessionnelMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var body map[string]string
		
		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(400, gin.H{"error": "Données iinvalides"})
			c.Abort()
			return
		}
		nomEntreprise := body["nom_entreprise"]
		email := body["email"]
		mdp := body["mot_de_passe"]
		numeroSiret := body["numero_siret"]
		telephone := body["telephone"]
		site_web := body["site_web"]
		adresse := body["adresse"]
		codePostal := body["code_postal"]
		ville := body["ville"]

		if nomEntreprise == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nom de l'entreprise du professionnel manquant"})
			return
		}
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email du professionnel manquant"})
			return
		}
		_, err = mail.ParseAddress(email)
		if err != nil {
			c.JSON(400, gin.H{"error": "Email invalide"})
			c.Abort()
			return
		}

		existing, err := repositories.GetProfessionnelByEmail(db, email)
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

		hasUpper := false
		hasLower := false
		hasDigit := false
		hasSpecial := false
		for _, c := range mdp {
			switch {
			case 'A' <= c && c <= 'Z':
				hasUpper = true
			case 'a' <= c && c <= 'z':
				hasLower = true
			case '0' <= c && c <= '9':
				hasDigit = true
			case (c >= 33 && c <= 47) || (c >= 58 && c <= 64) || (c >= 91 && c <= 96) || (c >= 123 && c <= 126):
				hasSpecial = true
			}
		}
		if !hasUpper || !hasLower || !hasDigit || !hasSpecial {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Mot de passe du professionnel invalide"})
			return
		}
		mdpHash, err := bcrypt.GenerateFromPassword([]byte(mdp), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du hachage du mot de passe du professionnel"})
			return
		}
		body["mot_de_passe"] = string(mdpHash)

		if numeroSiret == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Numéro de SIRET du professionnel manquant"})
			return
		}
		if len(numeroSiret) != 14 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Numéro de SIRET du professionnel invalide"})
			return
		}
		for _, C := range numeroSiret {
			if C < '0' || C > '9' {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Numéro de SIRET du professionnel invalide"})
				return
			}
		}
		// Possibilité de faire la vérif par api.insee.fr mais pas de clé d'api pour le moment

		if telephone == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Téléphone du professionnel manquant"})
			return
		}
		if len(telephone) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Téléphone du professionnel invalide"})
			return
		}
		for _, C := range telephone {
			if C < '0' || C > '9' {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Téléphone du professionnel invalide"})
				return
			}
		}

		if site_web == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Site web du professionnel manquant"})
			return
		}
		if adresse == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Adresse du professionnel manquante"})
			return
		}

		if codePostal == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Code postal du professionnel manquant"})
			return
		}
		if len(codePostal) != 5 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Code postal du professionnel invalide"})
			return
		}
		for _, C := range codePostal {
			if C < '0' || C > '9' {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Code postal du professionnel invalide"})
				return
			}
		}

		if ville == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ville du professionnel manquante"})
			return
		}
		c.Set("body", body)
		c.Next()
	}
}

func UpdateProfessionnelMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du professionnel manquant"})
			return
		}
		professionnel, err := repositories.GetProfessionnelByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du professionnel"})
			return
		}
		if professionnel == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Professionnel non trouvé"})
			return
		}
		c.Next()
	}
}

func DeleteProfessionnelMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du professionnel manquant"})
			return
		}
		professionnel, err := repositories.GetProfessionnelByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du professionnel"})
			return
		}
		if professionnel == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Professionnel non trouvé"})
			return
		}
		c.Next()
	}
}