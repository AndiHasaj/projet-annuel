package router

import (
	"API/handlers"
	"github.com/gin-gonic/gin"
	"database/sql"
)

var db *sql.DB

func SetupRouter(database *sql.DB) *gin.Engine {
	db = database
	router := gin.Default()
	
	// Routes pour les particuliers
	particulierGroup := router.Group("/particuliers")
	{
		particulierGroup.GET("/", handlers.GetParticuliersHandler(db))
		particulierGroup.GET("/:id", handlers.GetParticulierByIDHandler(db))
		particulierGroup.GET("/email/:email", handlers.GetParticulierByEmailHandler(db))
		particulierGroup.POST("/", handlers.CreateParticulierHandler(db))
		particulierGroup.PUT("/:id", handlers.UpdateParticulierHandler(db))
		particulierGroup.DELETE("/:id", handlers.DeleteParticulierHandler(db))
		particulierGroup.PUT("/score/:id",handlers.UpdateParticulierScoreHandler(db))
	}
	// Routes pour les salariés
	salarieGroup := router.Group("/salaries")
	{
		salarieGroup.GET("/", gin.WrapH(handlers.GetSalariesHandler(db)))
		salarieGroup.GET("/:id", gin.WrapH(handlers.GetSalarieByIDHandler(db)))
		salarieGroup.GET("/email/:email", gin.WrapH(handlers.GetSalarieByEmailHandler(db)))
		salarieGroup.POST("/", gin.WrapH(handlers.CreateSalarieHandler(db)))
		salarieGroup.PUT("/:id", gin.WrapH(handlers.UpdateSalarieHandler(db)))
		salarieGroup.DELETE("/:id", gin.WrapH(handlers.DeleteSalarieHandler(db)))
	}
	// Routes pour les professionnels
	professionnelGroup := router.Group("/professionnels")
	{
		professionnelGroup.GET("/", gin.WrapH(handlers.GetProfessionnelsHandler(db)))
		professionnelGroup.GET("/:id", gin.WrapH(handlers.GetProfessionnelByIDHandler(db)))
		professionnelGroup.GET("/email/:email", gin.WrapH(handlers.GetProfessionnelByEmailHandler(db)))
		professionnelGroup.POST("/", gin.WrapH(handlers.CreateProfessionnelHandler(db)))
		professionnelGroup.PUT("/:id", gin.WrapH(handlers.UpdateProfessionnelHandler(db)))
		professionnelGroup.DELETE("/:id", gin.WrapH(handlers.DeleteProfessionnelHandler(db)))
	}
	// Routes pour les annonces
	annonceGroup := router.Group("/annonces")
	{
		annonceGroup.GET("/", gin.WrapH(handlers.GetAnnoncesHandler(db)))
		annonceGroup.GET("/:id", gin.WrapH(handlers.GetAnnonceByIDHandler(db)))
		annonceGroup.GET("/annonceur/:idAnnonceur", gin.WrapH(handlers.GetAnnoncesByIDAnnonceurHandler(db)))
		annonceGroup.POST("/", gin.WrapH(handlers.CreateAnnonceHandler(db)))
		annonceGroup.PUT("/:id", gin.WrapH(handlers.UpdateAnnonceHandler(db)))
		annonceGroup.DELETE("/:id", gin.WrapH(handlers.DeleteAnnonceHandler(db)))
	}
	// Routes pour les objets
	objetGroup := router.Group("/objets")
	{
		objetGroup.GET("/", gin.WrapH(handlers.GetObjetsHandler(db)))
		objetGroup.GET("/:id", gin.WrapH(handlers.GetObjetByIDHandler(db)))
		objetGroup.GET("/conteneur/:id_conteneur", gin.WrapH(handlers.GetObjetsByIDConteneurHandler(db)))
		objetGroup.POST("/", gin.WrapH(handlers.CreateObjetHandler(db)))
		objetGroup.PUT("/:id", gin.WrapH(handlers.UpdateObjetHandler(db)))
		objetGroup.DELETE("/:id", gin.WrapH(handlers.DeleteObjetHandler(db)))
	}
	// Routes pour les conteneurs
	conteneurGroup := router.Group("/conteneurs")
	{
		conteneurGroup.GET("/", gin.WrapH(handlers.GetConteneursHandler(db)))
		conteneurGroup.GET("/:id", gin.WrapH(handlers.GetConteneurByIDHandler(db)))
		conteneurGroup.GET("/matricule/:matricule", gin.WrapH(handlers.GetConteneurByMatriculeHandler(db)))
		conteneurGroup.GET("/localisation/:localisation", gin.WrapH(handlers.GetConteneursByLocalisationHandler(db)))
		conteneurGroup.GET("/code/:code", gin.WrapH(handlers.GetConteneurByCodeHandler(db)))
		conteneurGroup.POST("/", gin.WrapH(handlers.CreateConteneurHandler(db)))
		conteneurGroup.PUT("/:id", gin.WrapH(handlers.UpdateConteneurHandler(db)))
		conteneurGroup.DELETE("/:id", gin.WrapH(handlers.DeleteConteneurHandler(db)))
	}
	// Routes pour les services
	serviceGroup := router.Group("/services")
	{
		serviceGroup.GET("/", gin.WrapH(handlers.GetServicesHandler(db)))
		serviceGroup.GET("/:id", gin.WrapH(handlers.GetServiceByIDHandler(db)))
		serviceGroup.GET("/type/:type", gin.WrapH(handlers.GetServicesByTypeHandler(db)))
		serviceGroup.POST("/", gin.WrapH(handlers.CreateServiceHandler(db)))
		serviceGroup.PUT("/:id", gin.WrapH(handlers.UpdateServiceHandler(db)))
		serviceGroup.DELETE("/:id", gin.WrapH(handlers.DeleteServiceHandler(db)))
	}
	// Routes pour les paiements
	paiementGroup := router.Group("/paiements")
	{
		paiementGroup.GET("/", gin.WrapH(handlers.GetPaiementsHandler(db)))
		paiementGroup.GET("/:id", gin.WrapH(handlers.GetPaiementByIDHandler(db)))
		paiementGroup.GET("/payeur/:idPayeur", gin.WrapH(handlers.GetPaiementsByIDPayeurHandler(db)))
		paiementGroup.GET("/receveur/:idReceveur", gin.WrapH(handlers.GetPaiementsByIDReceveurHandler(db)))
		paiementGroup.POST("/", gin.WrapH(handlers.CreatePaiementHandler(db)))
	}

	// Routes pour le suivi financier
	suiviFinancierGroup := router.Group("/suivi-financier")
	{
		suiviFinancierGroup.GET("/chiffre-d-affaires", gin.WrapH(handlers.GetChiffreAffairesHandler(db)))
		suiviFinancierGroup.GET("/charges", gin.WrapH(handlers.GetChargeHandler(db)))
		suiviFinancierGroup.GET("/benefice", gin.WrapH(handlers.GetBeneficeHandler(db)))
		suiviFinancierGroup.GET("/taux-benefice", gin.WrapH(handlers.GetTauxBeneficeHandler(db)))
	}

	// Routes pour les administrateurs
	adminGroup := router.Group("/admin")
	{
		adminGroup.GET("/", gin.WrapH(handlers.GetAdminsHandler(db)))
		adminGroup.GET("/?id=:id", gin.WrapH(handlers.GetAdminByIDHandler(db)))
	}

	login := router.Group("/login")
	{
		login.POST("/", handlers.LoginHandler(db))
	}
	return router
}