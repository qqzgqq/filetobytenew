package main

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"encoding/base64"
	"flag"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"unsafe"

	"github.com/makiuchi-d/gozxing/qrcode"
	// qrcode "github.com/skip2/go-qrcode"
)

// String1 just make byte to string
func String1(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Compress lasjdlfadlj
func Compress(s string) string {
	// //使用GBK字符集encode
	// gbk, err := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(s))
	// if err != nil {

	// return ""
	// }

	// //转为ISO8859_1，也就是latin1字符集
	// latin1, err := charmap.ISO8859_1.NewDecoder().Bytes(gbk)
	// if err != nil {
	// return ""
	// }

	//使用gzip压缩
	var buf bytes.Buffer
	// zw := gzip.NewWriter(&buf)
	zw, _ := gzip.NewWriterLevel(&buf, flate.BestCompression)
	_, err1 := zw.Write([]byte(s))
	if err1 != nil {
		panic(err1)
	}

	if err := zw.Close(); err != nil {
		panic(err)
	}
	//使用base64编码
	// encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	encoded := base64.StdEncoding.EncodeToString([]byte(s))

	// fmt.Println(encoded)
	return encoded
}

// GzipEncode asdfhka
func GzipEncode(in []byte) ([]byte, error) {
	var (
		buffer bytes.Buffer
		out    []byte
		err    error
	)
	writer, _ := gzip.NewWriterLevel(&buffer, flate.BestCompression)
	_, err = writer.Write(in)
	if err != nil {
		writer.Close()
		return out, err
	}
	err = writer.Close()
	if err != nil {
		return out, err
	}

	return buffer.Bytes(), nil
}

func main() {
	flag.Parse()
	var DIr = flag.Arg(0)

	Bytes, err := ioutil.ReadFile(DIr)
	if err != nil {
		fmt.Println("read ", DIr, " bad")
		log.Fatal(err)
	}
	fmt.Println(DIr, "========read====success:", len(Bytes))

	//make gzip encode
	YS, _ := GzipEncode(Bytes)
	fmt.Println(DIr, "========yasuo_byte===success:", len(YS))
	//make base64 encode
	YS2 := Compress(string(YS))
	fmt.Println(DIr, "===base64_encode==sucess:", len(YS2))
	qrWiter := qrcode.NewQRCodeWriter()
	qrRRR, error := qrWiter.EncodeWithoutHint(YS2, 11, 1024, 1024)
	if error != nil {
		log.Fatal(error)
	}

	FF, err := os.Create(DIr + ".png")
	defer FF.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	_ = png.Encode(FF, qrRRR)
}
