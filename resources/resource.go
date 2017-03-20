package resource

import (
	"github.com/jorjuela33/quality-api/controllers"
	"github.com/jorjuela33/quality-api/domain"
)

type ResourceInterface interface {
	//Context() IContext
	Routes() *domain.Routes
}

type Resource struct {
	Routes *domain.Routes
}

type Options struct {
	BasePath string
}

func NewResource(options *Options) *Resource {
	path := options.BasePath
	resource := &Resource{nil}

	if path == "" {
		path = DefaultApiPath + "/products"
	}

	var routes = domain.Routes{

		domain.Route{
			Name:           "products",
			Method:         "GET",
			Pattern:        "/api/products",
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": controller.Index,
			},
			Handler: controller.Index,
		},
	}

	for _, route := range routes {
		r := domain.Route{
			Name:           route.Name,
			Method:         route.Method,
			Pattern:        path,
			DefaultVersion: route.DefaultVersion,
			RouteHandlers:  route.RouteHandlers,
			Handler:        route.Handler,
		}

		routes = routes.Append(&domain.Routes{r})
	}

	resource.Routes = &routes
	return resource
}
