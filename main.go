package main

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/jorjuela33/quality-api/resources"
	"github.com/jorjuela33/quality-api/server"
)

func main() {
	router := server.NewRouter()
	_server := server.NewServer()

	// setup resources
	productResource := resource.NewResource(&resource.Options{
		BasePath: resource.DefaultApiPath + "/products",
	})
	router.AddResources(productResource)

	// setup router
	_server.UseRouter(router)

	_server.Run(":8080", server.Options{
		Timeout: 10 * time.Second,
	})
}
