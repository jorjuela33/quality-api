package main

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/jorjuela33/quality-api/server"
)

func main() {
	router := server.NewRouter()
	_server := server.NewServer()

	/// the router to use by the server
	_server.UserRouter(router)

	_server.Run(":8080", server.Options{
		Timeout: 10 * time.Second,
	})
}
