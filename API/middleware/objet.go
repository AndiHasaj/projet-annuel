package middleware

import (
	"database/sql"
	"net/http"
	"API/repositories"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetObjetMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		objets, err := repositories.GetAllObjets(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des objets"})
			return
		}
		if len(objets) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucun objet trouvé"})
			return
		}
		c.JSON(http.StatusOK, objets)
	}
}

func GetObjetByIDMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de l'objet manquant"})
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

func GetObjetsByConteneurIDMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		conteneurID, err := strconv.Atoi(c.Query("conteneur_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID du conteneur manquant"})
			return
		}
		objets, err := repositories.GetObjetsByIDConteneur(db, conteneurID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des objets"})
			return
		}
		if len(objets) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucun objet trouvé pour ce conteneur"})
			return
		}
		c.JSON(http.StatusOK, objets)
	}
}

func CreateObjetMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var body map[string]string
		
		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
			c.Abort()
			return
		}
		nom := body["nom"]
		type_objet := body["type_objet"]
		localisation := body["localisation"]
		if nom == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nom de l'objet manquant"})
			return
		}
		if type_objet == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Type de l'objet manquant"})
			return
		}
		if localisation == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Localisation de l'objet manquante"})
			return
		}
		c.Set("body", body)
		c.Next()
	}
}

func UpdateObjetMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de l'objet manquant"})
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
		c.Next()
	}
}

func DeleteObjetMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de l'objet manquant"})
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
		c.Next()
	}
}