package api

import (
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func randStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func PostUpload(c *gin.Context) {
	file, _ := c.FormFile("file")
	fileID := randStringBytes(16)
	dst := "./upload/" + fileID

	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(400, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	c.JSON(200, gin.H{
		"fileID": fileID,
	})
}
