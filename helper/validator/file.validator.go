package validator

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
)

func ValidateFileExtension(fileHeader *multipart.FileHeader) error {
	fileExtension := filepath.Ext(fileHeader.Filename)
	allowedExtensions := map[string]bool{
		".png":  true,
		".jpeg": true,
		".jpg":  true,
	}

	if !allowedExtensions[fileExtension] {
		//lint:ignore ST1005 Reason for ignoring this linter
		return fmt.Errorf("Mohon maaf, format file yang anda unggah tidak sesuai")
	}

	return nil
}

func ValidateFileSize(fileHeader *multipart.FileHeader, maxFileSize int64) error {
	fileSize := fileHeader.Size
	if fileSize > maxFileSize {
		//lint:ignore ST1005 Reason for ignoring this linter
		return fmt.Errorf("Mohon maaf, ukuran file Anda melebihi batas maksimum 4MB")
	}

	return nil
}
