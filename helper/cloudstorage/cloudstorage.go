package cloudstorage

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func UploadToBucket(ctx context.Context, fileHeader *multipart.FileHeader) (string, error) {
	bucket := "ecowave_storage"

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("storage.json"))
	if err != nil {
		return "", err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	objectName := "img/" + fileHeader.Filename
	sw := storageClient.Bucket(bucket).Object(objectName).NewWriter(ctx)

	if _, err := io.Copy(sw, file); err != nil {
		return "", err
	}

	if err := sw.Close(); err != nil {
		return "", err
	}

	u, err := url.Parse("/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		return "", err
	}

	PhotoUrl := fmt.Sprintf("https://storage.cloud.google.com%s", u.EscapedPath())
	return PhotoUrl, nil
}
