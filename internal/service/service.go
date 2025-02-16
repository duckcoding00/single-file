package service

import "mime/multipart"

type Service struct {
	File interface {
		SaveFile(file multipart.File, header *multipart.FileHeader) (string, error)
	}
}

func NewService() Service {
	return Service{
		File: &FileService{},
	}
}
