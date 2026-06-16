package middleware

import (
	"database/sql"
	"net/http"
	"API/repositories"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetAnnonceMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		annonces, err := repositories.GetAllAnnonces(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des annonces"})
			return
		}
		if len(annonces) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucune annonce trouvée"})
			return
		}
		c.JSON(http.StatusOK, annonces)
	}
}

func GetAnnonceByIDMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de l'annonce manquant"})
			return
		}
		annonce, err := repositories.GetAnnonceByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération de l'annonce"})
			return
		}
		if annonce == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Annonce non trouvée"})
			return
		}
		c.JSON(http.StatusOK, annonce)
	}
}

func GetAnnoncesByIDAnnonceurMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idAnnonceur, err := strconv.Atoi(c.Query("id_annonceur"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de l'annonceur manquant"})
			return
		}
		annonces, err := repositories.GetAnnoncesByIDAnnonceur(db, idAnnonceur)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des annonces"})
			return
		}
		if len(annonces) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucune annonce trouvée pour ce particulier"})
			return
		}
		c.JSON(http.StatusOK, annonces)
	}
}

func CreateAnnonceMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var body map[string]string

		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
			c.Abort()
			return
		}

		titre := body["titre"]
		description := body["description"]
		prix, err := strconv.ParseFloat(body["prix"], 64)
		particulierID, err := strconv.Atoi(body["particulier_id"])

		if titre == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Titre de l'annonce manquant"})
			return
		}
		if description == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Description de l'annonce manquante"})
			return
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Prix de l'annonce manquant ou invalide"})
			return
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du particulier manquant ou invalide"})
			return
		}
		body["prix"] = strconv.FormatFloat(prix, 'f', 2, 64)
		body["particulier_id"] = strconv.Itoa(particulierID)
		c.Set("body", body)
		c.Next()
	}
}

func UpdateAnnonceMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de l'annonce manquant"})
			return
		}
		annonce, err := repositories.GetAnnonceByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération de l'annonce"})
			return
		}
		if annonce == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Annonce non trouvée"})
			return
		}
		c.Next()
	}
}

func DeleteAnnonceMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de l'annonce manquant"})
			return
		}
		annonce, err := repositories.GetAnnonceByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération de l'annonce"})
			return
		}
		if annonce == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Annonce non trouvée"})
			return
		}
		c.Next()
	}
}