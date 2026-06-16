package handlers

import (
	"API/model"
	"API/repositories"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAnnoncesHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		annonces, err := repositories.GetAllAnnonces(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des annonces"})
			return
		}
		c.JSON(http.StatusOK, annonces)
	}
}

func GetAnnonceByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
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

func GetAnnoncesByIDAnnonceurHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id_annonceur"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de l'annonceur invalide"})
			return
		}
		annonces, err := repositories.GetAnnoncesByIDAnnonceur(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des annonces de l'annonceur"})
			return
		}
		c.JSON(http.StatusOK, annonces)
	}
}

func CreateAnnonceHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var a model.Annonce
		body, exists := c.Get("body")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données d'annonce manquantes"})
			return
		}
		data := body.(map[string]string)
		id_annonceur, err := strconv.Atoi(data["id_annonceur"])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de l'annonceur manquant ou invalide"})
			return
		}

		prix, err := strconv.ParseFloat(data["prix"], 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Prix de l'annonce manquant ou invalide"})
			return
		}
		
		a = model.Annonce{
			IDAnnonceur: id_annonceur,
			Titre:        data["titre"],
			Description:   data["description"],
			Prix:         prix,
			DatePublication: time.Now(),
		}
		
		c.JSON(http.StatusCreated, a)
	}
}

func UpdateAnnonceHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var annonce model.Annonce
		if err := json.NewDecoder(c.Request.Body).Decode(&annonce); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données d'annonce invalides"})
			return
		}
		if err := repositories.UpdateAnnonce(db, &annonce); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour de l'annonce"})
			return
		}
		c.JSON(http.StatusOK, annonce)
	}
}

func DeleteAnnonceHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}
		err = repositories.DeleteAnnonce(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression de l'annonce"})
			return
		}
		c.Status(http.StatusNoContent)
	}
}