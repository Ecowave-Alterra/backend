package validator

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
)

func ValidateVideoExtension(videoHeader *multipart.FileHeader) error {
	videoExtension := filepath.Ext(videoHeader.Filename)
	allowedExtensions := map[string]bool{
		".mp4": true,
	}

	if !allowedExtensions[videoExtension] {
		//lint:ignore ST1005 Reason for ignoring this linter
		return fmt.Errorf("Mohon maaf, format video yang anda unggah tidak sesuai")
	}

	return nil
}

func ValidateVideoSize(videoHeader *multipart.FileHeader, maxVideoSize int64) error {
	videoSize := videoHeader.Size
	if videoSize > maxVideoSize {
		//lint:ignore ST1005 Reason for ignoring this linter
		return fmt.Errorf("Mohon maaf, ukuran video Anda melebihi batas maksimum 4MB")
	}

	return nil
}
