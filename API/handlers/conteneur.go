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

func GetConteneursHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		conteneurs, err := repositories.GetAllConteneurs(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des conteneurs"})
			return
		}
		c.JSON(http.StatusOK, conteneurs)
	}
}

func GetConteneurByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
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

func GetConteneurByMatriculeHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		matricule := c.Query("matricule")
		conteneur, err := repositories.GetConteneurByMatricule(db, matricule)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du conteneur"})
			return
		}
		c.JSON(http.StatusOK, conteneur)
	}
}

func GetConteneursByLocalisationHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		localisation := c.Query("localisation")
		conteneurs, err := repositories.GetConteneursByLocalisation(db, localisation)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des conteneurs"})
			return
		}
		c.JSON(http.StatusOK, conteneurs)
	}
}

func GetConteneurByCodeHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")
		conteneur, err := repositories.GetConteneurByCode(db, code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du conteneur"})
			return
		}
		c.JSON(http.StatusOK, conteneur)
	}
}

func CreateConteneurHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cont model.Conteneur

		body, exists := c.Get("body")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données du conteneur manquantes"})
			return
		}
		data := body.(map[string]string)

		occupation, err := strconv.ParseInt(data["occupation"], 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Occupation du conteneur invalide"})
			return
		}

		cont = model.Conteneur{
			Matricule:   data["matricule"],
			Localisation: data["localisation"],
			Statut:       data["statut"],
			Occupation:    occupation,
			Code:         data["code"],
		}

		err = repositories.CreateConteneur(db, &cont)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création du conteneur"})
			return
		}
		c.JSON(http.StatusCreated, cont)
	}
}

func UpdateConteneurHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}
		var conteneur model.Conteneur
		if err := json.NewDecoder(c.Request.Body).Decode(&conteneur); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
			return
		}
		if conteneur.Matricule == ""{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Matricule du conteneur est requis"})
			return
		}
		conteneur.ID = id
		if err := repositories.UpdateConteneur(db, &conteneur); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour du conteneur"})
			return
		}
		c.JSON(http.StatusOK, conteneur)
	}
}

func DeleteConteneurHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}
		if err := repositories.DeleteConteneur(db, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression du conteneur"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Conteneur supprimé avec succès"})
	}
}
