package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeWorkers interface {
	IndexHome(c *gin.Context)
	//UpdateHome()
}

type homeWorkers struct {
	PageName string
}

func NewHomeWorkers() HomeWorkers {
	return &homeWorkers{
		PageName: "test",
	}
}

func (h homeWorkers) IndexHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}

//func (h *homeWorkers) UpdateHome()
