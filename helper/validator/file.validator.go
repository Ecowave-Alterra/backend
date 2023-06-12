package validator

import (
	"mime/multipart"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func ValidateFileExtension(fileHeader *multipart.FileHeader) error {
	fileExtension := filepath.Ext(fileHeader.Filename)
	allowedExtensions := map[string]bool{
		".png":  true,
		".jpeg": true,
		".jpg":  true,
	}

	if !allowedExtensions[fileExtension] {
		return echo.NewHTTPError(400, "Mohon maaf format file yang anda unggah tidak sesuai")
	}

	return nil
}

func ValidateFileSize(fileHeader *multipart.FileHeader, maxFileSize int64) error {
	fileSize := fileHeader.Size
	if fileSize > maxFileSize {
		return echo.NewHTTPError(413, "Mohon maaf ukuran file Anda melebihi batas maksimum 4MB")
	}

	return nil
}
