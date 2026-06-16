package middleware

import (
	"database/sql"
	"net/http"
	"API/repositories"
	"strconv"
	"math/rand"
	"github.com/gin-gonic/gin"
)

func GetConteneurMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		conteneurs, err := repositories.GetAllConteneurs(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des conteneurs"})
			return
		}
		if len(conteneurs) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucun conteneur trouvé"})
			return
		}
		c.JSON(http.StatusOK, conteneurs)
	}
}

func GetConteneurByIDMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du conteneur manquant"})
			return
		}
		conteneur, err := repositories.GetConteneurByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du conteneur"})
			return
		}
		if conteneur == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Conteneur non trouvé"})
			return
		}
		c.JSON(http.StatusOK, conteneur)
	}
}

func GetConteneurByMatriculeMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		matricule := c.Query("matricule")
		if matricule == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Matricule du conteneur manquant"})
			return
		}
		conteneur, err := repositories.GetConteneurByMatricule(db, matricule)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du conteneur"})
			return
		}
		if conteneur == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Conteneur non trouvé"})
			return
		}
		c.JSON(http.StatusOK, conteneur)
	}
}

func GetConteneursByLocalisationIDMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		localisation := c.Query("localisation")
		if localisation == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Localisation du conteneur manquante"})
			return
		}
		conteneurs, err := repositories.GetConteneursByLocalisation(db, localisation)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des conteneurs"})
			return
		}
		if len(conteneurs) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucun conteneur trouvé pour cette localisation"})
			return
		}
		c.JSON(http.StatusOK, conteneurs)
	}
}

func GetConteneurByCodeMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")
		if code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Code du conteneur manquant"})
			return
		}
		conteneur, err := repositories.GetConteneurByCode(db, code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du conteneur"})
			return
		}
		if conteneur == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucun conteneur trouvé pour ce code"})
			return
		}
		c.JSON(http.StatusOK, conteneur)
	}
}


func CreateConteneurMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body map[string]string

		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
			c.Abort()
			return
		}

		matricule := body["matricule"]
		localisation := body["localisation"]
		statut := body["statut"]
		occupation := body["occupation"]

		if matricule == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Matricule du conteneur manquant"})
			return
		}
		if localisation == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Localisation du conteneur manquante"})
			return
		}
		if statut == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Statut du conteneur manquant"})
			return
		}
		if occupation == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Occupation du conteneur manquante"})
			return
		}
		
		code := strconv.Itoa(rand.Intn(100000000))
		body["code"] = code

		c.Set("body", body)
		c.Next()
	}
}

func UpdateConteneurMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du conteneur manquant"})
			return
		}
		conteneur, err := repositories.GetConteneurByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du conteneur"})
			return
		}
		if conteneur == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Conteneur non trouvé"})
			return
		}
		c.Next()	
	}
}

func DeleteConteneurMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du conteneur manquant"})
			return
		}
		conteneur, err := repositories.GetConteneurByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du conteneur"})
			return
		}
		if conteneur == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Conteneur non trouvé"})
			return
		}
		c.Next()
	}
}