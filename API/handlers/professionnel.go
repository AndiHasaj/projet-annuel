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

func GetProfessionnelsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		professionnels, err := repositories.GetAllProfessionnels(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des professionnels"})
			return
		}
		c.JSON(http.StatusOK, professionnels)
	}
}

func GetProfessionnelByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
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

func GetProfessionnelByEmailHandler(db *sql.DB) gin.HandlerFunc {
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

func CreateProfessionnelHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p model.Professionnel

		body, existing := c.Get("body")
		if !existing {
			c.JSON(400, gin.H{"error": "Body manquant"})
			return
		}
		data := body.(map[string]string)

		p = model.Professionnel{
			NomEntreprise:  data["nom_entreprise"],
			Email:          data["email"],
			MotDePasse:     data["mot_de_passe"],
			NumeroSiret:    data["numero_siret"],
			Telephone:      data["telephone"],
			SiteWeb:        data["site_web"],
			Adresse:        data["adresse"],
			CodePostal:     data["code_postal"],
			Ville:          data["ville"],
		}

		_, err := repositories.CreateProfessionnel(db, &p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création du professionnel"})
			return
		}

		c.JSON(http.StatusOK, p)
	}
}

func UpdateProfessionnelHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}

		var p model.Professionnel
		if err := json.NewDecoder(c.Request.Body).Decode(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
			return
		}

		p.ID = id

		err = repositories.UpdateProfessionnel(db, &p)
		if  err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour du professionnel"})
			return
		}

		c.JSON(http.StatusOK, p)
	}
}

func DeleteProfessionnelHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
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
		
		err = repositories.DeleteProfessionnel(db, id)
		if  err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression du professionnel"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Professionnel supprimé avec succès"})
	}
}