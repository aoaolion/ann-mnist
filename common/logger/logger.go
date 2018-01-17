package logger

import (
	"io/ioutil"
	"os"
	"strings"

	log "github.com/cihub/seelog"
)

func InitLogger(confPath string, verose bool) {
	fp, err := os.Open(confPath)
	if err != nil {
		panic("conf file load error")
	}
	buf, err := ioutil.ReadAll(fp)
	if err != nil {
		panic("conf file load error")
	}
	confStr := string(buf)
	if verose {
		confStr = strings.Replace(confStr, "minlevel=\"info\"", "minlevel=\"debug\"", -1)
	}
	logger, err := log.LoggerFromConfigAsString(confStr)
	if err != nil {
		log.Critical("err parsing config log file", err)
		return
	}
	log.ReplaceLogger(logger)
	if verose {
		log.Info("logger init in verose mode")
		confStr = strings.Replace(confStr, "minlevel=\"info\"", "minlevel=\"debug\"", -1)
	} else {
		log.Info("logger init")
	}
}
