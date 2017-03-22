package server

/*import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/jorjuela33/quality-api/domain"
	"github.com/jorjuela33/quality-api/resources"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	router := mux.NewRouter().StrictSlash(true)
	return &Router{router}
}

func (router *Router) AddResources(resources ...*resource.Resource) *Router {
	for _, resource := range resources {
		if resource.Routes == nil {
			panic(errors.New(fmt.Sprintf("Routes definition missing: %v", resource)))
		}

		router.AddRoutes(resource.Routes)
	}

	return router
}

func (router *Router) AddRoutes(routes *domain.Routes) *Router {
	if routes == nil {
		return router
	}

	for _, route := range *routes {
		_, ok := route.RouteHandlers[route.DefaultVersion]
		if !ok {
			errors := errors.New(fmt.Sprintf("Routes definition error, missing default route handler for version `%v` in `%v`", route.DefaultVersion, route.Name))
			panic(errors)
		}

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).HandlerFunc(route.Handler)
	}

	return router
}*/
