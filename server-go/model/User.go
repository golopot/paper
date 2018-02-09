package model

import (
	"paper/server-go/db"
)

func CreateUserByGoogleId(google_id string) (User, error) {
	user := User{
		Google_ID: google_id,
	}
	err := db.Db.Create(&user).Error
	return user, err
}

func FindUserByGoogleID(google_id string) (User, error) {
	user := User{}
	err := db.Db.Take(&user, "Google_ID = ?", google_id).Error
	return user, err
}

func FindOrCreateUserByGoogleId(google_id string) (User, error) {
	user, err := FindUserByGoogleID(google_id)
	if user == (User{}) {
		return CreateUserByGoogleId(google_id)
	}
	if err != nil {
		return user, err
	}

	return user, err

}

type User struct {
	ID        uint   `gorm:"primary_key"`
	Google_ID string `gorm:"type:varchar(100)"`
}
