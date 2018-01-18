package main

import (
	"flag"
	"net/http"
	_ "net/http/pprof"

	"github.com/aoaolion/ann-mnist/common/logger"
	log "github.com/cihub/seelog"
)

var (
	mode         = flag.String("mode", "train", "working mode train or test")
	maxIteration = flag.Int("i", 100, "max iteration")
	maxTrainSize = flag.Int("t", 100, "max train set size")
	network      = flag.String("network", "data/network.json", "saved neural network")
)

const (
	trainLabel = "data/train-labels-idx1-ubyte"
	trainImage = "data/train-images-idx3-ubyte"
	testLabel  = "data/t10k-labels-idx1-ubyte"
	testImage  = "data/t10k-images-idx3-ubyte"
)

func main() {
	defer log.Flush()
	flag.Parse()

	logger.InitLogger("conf/logger.xml", true)
	log.Info(*mode, " mode is start")

	//远程获取pprof数据
	go func() {
		err := http.ListenAndServe("localhost:8888", nil)
		if err != nil {
			log.Error(err)
			return
		}
	}()

	if *mode == "train" {
		Train(*maxIteration, *maxTrainSize)
	} else if *mode == "test" {
		Test(*network)
	} else if *mode == "export" {
		Export()
	}
}
