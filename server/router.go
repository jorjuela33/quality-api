package server

import (
	"errors"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/jorjuela33/quality-api/domain"
)

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	router := mux.NewRouter().StrictSlash(true)
	return &Router{router}
}

func (router *Router) AddRoutes(routes *domain.Routes) *Router {
	if routes == nil {
		return router
	}

	for _, route := range *routes {
		defaultHandler, ok := route.RouteHandlers[route.DefaultVersion]
		if !ok {
			errors := errors.New(fmt.Sprintf("Routes definition error, missing default route handler for version `%v` in `%v`", route.DefaultVersion, route.Name))
			panic(errors)
		}

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name)
	}

	return router
}
