package router

import (
	"API/handlers"
	"API/middleware"
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
		particulierGroup.GET("/", middleware.GetParticulierMiddleware(db), handlers.GetParticuliersHandler(db))
		particulierGroup.GET("/:id",middleware.GetParticulierByIDMiddleware(db), handlers.GetParticulierByIDHandler(db))
		particulierGroup.GET("/email/:email",middleware.GetParticulierByEmailMiddleware(db), handlers.GetParticulierByEmailHandler(db))
		particulierGroup.POST("/",middleware.CreateParticulierMiddleware(db), handlers.CreateParticulierHandler(db))
		particulierGroup.PUT("/:id", middleware.UpdateParticulierMiddleware(db), handlers.UpdateParticulierHandler(db))
		particulierGroup.DELETE("/:id", middleware.DeleteParticulierMiddleware(db), handlers.DeleteParticulierHandler(db))
		particulierGroup.PUT("/score/:id",middleware.UpdateScoreParticulierMiddleware(db),handlers.UpdateParticulierScoreHandler(db))
	}
	// Routes pour les salariés
	salarieGroup := router.Group("/salaries")
	{
		salarieGroup.GET("/", middleware.GetSalarieMiddleware(db), handlers.GetSalariesHandler(db))
		salarieGroup.GET("/:id", middleware.GetSalarieByIDMiddleware(db), handlers.GetSalarieByIDHandler(db))
		salarieGroup.GET("/email/:email", middleware.GetSalarieByEmailMiddleware(db), handlers.GetSalarieByEmailHandler(db))
		salarieGroup.POST("/", middleware.CreateSalarieMiddleware(db), handlers.CreateSalarieHandler(db))
		salarieGroup.PUT("/:id", middleware.UpdateSalarieMiddleware(db), handlers.UpdateSalarieHandler(db))
		salarieGroup.DELETE("/:id", middleware.DeleteSalarieMiddleware(db), handlers.DeleteSalarieHandler(db))
	}
	// Routes pour les professionnels
	professionnelGroup := router.Group("/professionnels")
	{
		professionnelGroup.GET("/", middleware.GetProfessionnelMiddleware(db), handlers.GetProfessionnelsHandler(db))
		professionnelGroup.GET("/:id", middleware.GetProfessionnelByIDMiddleware(db), handlers.GetProfessionnelByIDHandler(db))
		professionnelGroup.GET("/email/:email", middleware.GetProfessionnelByEmailMiddleware(db), handlers.GetProfessionnelByEmailHandler(db))
		professionnelGroup.POST("/", middleware.CreateProfessionnelMiddleware(db), handlers.CreateProfessionnelHandler(db))
		professionnelGroup.PUT("/:id", middleware.UpdateProfessionnelMiddleware(db), handlers.UpdateProfessionnelHandler(db))
		professionnelGroup.DELETE("/:id", middleware.DeleteProfessionnelMiddleware(db), handlers.DeleteProfessionnelHandler(db))
	}
	// Routes pour les annonces
	annonceGroup := router.Group("/annonces")
	{
		annonceGroup.GET("/", middleware.GetAnnonceMiddleware(db), handlers.GetAnnoncesHandler(db))
		annonceGroup.GET("/:id", middleware.GetAnnonceByIDMiddleware(db), handlers.GetAnnonceByIDHandler(db))
		annonceGroup.GET("/annonceur/:idAnnonceur", middleware.GetAnnoncesByIDAnnonceurMiddleware(db), handlers.GetAnnoncesByIDAnnonceurHandler(db))
		annonceGroup.POST("/", middleware.CreateAnnonceMiddleware(db), handlers.CreateAnnonceHandler(db))
		annonceGroup.PUT("/:id", middleware.UpdateAnnonceMiddleware(db), handlers.UpdateAnnonceHandler(db))
		annonceGroup.DELETE("/:id", middleware.DeleteAnnonceMiddleware(db), handlers.DeleteAnnonceHandler(db))
	}
	// Routes pour les objets
	objetGroup := router.Group("/objets")
	{
		objetGroup.GET("/", middleware.GetObjetMiddleware(db), handlers.GetObjetsHandler(db))
		objetGroup.GET("/:id", middleware.GetObjetByIDMiddleware(db), handlers.GetObjetByIDHandler(db))
		objetGroup.GET("/conteneur/:id_conteneur", middleware.GetObjetsByConteneurIDMiddleware(db), handlers.GetObjetsByIDConteneurHandler(db))
		objetGroup.POST("/", middleware.CreateObjetMiddleware(db), handlers.CreateObjetHandler(db))
		objetGroup.PUT("/:id", middleware.UpdateObjetMiddleware(db), handlers.UpdateObjetHandler(db))
		objetGroup.DELETE("/:id", middleware.DeleteObjetMiddleware(db), handlers.DeleteObjetHandler(db))
	}
	// Routes pour les conteneurs
	conteneurGroup := router.Group("/conteneurs")
	{
		conteneurGroup.GET("/", middleware.GetConteneurMiddleware(db), handlers.GetConteneursHandler(db))
		conteneurGroup.GET("/:id", middleware.GetConteneurByIDMiddleware(db), handlers.GetConteneurByIDHandler(db))
		conteneurGroup.GET("/matricule/:matricule", middleware.GetConteneurByMatriculeMiddleware(db), handlers.GetConteneurByMatriculeHandler(db))
		conteneurGroup.GET("/localisation/:localisation", middleware.GetConteneursByLocalisationIDMiddleware(db), handlers.GetConteneursByLocalisationHandler(db))
		conteneurGroup.GET("/code/:code", middleware.GetConteneurByCodeMiddleware(db), handlers.GetConteneurByCodeHandler(db))
		conteneurGroup.POST("/", middleware.CreateConteneurMiddleware(db), handlers.CreateConteneurHandler(db))
		conteneurGroup.PUT("/:id", middleware.UpdateConteneurMiddleware(db), handlers.UpdateConteneurHandler(db))
		conteneurGroup.DELETE("/:id", middleware.DeleteConteneurMiddleware(db), handlers.DeleteConteneurHandler(db))
	}
	// Routes pour les services
	serviceGroup := router.Group("/services")
	{
		serviceGroup.GET("/", middleware.GetServiceMiddleware(db), handlers.GetServicesHandler(db))
		serviceGroup.GET("/:id", middleware.GetServiceByIDMiddleware(db), handlers.GetServiceByIDHandler(db))
		serviceGroup.GET("/type/:type", middleware.GetServicesByTypeMiddleware(db), handlers.GetServicesByTypeHandler(db))
		serviceGroup.POST("/", middleware.CreateServiceMiddleware(db), handlers.CreateServiceHandler(db))
		serviceGroup.PUT("/:id", middleware.UpdateServiceMiddleware(db), handlers.UpdateServiceHandler(db))
		serviceGroup.DELETE("/:id", middleware.DeleteServiceMiddleware(db), handlers.DeleteServiceHandler(db))
	}
	// Routes pour les paiements
	paiementGroup := router.Group("/paiements")
	{
		paiementGroup.GET("/", middleware.GetPaiementMiddleware(db), handlers.GetPaiementsHandler(db))
		paiementGroup.GET("/:id", middleware.GetPaiementByIDMiddleware(db), handlers.GetPaiementByIDHandler(db))
		paiementGroup.GET("/payeur/:idPayeur", middleware.GetPaiementsByIDPayeurMiddleware(db), handlers.GetPaiementsByIDPayeurHandler(db))
		paiementGroup.GET("/receveur/:idReceveur", middleware.GetPaiementsByIDReceveurMiddleware(db), handlers.GetPaiementsByIDReceveurHandler(db))
		paiementGroup.POST("/", middleware.CreatePaiementMiddleware(db), handlers.CreatePaiementHandler(db))
	}

	// Routes pour le suivi financier
	suiviFinancierGroup := router.Group("/suivi-financier")
	{
		suiviFinancierGroup.GET("/chiffre-d-affaires", middleware.GetChiffreAffairesMiddleware(db), handlers.GetChiffreAffairesHandler(db))
		suiviFinancierGroup.GET("/charges", middleware.GetChargeMiddleware(db), handlers.GetChargeHandler(db))
		suiviFinancierGroup.GET("/benefice", middleware.GetBeneficeMiddleware(db), handlers.GetBeneficeHandler(db))
		suiviFinancierGroup.GET("/taux-benefice", middleware.GetTauxBeneficeMiddleware(db), handlers.GetTauxBeneficeHandler(db))
	}

	// Routes pour les administrateurs
	adminGroup := router.Group("/admin")
	{
		adminGroup.GET("/", middleware.GetAdminMiddleware(db), handlers.GetAdminsHandler(db))
		adminGroup.GET("/?id=:id", middleware.GetAdminByIDMiddleware(db), handlers.GetAdminByIDHandler(db))
	}

	login := router.Group("/login")
	{
		login.POST("/particulier", handlers.LoginParticulierHandler(db))
		login.POST("/salarie", handlers.LoginSalarieHandler(db))
	}
	return router
}