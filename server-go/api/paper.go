package api

import (
	"encoding/json"
	"log"
	"paper/server-go/db"
	"paper/server-go/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPaper(c *gin.Context) {
	var papers []model.Paper
	db.Db.Find(&papers)
	c.JSON(200, papers)
}

func GetPaperID(c *gin.Context) {

	ID, err := strconv.ParseUint(c.Param("ID"), 10, 32)
	if err != nil {
		panic("")
	}
	paper := model.Paper{
		ID: uint(ID),
	}

	db.Db.Find(&paper)
	c.JSON(200, paper)
}

func PostPaper(c *gin.Context) {

	var b struct {
		Title    string
		Authors  string
		Abstract string
		FileID   string
	}

	err := json.NewDecoder(c.Request.Body).Decode(&b)
	if err != nil {
		panic(err)
	}
	defer c.Request.Body.Close()

	log.Println(b)

	paper := model.Paper{
		UserID:   "",
		Title:    b.Title,
		Authors:  b.Authors,
		Abstract: b.Abstract,
	}

	model.PaperCreate(paper)

	c.JSON(200, gin.H{
		"ID": paper.ID,
	})
}
