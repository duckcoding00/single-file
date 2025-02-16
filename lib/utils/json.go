package utils

import (
	"encoding/json"
	"net/http"
)

type RespErr struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type RespOk struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
}

func WriteErr(w http.ResponseWriter, statusCode int, err error) error {
	var message string
	if statusCode == http.StatusBadRequest {
		message = "BAD REQUEST"
	}

	if statusCode == http.StatusInternalServerError {
		message = "INTERNAL SERVER ERROR"
	}

	return writeJSON(w, statusCode, RespErr{
		Message: message,
		Error:   err.Error(),
	})
}

func WriteOk(w http.ResponseWriter, statusCode int, data interface{}) error {
	var message string
	if statusCode == http.StatusOK {
		message = "OK"
	}

	if statusCode == http.StatusCreated {
		message = "Created"
	}

	return writeJSON(w, statusCode, RespOk{
		Message: message,
		Data:    data,
	})
}
