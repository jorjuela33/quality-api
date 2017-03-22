package server

import (
	"errors"
	"fmt"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/jorjuela33/quality-api/domain"
	resource "github.com/jorjuela33/quality-api/resources"
)

type Server struct {
	router  *gin.Engine
	options *Options
}

// Options for running the server
type Options struct {
	Port string
}

func NewServer(options *Options) *Server {
	router := gin.Default()
	server := &Server{router, options}
	return server
}

func (server *Server) AddResources(resources ...*resource.Resource) *Server {
	for _, resource := range resources {
		if resource.Routes == nil {
			panic(errors.New(fmt.Sprintf("Routes definition missing: %v", resource)))
		}

		server.AddRoutes(resource.Routes)
	}

	return server
}

func (server *Server) AddRoutes(routes *domain.Routes) *Server {
	if routes == nil {
		return server
	}

	for _, route := range *routes {
		_, ok := route.RouteHandlers[route.DefaultVersion]
		if !ok {
			errors := errors.New(fmt.Sprintf("Routes definition error, missing default route handler for version `%v` in `%v`", route.DefaultVersion, route.Name))
			panic(errors)
		}

		switch route.Method {
		case "GET":
			server.router.GET(route.Pattern, route.Handler)

		case "DELETE":
			server.router.DELETE(route.Pattern, route.Handler)

		case "POST":
			server.router.POST(route.Pattern, route.Handler)

		case "PUT":
			server.router.PUT(route.Pattern, route.Handler)
		}
	}

	return server
}

func (server *Server) Stop() {
	endless.ListenAndServe(server.options.Port, server.router)
}

func (server *Server) Run(address string) *Server {
	server.router.Run(server.options.Port)
	return server
}
