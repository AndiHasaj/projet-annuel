package handlers

import (
	"API/model"
	"API/repositories"
	"database/sql"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetAdminsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		admins, err := repositories.GetAllAdmins(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des administrateurs"})
			return
		}
		c.JSON(http.StatusOK, admins)
	}
}

func GetAdminByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}
		admin, err := repositories.GetAdminByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération de l'administrateur"})
			return
		}
		if admin == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Administrateur non trouvé"})
			return
		}
		c.JSON(http.StatusOK, admin)
	}
}

func CreateAdminHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var a model.Admin

		body, exists := c.Get("body")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données manquantes"})
			return
		}

		data := body.(map[string]string)

		a = model.Admin{
			MotDePasse:  data["mot_de_passe"],
		}

		_, err := repositories.CreateAdmin(db, &a)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de l'administrateur"})
			return
		}
		c.JSON(http.StatusOK, a)
		
	}
}

func DeleteAdminHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}
		admin, err := repositories.GetAdminByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération de l'administrateur"})
			return
		}
		if admin == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Administrateur non trouvé"})
			return
		}
		if err := repositories.DeleteAdmin(db, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression de l'administrateur"})
			return
		}
		c.JSON(http.StatusOK, admin)
	}
}