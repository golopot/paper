package main

import (
	"log"
	"math/rand"
	_ "net/http"
	"paper/server-go/api"
	"paper/server-go/controller"
	"paper/server-go/db"
	"paper/server-go/model"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func serverError(c gin.Context, err error) {
	c.String(500, "Server Error")
	log.Println(err.Error())
}

func main() {

	rand.Seed(time.Now().UnixNano())

	db.Open()
	defer db.Close()

	db.Db.AutoMigrate(&model.Paper{})
	db.Db.AutoMigrate(&model.User{})

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "../client/dist/")

	router.GET("/", controller.GetRoot)

	router.GET("/login", controller.GetLogin)

	router.GET("/paper", controller.GetPaper)

	router.GET("/oauth-google", controller.OauthGoogle)

	router.GET("/oauth-callback", controller.OauthCallback)

	router.GET("/paper/:ID", controller.GetPaperID)

	router.POST("/api/upload", api.PostUpload)

	router.GET("/api/paper", api.GetPaper)

	router.GET("/api/paper/:ID", api.GetPaperID)

	router.POST("/api/paper", api.PostPaper)

	router.Run()
}
