package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type FileService struct {
}

func (s *FileService) createFolder() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working dirr: %w", err)
	}

	rootDir := filepath.Join(currentDir, "..")
	dataDir := filepath.Join(rootDir, "data")
	return dataDir, nil
}

func (s *FileService) validatingFile(file multipart.File, header *multipart.FileHeader) error {
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

	// validation MIME
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return fmt.Errorf("unable read file")
	}
	file.Seek(0, 0)

	mimeType := http.DetectContentType(buffer)
	if !strings.HasPrefix(mimeType, "image/") {
		return fmt.Errorf("invalid type, only image are allowed")
	}

	return nil
}

func (s *FileService) createUniqueFilename(dir, fileName string) string {
	ext := filepath.Ext(fileName)
	nameWithoutExt := strings.TrimSuffix(fileName, ext)
	nameWithoutExt = strings.ReplaceAll(nameWithoutExt, " ", "_")

	count := 1

	fileName = strings.ReplaceAll(fileName, " ", "_")
	newFilename := fileName
	for {
		filePath := filepath.Join(dir, newFilename)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return filePath
		}
		newFilename = fmt.Sprintf("%s(%d)%s", nameWithoutExt, count, ext)
		count++
	}
}

func (s *FileService) SaveFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	// create folder if not didnt exits
	dataDir, err := s.createFolder()
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		return "", err
	}

	// check validation file
	if err := s.validatingFile(file, header); err != nil {
		return "", err
	}

	// create new file
	filePath := s.createUniqueFilename(dataDir, header.Filename)

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

func (s *FileService) GetFiles() ([]string, error) {
	dir, err := s.createFolder()
	if err != nil {
		return nil, err
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				return nil, err
			}
			return []string{}, nil
		}
		return nil, err
	}
	var images []string
	for _, file := range files {
		images = append(images, file.Name())
	}

	return images, nil
}

func (s *FileService) GetFile(filename string) (string, string, error) {
	mimeType := map[string]string{
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".png":  "image/png",
		".gif":  "image/png",
	}

	dir, err := s.createFolder()
	if err != nil {
		return "", "", err
	}

	filePath := filepath.Join(dir, filename)
	//log.Print(filePath)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", "", fmt.Errorf("file didnt exist")
	}

	contentType := mimeType[strings.ToLower(filepath.Ext(filePath))]
	//log.Println(contentType)

	return filePath, contentType, nil
}
