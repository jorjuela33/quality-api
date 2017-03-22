package renderer

import "github.com/gin-gonic/gin"

const JSON = "json"
const XML = "xml"
const Data = "octet-stream"
const Text = "text"

type Renderer struct {
	Type string
}

func New(renderType string) *Renderer {
	return &Renderer{Type: renderType}
}

func (renderer *Renderer) Render(context *gin.Context, status int) {
	switch renderer.Type {
	case XML:
		context.XML(status, gin.H{
			"status": "Fine",
			"foo":    "bar",
		})

	case Data:
		/// TO BE DEFINED

	case Text:
		context.String(status, "")

	default:
		context.JSON(status, gin.H{
			"status": "Fine",
			"foo":    "bar",
		})
	}
}
