package service

import (
	"errors"
	"mime/multipart"
	"os"

	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/global"
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessURL string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName

	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported.")
	}

	if upload.CheckSavePath(uploadSavePath) {
		err := upload.CreateSavePath(uploadSavePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}

	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}

	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}

	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessURL := global.AppSetting.UploadServerURL + "/" + fileName
	return &FileInfo{
		Name:      fileName,
		AccessURL: accessURL,
	}, nil
}
