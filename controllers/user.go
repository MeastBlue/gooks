package controllers

import (
	"books/handlers"
	"books/services"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := services.GetUsers()
	if users == nil {
		return handlers.SendJsonNoContent(w, users)
	}

	return handlers.SendJsonOK(w, &users)
}
