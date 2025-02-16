package service

import "mime/multipart"

type Service struct {
	File interface {
		SaveFile(file multipart.File, header *multipart.FileHeader) (string, error)
		GetFiles() ([]string, error)
		GetFile(filename string) (string, string, error)
	}
}

func NewService() Service {
	return Service{
		File: &FileService{},
	}
}
