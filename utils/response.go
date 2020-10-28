package utils

import (
	"books/models"
)

func FormatJsonResponse(code int, data interface{}) models.Response {
	res := models.Response{}
	res.Status = code
	res.Data = data

	return res
}
