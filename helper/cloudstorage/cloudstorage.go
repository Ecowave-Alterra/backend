package cloudstorage

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"path"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

var Folder string

func UploadToBucket(ctx context.Context, fileHeader *multipart.FileHeader) (string, error) {
	bucket := "ecowave"

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("storage.json"))
	if err != nil {
		return "", err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", echo.NewHTTPError(500, err)
	}
	defer file.Close()

	objectName := Folder + fileHeader.Filename
	sw := storageClient.Bucket(bucket).Object(objectName).NewWriter(ctx)

	if _, err := io.Copy(sw, file); err != nil {
		return "", echo.NewHTTPError(500, err)
	}

	if err := sw.Close(); err != nil {
		return "", echo.NewHTTPError(500, err)
	}

	u, err := url.Parse("/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		return "", echo.NewHTTPError(500, err)
	}

	PhotoUrl := fmt.Sprintf("https://storage.googleapis.com%s", u.EscapedPath())
	return PhotoUrl, nil
}

func GetFileName(filePath string) string {
	decodeFilePath, err := url.PathUnescape(filePath)
	if err != nil {
		return ""
	}

	fileName := path.Base(decodeFilePath)
	return fileName
}

func DeleteImage(fileName string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("storage.json"))
	if err != nil {
		return echo.NewHTTPError(500, err)
	}

	bucketName := "ecowave"
	objectPath := Folder + fileName

	obj := client.Bucket(bucketName).Object(objectPath)

	err = obj.Delete(ctx)
	if err != nil {
		return echo.NewHTTPError(500, "Gagal menghapus file pada cloud storage")
	}

	return nil
}
