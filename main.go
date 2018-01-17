package main

import (
	"flag"
	"fmt"

	neural "github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/learn"
	"github.com/NOX73/go-neural/persist"
	"github.com/aoaolion/ann-mnist/common/logger"
	log "github.com/cihub/seelog"
)

var (
	mode = flag.String("mode", "train", "working mode train or test")
)

const (
	trainLabel = "data/train-labels-idx1-ubyte"
	trainImage = "data/train-images-idx3-ubyte"
	testLabel  = "data/t10k-labels-idx1-ubyte"
	testImage  = "data/t10k-images-idx3-ubyte"
)

func train() {
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
	lastAvg := 1.0

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
		}
		avg = avg / float64(maxTrainSet)
		log.Infof("round:%d, e:%f", round, avg)
		if avg < 0.01 {
			break
		}
		if lastAvg < avg {
			log.Info("train too much")
			break
		}
		lastAvg = avg
		persist.ToFile("data/network.json", network)
	}
}

func test() {
	labelFile, imageFile, err := loadDataSet(testLabel, testImage)
	if err != nil {
		log.Error(err)
		return
	}
	network := persist.FromFile("data/network.json")
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
		log.Infof("[%d] %d -> %d, possible:%f", i, labelFile.Label[i], resultNum, resultRate)
		if int(labelFile.Label[i]) == resultNum {
			accurateNum++
		}
	}
	log.Infof("Accurate:%d, Rate:%f", accurateNum, float64(accurateNum)/float64(maxTestSet))
}

func export() {
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

func main() {
	defer log.Flush()
	flag.Parse()

	logger.InitLogger("conf/logger.xml", true)
	log.Info(*mode, " mode is start")

	if *mode == "train" {
		train()
	} else if *mode == "test" {
		test()
	} else if *mode == "export" {
		export()
	}
}
