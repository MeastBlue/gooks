package handlers

import (
	"books/utils"
	"encoding/json"
	"net/http"
)

func sendJsonWithStatus(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(status)

	if status == http.StatusNoContent {
		return
	}

	if err := json.NewEncoder(w).Encode(&data); err != nil {
		panic(err)
	}
}

func SendJsonOK(w http.ResponseWriter, data interface{}) {
	formatData := utils.FormatJsonResponse(http.StatusOK, &data)
	sendJsonWithStatus(w, formatData, http.StatusOK)
}

func SendJsonNotFound(w http.ResponseWriter, data interface{}) {
	formatData := utils.FormatJsonResponse(http.StatusNotFound, nil)
	sendJsonWithStatus(w, formatData, http.StatusNotFound)
}

func SendJsonNoContent(w http.ResponseWriter, data interface{}) {
	formatData := utils.FormatJsonResponse(http.StatusNoContent, nil)
	sendJsonWithStatus(w, formatData, http.StatusNotFound)
}
