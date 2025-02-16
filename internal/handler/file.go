package handler

import (
	"fmt"
	"net/http"

	"github.com/duckcoding00/single-file/internal/service"
	"github.com/duckcoding00/single-file/lib/utils"
)

type FileHanlder struct {
	service service.Service
}

func (h *FileHanlder) SaveFile(w http.ResponseWriter, r *http.Request) {
	// Set the maximum file size to 2 MB
	const maxSize = 2 << 20 // 2 MB

	r.Body = http.MaxBytesReader(w, r.Body, maxSize)

	// Parse the multipart form
	if err := r.ParseMultipartForm(maxSize); err != nil {
		if err.Error() == "http: request body too large" {
			err = fmt.Errorf("file too large. Max file size is 2 MB")
			utils.WriteErr(w, http.StatusBadRequest, err)
			return
		}
		utils.WriteErr(w, http.StatusBadRequest, fmt.Errorf("failed to parse multipart form: %w", err))
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		utils.WriteErr(w, http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	resp, err := h.service.File.SaveFile(file, header)
	if err != nil {
		utils.WriteErr(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteOk(w, http.StatusCreated, resp)
}
