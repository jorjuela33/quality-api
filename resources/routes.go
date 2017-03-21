package resource

import "github.com/jorjuela33/quality-api/domain"

const DefaultApiPath string = "/api/products"

func (resource *Resource) createRoutes(path string) *domain.Routes {
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
				"0.0": resource.List,
			},
			Handler: resource.List,
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
	return resource.Routes
}
