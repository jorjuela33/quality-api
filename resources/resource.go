package resource

import (
	"encoding/json"
	"net/http"

	"github.com/jorjuela33/quality-api/database"
	"github.com/jorjuela33/quality-api/domain"
	"github.com/jorjuela33/quality-api/models/product"
)

type ResourceInterface interface {
	//Context() IContext
	Routes() *domain.Routes
}

type Resource struct {
	Routes   *domain.Routes
	Database database.DatabaseInterface
}

type Options struct {
	BasePath string
	Database database.DatabaseInterface
}

func NewResource(options *Options) *Resource {
	database := options.Database
	resource := &Resource{nil, database}
	resource.createRoutes(options.BasePath)
	return resource
}

func (resource *Resource) List(responseWriter http.ResponseWriter, request *http.Request) {
	var products []model.Product
	resource.Database.DB().Table("alm_insumos").Scan(&products)
	json.NewEncoder(responseWriter).Encode(products)
}
