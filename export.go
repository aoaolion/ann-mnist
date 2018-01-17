package main

import (
	"fmt"

	log "github.com/cihub/seelog"
)

func Export() {
	//labelFile, imageFile, err := loadDataSet(trainLabel, trainImage)
	labelFile, imageFile, err := loadDataSet(testLabel, testImage)
	if err != nil {
		log.Error(err)
		return
	}
	for i := 0; i < imageFile.Num; i++ {
		path := fmt.Sprintf("pic/t_%d_%d.png", int(labelFile.Label[i]), i)
		imageFile.SaveImage(path, i)
		log.Infof("save %s", path)
	}
}
