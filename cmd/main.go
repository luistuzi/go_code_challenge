package main

import (
	"database/sql"
	"go_code_challenge/cmd/api"
	"go_code_challenge/config"
	"go_code_challenge/repository"
	"log"

	"github.com/go-sql-driver/mysql"
)

// @title Device API
// @version 1.0
// @description API used for device management in this code challenge
// @host localhost:8080
// @BasePath /api/v1/device
func main() {

	db, err := repository.NewSQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database successfully")
}
