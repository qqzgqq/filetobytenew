package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"image/png"
	"os"
	"strconv"
)

func main() {

	flag.Parse()
	var GIFFILE = flag.Arg(0)
	var PNGFILE = flag.Arg(1)
	PNGFILEDIR := PNGFILE
	f, err := os.Open(GIFFILE)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	inGif, err := gif.DecodeAll(f)
	if err != nil {
		panic(err)
	}
	config, _ := gif.DecodeConfig(f)
	//  获取GIF的高宽
	rect := image.Rect(0, 0, config.Width, config.Height)
	if rect.Min == rect.Max {
		var max image.Point
		for _, frame := range inGif.Image {
			maxF := frame.Bounds().Max
			if max.X < maxF.X {
				max.X = maxF.X
			}
			if max.Y < maxF.Y {
				max.Y = maxF.Y
			}
		}
		rect.Max = max
	}
	// format := fmt.Sprintf("%%dd", len(string(len(inGif.Image)))+1)
	for i, srcimg := range inGif.Image {
		img := image.NewRGBA(rect)
		subfn := strconv.Itoa(i) + ".png"
		PNGFILE = PNGFILEDIR + subfn
		f, err := os.Create(PNGFILE)
		if err != nil {
			panic(err)
		}
		PNGFILE = ""
		defer f.Close()
		draw.Draw(img, srcimg.Bounds(), srcimg, srcimg.Rect.Min, draw.Src)
		fmt.Printf("\r%s", subfn)
		png.Encode(f, img)
	}
	fmt.Println("\r", PNGFILE)

}
