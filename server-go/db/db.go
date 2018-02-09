package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

func Open() {
	var err error
	Db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=paper sslmode=disable")
	if err != nil {
		log.Println("Connection Error" + err.Error())
		return
	}
}

func Close() {
	Db.Close()
}
