package main

import (
	log "github.com/cihub/seelog"
)

func loadDataSet(lablePath, imgPath string) (*LabelFile, *ImageFile, error) {
	// load training set
	labelFile, err := NewLabel(lablePath)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}
	//log.Info(labelFile)
	imageFile, err := NewImageFile(imgPath)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}
	//log.Info(imageFile)
	return labelFile, imageFile, nil
}
