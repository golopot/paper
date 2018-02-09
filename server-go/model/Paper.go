package model

import (
	"paper/server-go/db"
)

func PaperCreate(paper Paper) {
	db.Db.Create(&paper)
}

type Paper struct {
	ID       uint `gorm:"primary_key"`
	UserID   string
	Title    string `gorm:"type:varchar(100)"`
	Authors  string `gorm:"type:varchar(100)"`
	Abstract string `gorm:"type:varchar(100)"`
}
