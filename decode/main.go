package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func main() {
	flag.Parse()
	var PNGfile = flag.Arg(0)

	// open and decode image file
	file, _ := os.Open(PNGfile)
	img, _, _ := image.Decode(file)

	// prepare BinaryBitmap
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, _ := qrReader.Decode(bmp, nil)

	fmt.Println(result)

}
