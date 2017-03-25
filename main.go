package main

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/jorjuela33/quality-api/database"
	"github.com/jorjuela33/quality-api/middlewares/renderer"
	"github.com/jorjuela33/quality-api/mssql"
	"github.com/jorjuela33/quality-api/resources"
	"github.com/jorjuela33/quality-api/server"
)

func main() {
	_server := server.NewServer(&server.Options{
		Port: ":8080",
	})

	// setup database connection
	database := mssql.New(&database.Options{
		ServerName:   "181.49.12.194",
		DatabaseName: "BD_TEMP",
	})
	_ = database.NewSession()

	// setup renderer
	renderer := renderer.New(renderer.JSON)

	// setup resources
	productResource := resource.NewResource(&resource.Options{
		BasePath: resource.DefaultApiPath + "/products",
		Database: database,
		Renderer: renderer,
	})
	_server.AddResources(productResource)

	_server.Run(":8080")
}
