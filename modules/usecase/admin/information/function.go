package information

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
	"github.com/berrylradianh/ecowave-go/helper/randomid"
	vld "github.com/berrylradianh/ecowave-go/helper/validator"
	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
)

func (ic *informationUsecase) GetAllInformationsNoPagination() (*[]ie.Information, error) {
	informations, err := ic.informationRepo.GetAllInformationsNoPagination()
	return informations, err
}

func (ic *informationUsecase) GetAllInformations(offset, pageSize int) (*[]ie.Information, int64, error) {
	informations, count, err := ic.informationRepo.GetAllInformations(offset, pageSize)
	return informations, count, err
}

func (ic *informationUsecase) GetInformationById(informationId string) (*ie.Information, error) {
	information, err := ic.informationRepo.GetInformationById(informationId)
	return information, err
}

func (ic *informationUsecase) CreateInformation(fileHeader *multipart.FileHeader, title, content, status string) error {
	if fileHeader == nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return errors.New("Mohon maaf, anda harus mengunggah file")
	}

	if err := vld.ValidateFileExtension(fileHeader); err != nil {
		return err
	}

	maxFileSize := 4 * 1024 * 1024
	if err := vld.ValidateFileSize(fileHeader, int64(maxFileSize)); err != nil {
		return err
	}

	PhotoUrl, err := cloudstorage.UploadToBucket(context.Background(), fileHeader)
	if err != nil {
		return err
	}

	information := &ie.Information{
		Title:           title,
		Content:         content,
		PhotoContentUrl: PhotoUrl,
		Status:          status,
	}

	if err := vld.Validation(information); err != nil {
		return err
	}

	for {
		informationId := randomid.GenerateRandomID()

		exists, err := ic.informationRepo.CheckInformationExists(informationId)
		if err != nil {
			return err
		}
		if !exists {
			information.InformationId = informationId
			break
		}
	}

	err = ic.informationRepo.CreateInformation(information, nil)
	return err
}

func (ic *informationUsecase) CreateInformationDraft(fileHeader *multipart.FileHeader, title, content, status string) error {
	var information ie.InformationDraftRequest

	information.Status = status

	if fileHeader == nil && title == "" && content == "" {
		//lint:ignore ST1005 Reason for ignoring this linter
		return fmt.Errorf("Masukkan data")
	}

	if title == "" {
		information.Title = ""
	} else {
		information.Title = title
	}

	if content == "" {
		information.Content = ""
	} else {
		information.Content = content
	}

	if fileHeader == nil {
		information.PhotoContentUrl = ""
	} else {
		if err := vld.ValidateFileExtension(fileHeader); err != nil {
			return err
		}

		maxFileSize := 4 * 1024 * 1024
		if err := vld.ValidateFileSize(fileHeader, int64(maxFileSize)); err != nil {
			return err
		}

		PhotoUrl, err := cloudstorage.UploadToBucket(context.Background(), fileHeader)
		if err != nil {
			return err
		}

		information.PhotoContentUrl = PhotoUrl
	}

	for {
		informationId := randomid.GenerateRandomID()

		exists, err := ic.informationRepo.CheckInformationExists(informationId)
		if err != nil {
			return err
		}
		if !exists {
			information.InformationId = informationId
			break
		}
	}

	err := ic.informationRepo.CreateInformation(nil, &information)
	return err
}

func (ic *informationUsecase) UpdateInformation(informationId string, information *ie.Information) error {
	if err := vld.Validation(information); err != nil {
		return err
	}

	result := ic.informationRepo.UpdateInformation(informationId, information)
	return result
}

func (ic *informationUsecase) DeleteInformation(informationId string) error {
	err := ic.informationRepo.DeleteInformation(informationId)
	return err
}

func (ic *informationUsecase) SearchInformations(search, filter string, offset, pageSize int) (*[]ie.Information, int64, error) {
	informations, count, err := ic.informationRepo.SearchInformations(search, filter, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return informations, count, nil
}
