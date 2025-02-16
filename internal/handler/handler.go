package handler

import (
	"net/http"

	"github.com/duckcoding00/single-file/internal/service"
)

type Handler struct {
	File interface {
		SaveFile(w http.ResponseWriter, r *http.Request)
	}
}

func NewHandler() Handler {
	service := service.NewService()
	return Handler{
		File: &FileHanlder{
			service: service,
		},
	}
}
