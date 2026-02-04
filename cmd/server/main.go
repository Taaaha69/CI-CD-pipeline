package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Taaaha69/CI-CD-pipeline/internal/server/api"
	"github.com/Taaaha69/CI-CD-pipeline/internal/server/config"
	"github.com/Taaaha69/CI-CD-pipeline/internal/server/db"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	database, err := db.Connect(cfg.Database.Path)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Migrate(database); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server running port", cfg.Server.Port)
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, router))
}
