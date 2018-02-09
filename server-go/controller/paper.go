package controller

import (
	"github.com/gin-gonic/gin"
)

func GetPaper(c *gin.Context) {
	c.HTML(200, "index.tmpl", nil)
}

func GetPaperID(c *gin.Context) {
	c.HTML(200, "index.tmpl", nil)
}
