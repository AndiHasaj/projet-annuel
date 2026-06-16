package handlers

import (
	"API/model"
	"API/repositories"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetSalariesHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		salaries, err := repositories.GetAllSalaries(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des salariés"})
			return
		}
		c.JSON(http.StatusOK, salaries)
	}
}

func GetSalarieByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}
		salarie, err := repositories.GetSalarieByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du salarié"})
			return
		}
		if salarie == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Salarié non trouvé"})
			return
		}
		c.JSON(http.StatusOK, salarie)
	}
}

func GetSalarieByEmailHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Query("email")
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email du salarié manquant"})
			return
		}

		salarie, err := repositories.GetSalarieByEmail(db, email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du salarié"})
			return
		}
		if salarie == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Salarié non trouvé"})
			return
		}
		c.JSON(http.StatusOK, salarie)
	}
}

func CreateSalarieHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var s model.Salarie
		body, exists := c.Get("body")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Body manquant"})
			return
		}
		data := body.(map[string]string)
		salaire, err := strconv.Atoi(data["salaire"])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Salaire invalide"})
			return
		}
		date_embauche, err := time.Parse("2006-01-02", data["date_embauche"])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Date d'embauche invalide"})
			return
		}
		s = model.Salarie{
			Nom:      data["nom"],
			Prenom:   data["prenom"],
			Email:    data["email"],
			MotDePasse: data["mot_de_passe"],
			IntitulePoste: data["intitule_poste"],
			TypeContrat: data["type_contrat"],
			Salaire: salaire,
			DateEmbauche: date_embauche,
		}
		_, err = repositories.CreateSalarie(db, &s)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création du salarié"})
			return
		}
		c.JSON(http.StatusOK, s)
	}
}

func UpdateSalarieHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}

		var s model.Salarie
		if err := json.NewDecoder(c.Request.Body).Decode(&s); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
			return
		}

		s.ID = id

		err = repositories.UpdateSalarie(db, &s);
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour du salarié"})
			return
		}
		c.JSON(http.StatusOK, s)
	}
}

func DeleteSalarieHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}
		salarie, err := repositories.GetSalarieByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du salarié"})
			return
		}
		if salarie == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Salarié non trouvé"})
			return
		}
		err = repositories.DeleteSalarie(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression du salarié"})
			return
		}
		c.Status(http.StatusNoContent)
	}
}