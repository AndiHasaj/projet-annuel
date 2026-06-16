package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	host := "localhost"
	port := 3306
	user := "root"
	password := "MDPdbPA*67*"
	dbname := "upcycleconnect_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Erreur lors de l'ouverture de la base de données: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Erreur lors de la connexion à la base de données: %v", err)
		return nil, err
	}

	log.Println("Connexion à la base de données réussie")
	return db, nil
}