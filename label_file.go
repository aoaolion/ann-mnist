package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
)

type LabelFile struct {
	Msb   int
	Num   int
	Label []byte
}

func NewLabel(path string) (*LabelFile, error) {
	labelFile := &LabelFile{}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	labelFile.Msb = int(buf[0])<<24 | int(buf[1])<<16 | int(buf[2])<<8 | int(buf[3])
	if labelFile.Msb != 2049 {
		return nil, errors.New("format error")
	}

	labelFile.Num = int(buf[4])<<24 | int(buf[5])<<16 | int(buf[6])<<8 | int(buf[7])
	labelFile.Label = bytes.NewBuffer(buf[8:]).Bytes()
	return labelFile, nil
}

func (lf *LabelFile) String() string {
	return fmt.Sprintf("LableFile=>Msb:%d, Num:%d", lf.Msb, lf.Num)
}
