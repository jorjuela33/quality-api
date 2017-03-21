package resource

import (
	"encoding/json"
	"net/http"

	"github.com/jorjuela33/quality-api/domain"
	"github.com/jorjuela33/quality-api/models/product"
	"github.com/jorjuela33/quality-api/mssql"
)

type ResourceInterface interface {
	//Context() IContext
	Routes() *domain.Routes
}

type Resource struct {
	Routes   *domain.Routes
	Database *mssql.MSSQLDB
}

type Options struct {
	BasePath string
	Database *mssql.MSSQLDB
}

func NewResource(options *Options) *Resource {
	database := options.Database
	path := options.BasePath
	resource := &Resource{nil, database}

	/*if database == nil {
		panic("sessions.Options.Database is required")
	}*/

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
	return resource
}

func (resource *Resource) List(responseWriter http.ResponseWriter, request *http.Request) {
	var products []model.Product
	resource.Database.GormDB.Table("alm_insumos").Scan(&products)
	json.NewEncoder(responseWriter).Encode(products)
}
