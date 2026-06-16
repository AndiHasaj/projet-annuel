package handlers

import (
	"API/model"
	"API/repositories"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetObjetsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		objets, err := repositories.GetAllObjets(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des objets"})
			return
		}
		c.JSON(http.StatusOK, objets)
	}
}

func GetObjetByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}
		objet, err := repositories.GetObjetByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération de l'objet"})
			return
		}
		if objet == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Objet non trouvé"})
			return
		}
		c.JSON(http.StatusOK, objet)
	}
}

func GetObjetsByIDConteneurHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idConteneur, err := strconv.Atoi(c.Query("id_conteneur"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du conteneur invalide"})
			return
		}
		objets, err := repositories.GetObjetsByIDConteneur(db, idConteneur)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des objets du conteneur"})
			return
		}
		c.JSON(http.StatusOK, objets)
	}
}

func CreateObjetHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var o model.Objet

		body, exists := c.Get("body")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Body manquant"})
			return
		}

		data := body.(map[string]string)

		o = model.Objet{
			Nom:          data["nom"],
			TypeObjet:    data["type_objet"],
			Description:  data["description"],
			Localisation: data["localisation"],
		}
		
		err := repositories.CreateObjet(db, &o)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de l'objet"})
			return
		}
		c.JSON(http.StatusCreated, o)
	}
}

func UpdateObjetHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}
		var objet model.Objet
		if err := json.NewDecoder(c.Request.Body).Decode(&objet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
			return
		}
		objet.ID = id
		if err := repositories.UpdateObjet(db, &objet); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour de l'objet"})
			return
		}
		c.JSON(http.StatusOK, objet)
	}
}

func DeleteObjetHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}
		objet, err := repositories.GetObjetByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération de l'objet"})
			return
		}
		if objet == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Objet non trouvé"})
			return
		}
		if err := repositories.DeleteObjet(db, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression de l'objet"})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{})
	}
}