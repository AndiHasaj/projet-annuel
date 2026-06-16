package handlers

import (
	"API/model"
	"API/repositories"
	"database/sql"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"time"
)

func GetServicesHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		services, err := repositories.GetAllServices(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des services"})
			return
		}
		c.JSON(http.StatusOK, services)
	}
}

func GetServiceByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}
		service, err := repositories.GetServiceByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du service"})
			return
		}
		if service == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Service non trouvé"})
			return
		}
		c.JSON(http.StatusOK, service)
	}
}

func GetServicesByTypeHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		type_service := c.Query("type")
		if type_service == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Type de service manquant"})
			return
		}
		services, err := repositories.GetServicesByType(db, type_service)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des services"})
			return
		}
		if len(services) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucun service trouvé pour ce type"})
			return
		}
		c.JSON(http.StatusOK, services)
	}
}

func CreateServiceHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var s model.Service
		
		body, exists := c.Get("body")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Body manquant"})
			return
		}

		data := body.(map[string]string)

		prix, err := strconv.ParseFloat(data["prix"], 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Prix invalide"})
			return
		}

		date_service, err := time.Parse("2006-01-02", data["date_service"])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Date de service invalide, format attendu : YYYY-MM-DD"})
			return
		}

		s = model.Service{
			TypeService:        data["type_service"],
			Titre:              data["titre"],
			Description:        data["description"],
			Prix:               prix,
			DateService:        date_service,
		}

		_, err = repositories.CreateService(db, &s)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création du service"})
			return
		}

		c.JSON(http.StatusOK, s)
	}
}

func UpdateServiceHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}
		var service model.Service
		if err := c.ShouldBindJSON(&service); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
			return
		}
		service.ID = id
		if err := repositories.UpdateService(db, &service); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour du service"})
			return
		}
		c.JSON(http.StatusOK, service)
	}
}

func DeleteServiceHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}
		service, err := repositories.GetServiceByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du service"})
			return
		}
		if service == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Service non trouvé"})
			return
		}
		if err := repositories.DeleteService(db, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression du service"})
			return
		}
		c.JSON(http.StatusOK, service)
	}
}