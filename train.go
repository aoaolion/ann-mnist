package main

import (
	"fmt"

	neural "github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/learn"
	"github.com/NOX73/go-neural/persist"
	log "github.com/cihub/seelog"
)

func Train(maxIteration, maxTrainSize int) {
	labelFile, imageFile, err := loadDataSet(trainLabel, trainImage)
	if err != nil {
		log.Error(err)
		return
	}
	// 24*24 = 784
	network := neural.NewNetwork(784, []int{300, 100, 10})
	network.RandomizeSynapses()

	log.Infof("maxIteration: %d, trainSet: %d", maxIteration, maxTrainSize)

	if maxTrainSize > imageFile.Num {
		maxTrainSize = imageFile.Num
	}

	for iteration := 0; iteration < maxIteration; iteration++ {
		avg := 0.0
		for i := 0; i < maxTrainSize; i++ {
			//s := time.Now()
			in := make([]float64, 0)
			buf := imageFile.GetImage(i)
			for _, v := range buf {
				in = append(in, float64(v))
			}

			ideal := make([]float64, 0)
			for j := 0; j < 10; j++ {
				if j == int(labelFile.Label[i]) {
					ideal = append(ideal, 1)
				} else {
					ideal = append(ideal, 0)
				}
			}

			learn.Learn(network, in, ideal, 0.2)
			estimate := learn.Evaluation(network, in, ideal)
			avg += estimate
			//log.Info("4. ", time.Since(s))
			//			if i%100 == 0 {
			//				log.Infof("iteration:%d, training:%d, estimate:%f", iteration, i, estimate)
			//			}
		}
		avg = avg / float64(maxTrainSize)
		log.Infof("iteration:%d, e:%f", iteration, avg)
		if avg < 0.01 {
			break
		}
		path := fmt.Sprintf("data/network_%d_%f.json", iteration, avg)
		persist.ToFile(path, network)
		go Test(path)
	}
}
