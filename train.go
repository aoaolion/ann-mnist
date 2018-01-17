package main

import (
	neural "github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/learn"
	"github.com/NOX73/go-neural/persist"
	log "github.com/cihub/seelog"
)

func Train() {
	labelFile, imageFile, err := loadDataSet(trainLabel, trainImage)
	if err != nil {
		log.Error(err)
		return
	}
	// 24*24 = 784
	network := neural.NewNetwork(784, []int{784, 30, 10})
	network.RandomizeSynapses()

	maxTrainRound := 100
	maxTrainSet := imageFile.Num

	for round := 0; round < maxTrainRound; round++ {
		avg := 0.0
		for i := 0; i < maxTrainSet; i++ {
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
			//			if i%100 == 0 {
			//				log.Infof("round:%d, training:%d, estimate:%f", round, i, estimate)
			//			}
		}
		avg = avg / float64(maxTrainSet)
		log.Infof("round:%d, e:%f", round, avg)
		if avg < 0.01 {
			break
		}
		persist.ToFile("data/network.json", network)
	}
}
