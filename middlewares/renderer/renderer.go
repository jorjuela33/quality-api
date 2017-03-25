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

func (renderer *Renderer) Render(context *gin.Context, status int, v interface{}) {
	switch renderer.Type {
	case XML:
		renderer.XML(context, status, v)

	case Data:
		renderer.Data(context, status, v.([]byte))

	case Text:
		renderer.Text(context, status, v.([]byte))

	default:
		renderer.JSON(context, status, v)
	}
}

func (renderer *Renderer) Data(context *gin.Context, status int, v []byte) {
	context.Data(status, "contentType", v)
}

func (renderer *Renderer) JSON(context *gin.Context, status int, v interface{}) {
	context.JSON(status, gin.H{
		"status": "Fine",
		"foo":    v,
	})
}

func (renderer *Renderer) Text(context *gin.Context, status int, v []byte) {
	context.String(status, "")
}

func (renderer *Renderer) XML(context *gin.Context, status int, v interface{}) {
	context.XML(status, gin.H{
		"status": "Fine",
		"foo":    v,
	})
}
