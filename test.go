package main

import (
	"github.com/NOX73/go-neural/persist"
	log "github.com/cihub/seelog"
)

func Test(netwrok string) {
	labelFile, imageFile, err := loadDataSet(testLabel, testImage)
	if err != nil {
		log.Error(err)
		return
	}
	network := persist.FromFile(network)
	//maxTestSet := 1000
	maxTestSet := labelFile.Num

	accurateNum := 0
	var resultNum int
	var resultRate float64

	for i := 0; i < maxTestSet; i++ {
		resultNum = -1
		resultRate = 0.0

		test := make([]float64, 0)
		for _, v := range imageFile.GetImage(i) {
			test = append(test, float64(v))
		}

		result := network.Calculate(test)
		for i, v := range result {
			if v > resultRate {
				resultNum = i
				resultRate = v
			}
			//log.Infof("[%d]:%f", i, v)
		}
		//log.Infof("[%d] %d -> %d, possible:%f", i, labelFile.Label[i], resultNum, resultRate)
		if int(labelFile.Label[i]) == resultNum {
			accurateNum++
		}
	}
	log.Infof("Accurate:%d, Rate:%f", accurateNum, float64(accurateNum)/float64(maxTestSet))
}
