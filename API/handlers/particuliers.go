package handlers

import (
	"API/model"
	"API/repositories"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)

func GetParticuliersHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		particuliers, err := repositories.GetAllParticuliers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des particuliers"})
			return
		}
		c.JSON(http.StatusOK, particuliers)
	}
}

func GetParticulierByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du particulier invalide"})
			return
		}

		particulier, err := repositories.GetParticulierByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du particulier"})
			return
		}
		if particulier == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Particulier non trouvé"})
			return
		}

		c.JSON(http.StatusOK, particulier)
	}
}

func GetParticulierByEmailHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Query("email")
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email du particulier manquant"})
			return
		}

		particulier, err := repositories.GetParticulierByEmail(db, email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du particulier"})
			return
		}
		if particulier == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Particulier non trouvé"})
			return
		}

		c.JSON(http.StatusOK, particulier)
	}
}

func CreateParticulierHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p model.Particulier

		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
			return
		}

		_, err := repositories.CreateParticulier(db, &p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création"})
			return
		}

		c.JSON(http.StatusOK, p)
	}
}

func UpdateParticulierHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}

		var p model.Particulier
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
			return
		}

		p.ID = id

		err = repositories.UpdateParticulier(db, &p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur update"})
			return
		}

		c.JSON(http.StatusOK, p)
	}
}

func DeleteParticulierHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}

		err = repositories.DeleteParticulier(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur delete"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Particulier supprimé avec succès"})
	}

}

func UpdateParticulierScoreHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du particulier invalide"})
			return
		}

		var body struct {
			Score int `json:"score"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
			return
		}

		err = repositories.UpdateScoreParticulier(db, id, body.Score)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour du score"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Score mis à jour avec succès"})
	}
}