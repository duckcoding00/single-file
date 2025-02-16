package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type FileService struct {
}

func (s *FileService) validatingFile(header *multipart.FileHeader) error {
	allowedExtension := map[string]bool{
		".jpg":  true,
		".png":  true,
		".jpeg": true,
		".gif":  true,
	}
	// check extension
	fileExtension := strings.ToLower(filepath.Ext(header.Filename))
	if !allowedExtension[fileExtension] {
		return fmt.Errorf("only images extension allowed (jpg, jpeg, png, gif)")
	}

	// // validation MIME
	// buffer := make([]byte, 512)
	// _, err := file.Read(buffer)
	// if err != nil {
	// 	return fmt.Errorf("unable read file")
	// }
	// file.Seek(0, 0)

	// mimeType := http.DetectContentType(buffer)
	// if !strings.HasPrefix(mimeType, "images/") {
	// 	return fmt.Errorf("invalid type, only images are allowed")
	// }

	return nil
}

func (s *FileService) SaveFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	// create folder if not didnt exits
	if err := os.MkdirAll("data", os.ModePerm); err != nil {
		return "", err
	}

	// check validation file
	if err := s.validatingFile(header); err != nil {
		return "", err
	}

	// create new file
	filePath := filepath.Join("data", header.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
