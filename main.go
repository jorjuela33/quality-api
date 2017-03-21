package main

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/jorjuela33/quality-api/database"
	"github.com/jorjuela33/quality-api/mssql"
	"github.com/jorjuela33/quality-api/resources"
	"github.com/jorjuela33/quality-api/server"
)

func main() {
	router := server.NewRouter()
	_server := server.NewServer()

	// setup database connection
	database := mssql.New(&database.Options{
		ServerName:   "181.49.12.194",
		DatabaseName: "BD_TEMP",
	})
	_ = database.NewSession()

	// setup resources
	productResource := resource.NewResource(&resource.Options{
		BasePath: resource.DefaultApiPath + "/products",
		Database: database,
	})
	router.AddResources(productResource)

	// setup router
	_server.UseRouter(router)

	_server.Run(":8080", server.Options{
		Timeout: 10 * time.Second,
	})
}
