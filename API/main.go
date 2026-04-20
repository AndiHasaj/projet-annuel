package main

import (
	"fmt"
	"API/router"
	"API/db"
	"log"
)



func main() {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Erreur de connexion à la base de données: %v", err)
	}
	router := router.SetupRouter(db)
	router.Run(":8080")
	fmt.Println("Serveur démarré sur le port 8080")

}