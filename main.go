package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jorjuela33/quality-api/router"

	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func main() {
	router := router.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}
