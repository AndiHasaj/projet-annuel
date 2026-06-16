package middleware

import (
	"database/sql"
	"net/http"
	"API/repositories"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetServiceMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		services, err := repositories.GetAllServices(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des services"})
			return
		}
		if len(services) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucun service trouvé"})
			return
		}
		c.JSON(http.StatusOK, services)
	}
}

func GetServiceByIDMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du service manquant"})
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

func GetServicesByTypeMiddleware(db *sql.DB) gin.HandlerFunc {
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

func CreateServiceMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body map[string]string
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
			return
		}

		type_service := body["type"]
		titre := body["titre"]
		description := body["description"]
		prix, err := strconv.ParseFloat(body["prix"], 64)
		date_service := body["date_service"]

		if type_service == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Type de service manquant"})
			return
		}
		if titre == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Titre du service manquant"})
			return
		}
		if description == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Description du service manquante"})
			return
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Prix du service invalide"})
			return
		}
		if date_service == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Date du service manquante"})
			return
		}

		body["prix"] = strconv.FormatFloat(prix, 'f', 2, 64)
		c.Set("body", body)
		c.Next()
	}

}

func UpdateServiceMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du service manquant"})
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
		c.Next()
	}
}

func DeleteServiceMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du service manquant"})
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
		c.Next()
	}
}