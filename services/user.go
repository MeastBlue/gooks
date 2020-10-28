package services

import (
	"books/config"
	"books/models"
	"fmt"
	"log"
)

func GetUsers() models.Users  {
	users := models.Users{}
	query := fmt.Sprintf("select * from user")
	db, err := config.DatabaseConnect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	if err = db.Select(&users, query); err != nil {
		log.Fatal(err)
	}

	return users
}
