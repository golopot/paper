package controller

import "github.com/gin-gonic/gin"

func ClientError(e error, c *gin.Context) {
	println(e)
	c.String(400, "Bad Request")
}
