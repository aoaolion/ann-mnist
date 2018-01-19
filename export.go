package main

import (
	"fmt"

	log "github.com/cihub/seelog"
)

func Export(export string, maxExportSetSize int) {

	labelFile, imageFile, err := loadDataSet(testLabel, testImage)
	if export == "train" {
		labelFile, imageFile, err = loadDataSet(trainLabel, trainImage)
	}
	if err != nil {
		log.Error(err)
		return
	}

	if maxExportSetSize > imageFile.Num {
		maxExportSetSize = imageFile.Num
	}

	for i := 0; i < maxExportSetSize; i++ {
		path := fmt.Sprintf("pic/t_%d_%d.png", int(labelFile.Label[i]), i)
		imageFile.SaveImage(path, i)
		log.Infof("save %s", path)
	}
}
