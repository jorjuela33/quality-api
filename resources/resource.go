package resource

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jorjuela33/quality-api/database"
	"github.com/jorjuela33/quality-api/domain"
	"github.com/jorjuela33/quality-api/models/product"
)

type Products []model.Product

type Response_v0 struct {
	Products Products `json:"products"`
	Success  bool     `json:"success"`
}

type ResourceInterface interface {
	//Context() IContext
	Routes() *domain.Routes
}

type Resource struct {
	Routes   *domain.Routes
	Database database.DatabaseInterface
	Renderer domain.RendererInterface
}

type Options struct {
	BasePath string
	Database database.DatabaseInterface
	Renderer domain.RendererInterface
}

func NewResource(options *Options) *Resource {
	resource := &Resource{nil, options.Database, options.Renderer}
	resource.createRoutes(options.BasePath)
	return resource
}

func (resource *Resource) List(context *gin.Context) {
	var products []model.Product
	resource.Database.DB().Table("alm_insumos").Scan(&products)
	resource.Renderer.JSON(context, http.StatusOK, products)
}
