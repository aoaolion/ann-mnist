package main

import (
	"fmt"

	neural "github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/learn"
	"github.com/NOX73/go-neural/persist"
	log "github.com/cihub/seelog"
)

func InData(imageFile *ImageFile) [][]float64 {
	inList := make([][]float64, 0)
	for i := 0; i < imageFile.Num; i++ {
		in := make([]float64, 0)
		buf := imageFile.GetImage(i)
		for _, v := range buf {
			in = append(in, float64(v))
		}
		inList = append(inList, in)
	}
	return inList
}

func IdealData(labelFile *LabelFile) [][]float64 {
	idealList := make([][]float64, 0)
	for i := 0; i < labelFile.Num; i++ {
		ideal := make([]float64, 0)
		for j := 0; j < 10; j++ {
			if j == int(labelFile.Label[i]) {
				ideal = append(ideal, 1)
			} else {
				ideal = append(ideal, 0)
			}
		}
		idealList = append(idealList, ideal)
	}
	return idealList

}

func Train(maxIteration, maxSetSize int, speed float64) {
	labelFile, imageFile, err := loadDataSet(trainLabel, trainImage)
	if err != nil {
		log.Error(err)
		return
	}
	// 28*28 = 784 inputs, 3 layers
	network := neural.NewNetwork(784, []int{300, 100, 10})
	network.RandomizeSynapses()

	if maxSetSize > imageFile.Num {
		maxSetSize = imageFile.Num
	}

	log.Infof("maxIteration: %d, trainSet: %d", maxIteration, maxSetSize)

	idealList := IdealData(labelFile)
	inList := InData(imageFile)

	for iteration := 0; iteration < maxIteration; iteration++ {
		avge := 0.0
		for i := 0; i < maxSetSize; i++ {
			//s := time.Now()

			learn.Learn(network, inList[i], idealList[i], speed)
			e := learn.Evaluation(network, inList[i], idealList[i])
			avge += e

			//log.Infof("iteration:%d, training:%d, estimate:%f", iteration, i, estimate)
		}
		avge = avge / float64(maxSetSize)
		log.Infof("%5d, %f", iteration, avge)
		path := fmt.Sprintf("data/network_%d_%f.json", iteration, avge)
		persist.ToFile(path, network)
		if avge < 0.001 {
			break
		}
	}
}
