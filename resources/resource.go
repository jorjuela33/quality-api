package resource

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (resource *Resource) List(context *gin.Context) {
	var products []model.Product
	resource.Database.DB().Table("alm_insumos").Scan(&products)
	context.JSON(http.StatusOK, gin.H{
		"status":   "Fine",
		"products": products,
	})
}
