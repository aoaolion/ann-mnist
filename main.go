package main

import (
	"flag"

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

func main() {
	defer log.Flush()
	flag.Parse()

	logger.InitLogger("conf/logger.xml", true)
	log.Info(*mode, " mode is start")

	if *mode == "train" {
		Train()
	} else if *mode == "test" {
		Test()
	} else if *mode == "export" {
		Export()
	}
}
