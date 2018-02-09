package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRoot(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title": "Main website",
	})
}
