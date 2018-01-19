package main

import (
	"fmt"

	neural "github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/learn"
	"github.com/NOX73/go-neural/persist"
	log "github.com/cihub/seelog"
)

func InData(imageFile *ImageFile, idx int, in *[]float64) {
	buf := imageFile.GetImage(idx)
	for _, v := range buf {
		*in = append(*in, float64(v))
	}
}

func IdealData(labelFile *LabelFile, idx int, ideal *[]float64) {
	for i := 0; i < 10; i++ {
		if i == int(labelFile.Label[idx]) {
			*ideal = append(*ideal, 1)
		} else {
			*ideal = append(*ideal, 0)
		}
	}
}

func Train(maxIteration, maxSetSize int) {
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

	in := make([]float64, 0)
	ideal := make([]float64, 0)

	for iteration := 0; iteration < maxIteration; iteration++ {
		avg := 0.0
		for i := 0; i < maxSetSize; i++ {
			//s := time.Now()
			in = in[:0]
			ideal = ideal[:0]

			InData(imageFile, i, &in)
			IdealData(labelFile, i, &ideal)

			learn.Learn(network, in, ideal, 0.2)
			estimate := learn.Evaluation(network, in, ideal)
			avg += estimate

			//log.Infof("iteration:%d, training:%d, estimate:%f", iteration, i, estimate)
		}
		avg = avg / float64(maxSetSize)
		log.Infof("iteration:%5d, e:%f", iteration, avg)
		path := fmt.Sprintf("data/network_%d_%f.json", iteration, avg)
		persist.ToFile(path, network)
		if avg < 0.001 {
			break
		}
	}
}
