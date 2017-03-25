package domain

import "github.com/gin-gonic/gin"

type RendererInterface interface {
	Render(context *gin.Context, status int, v interface{})
	JSON(context *gin.Context, status int, v interface{})
	XML(context *gin.Context, status int, v interface{})
	Data(context *gin.Context, status int, v []byte)
	Text(context *gin.Context, status int, v []byte)
}
