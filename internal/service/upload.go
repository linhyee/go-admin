package service

import (
	"blog/global"
	"blog/pkg/upload"
	"errors"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func UploadFile(fileType upload.FileType, file multipart.File,
	fileHeader *multipart.FileHeader) (*FileInfo, error) {

	fileName := upload.GetFileName(fileHeader.Filename)
	savePath := upload.GetSavePath()
	dst := savePath + "/" + fileName
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if upload.CheckSavePath(savePath) {
		err := upload.CreateSavePath(savePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}
	if upload.CheckPermission(savePath) {
		return nil, errors.New("insufficient file permissions")
	}
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
