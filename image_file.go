package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
)

type ImageFile struct {
	Msb   int
	Num   int
	Row   int
	Col   int
	Pixel []byte
}

func NewImageFile(path string) (*ImageFile, error) {
	imageFile := &ImageFile{}
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	imageFile.Msb = int(buf[0])<<24 | int(buf[1])<<16 | int(buf[2])<<8 | int(buf[3])
	if imageFile.Msb != 2051 {
		return nil, errors.New("format error")
	}

	imageFile.Num = int(buf[4])<<24 | int(buf[5])<<16 | int(buf[6])<<8 | int(buf[7])
	imageFile.Row = int(buf[8])<<24 | int(buf[9])<<16 | int(buf[10])<<8 | int(buf[11])
	imageFile.Col = int(buf[12])<<24 | int(buf[13])<<16 | int(buf[14])<<8 | int(buf[15])

	imageFile.Pixel = bytes.NewBuffer(buf[16:]).Bytes()
	return imageFile, nil
}

func (imgf *ImageFile) GetImage(idx int) []byte {
	offset := imgf.Row * imgf.Col * idx
	offset2 := imgf.Row * imgf.Col * (idx + 1)
	buf := bytes.NewBuffer(imgf.Pixel[offset:offset2]).Bytes()
	return buf
}

func (imgf *ImageFile) SaveImage(path string, idx int) error {
	buf := imgf.GetImage(idx)
	fp, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fp.Close()
	img := image.NewRGBA(image.Rect(0, 0, imgf.Col, imgf.Row))
	for x := 0; x < imgf.Row; x++ {
		for y := 0; y < imgf.Col; y++ {
			v := 255 - buf[y*imgf.Col+x]
			img.Set(x, y, color.RGBA{v, v, v, 255})
		}
	}
	err = png.Encode(fp, img)
	if err != nil {
		return err
	}

	return nil
}

func (imgf *ImageFile) String() string {
	return fmt.Sprintf("ImageFile=>Msb:%d, Num:%d, Row:%d, Col:%d",
		imgf.Msb, imgf.Num, imgf.Row, imgf.Col)
}
