package controller

import (
	"github.com/gin-gonic/gin"
)

func GetLogin(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title": "Main website",
	})
}
