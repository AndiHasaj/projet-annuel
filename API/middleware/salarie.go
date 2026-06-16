package middleware

import (
	"API/repositories"
	"database/sql"
	"net/http"
	"strconv"
	"math/rand"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetSalarieMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		salaries, err := repositories.GetAllSalaries(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des salariés"})
			return
		}
		if len(salaries) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucun salarié trouvé"})
			return
		}
		c.JSON(http.StatusOK, salaries)
	}
}

func GetSalarieByIDMiddleware(db *sql.DB,) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du salarié manquant"})
			return
		}

		salarie, err := repositories.GetSalarieByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du salarié"})
			return
		}
		if salarie == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Salarié non trouvé"})
			return
		}
		c.JSON(http.StatusOK, salarie)
	}
}

func GetSalarieByEmailMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Query("email")
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email du salarié manquant"})
			return
		}
		salarie, err := repositories.GetSalarieByEmail(db, email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du salarié"})
			return
		}
		if salarie == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Salarié non trouvé"})
			return
		}
		c.JSON(http.StatusOK, salarie)
	}
}

func CreateSalarieMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var body map[string]string

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
			return
		}
		nom := body["nom"]
		prenom := body["prenom"]
		email := prenom + "." + nom + string(rand.Intn(100)) + "@upcycle.fr"
		mdp := body["mot_de_passe"]
		intitulePoste := body["intitule_poste"]
		typeContrat := body["type_contrat"]
		salaire := body["salaire"]
		dateEmbauche := body["date_embauche"]

		if nom == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nom du salarié manquant"})
			return
		}
		if prenom == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Prénom du salarié manquant"})
			return
		}
		salarie, err := repositories.GetSalarieByEmail(db, email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la vérification de l'email du salarié"})
			return
		}
		if salarie != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Un salarié avec cet email existe déjà"})
			return
		}
		body["email"] = email
		if mdp == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Mot de passe du particulier manquant"})
			return
		}

		if len(mdp) < 8 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Mot de passe du particulier invalide"}	)
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Mot de passe du particulier invalide"})
			return
		}
		mdpHash, err := bcrypt.GenerateFromPassword([]byte(mdp), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du hachage du mot de passe du particulier"})
			return
		}
		body["mot_de_passe"] = string(mdpHash)

		if intitulePoste == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Intitulé du poste du salarié manquant"})
			return
		}

		if typeContrat == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Type de contrat du salarié manquant"})
			return
		}

		if salaire == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Salaire du salarié manquant"})
			return
		}

		
		if dateEmbauche == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Date d'embauche du salarié manquante"})
			return
		}
		c.Set("body", body)
		c.Next()
	}
}

func UpdateSalarieMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du salarié manquant"})
			return
		}
		salarie, err := repositories.GetSalarieByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du salarié"})
			return
		}
		if salarie == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Salarié non trouvé"})
			return
		}
		
		c.Next()
	}
}

func DeleteSalarieMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du salarié manquant"})
			return
		}
		salarie, err := repositories.GetSalarieByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du salarié"})
			return
		}
		if salarie == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Salarié non trouvé"})
			return
		}
		c.Next()
	}
}
