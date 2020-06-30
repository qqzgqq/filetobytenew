package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"log"
	"encoding/base64"
)

// GzipDecode aksdjalsd
func GzipDecode(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		var out []byte
		return out, err
	}
	defer reader.Close()

	return ioutil.ReadAll(reader)
}

// Read3 read file to byte
func Read3(s string) []byte {
	f, err := os.Open(s)
	if err != nil {
		fmt.Println("read file fail", err)
		os.Exit(0)
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		os.Exit(0)

	}

	return fd
}

// Write3 write file
func Write3(ttt []byte, ss string) {
	fileName := ss
	f, err3 := os.Create(fileName) //创建文件
	if err3 != nil {
		fmt.Println("create file fail")
	}
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n4, err3 := w.WriteString(string(ttt))
	fmt.Printf("写入 %d 个字节n", n4)
	w.Flush()
	f.Close()
}

// Stringtobyte String to byte
func Stringtobyte(ss string) []byte {
	// checkout [] exist
	var ii string
	var test []byte
	kuohao := strings.ContainsAny(ss, "[&]")
	if kuohao {
		ss = strings.Replace(ss, "[", "", -1)
		ss = strings.Replace(ss, "]", "", -1)
		ss = strings.Replace(ss, "\r", " ", -1)
		ss = strings.Replace(ss, "\n", "", -1)
		i := strings.Count(ss, " ")
		cc := strings.SplitAfterN(ss, " ", i+1)
		for _, ii = range cc {
			ii = strings.Replace(ii, " ", "", -1)
			Shuzui, _ := strconv.Atoi(ii)
			test = append(test, byte(Shuzui))

		}
	} else {
		fmt.Println("can't find kuohao")
		os.Exit(2)
	}
	return test
}

// OsIoutil read file to byte
func OsIoutil(name string) []byte {
	var contents []byte
	if fileObj, err := os.Open(name); err == nil {
		//if fileObj,err := os.OpenFile(name,os.O_RDONLY,0644); err == nil {
		defer fileObj.Close()
		if contents, err = ioutil.ReadAll(fileObj); err != nil {

			panic(err)
		}

	}
	return contents
}

func main() {
	flag.Parse()
	var file = flag.Arg(0)
	var file2 = flag.Arg(1)
	// test := Read3(file)
	// fmt.Println(test)
	var Result []byte

	Result = OsIoutil(file)
	// fmt.Println(Result)

	DecodeBytes, err := base64.StdEncoding.DecodeString(string(Result))
	if err != nil {
		log.Fatalln(err)
	}
	JY, _ := GzipDecode(DecodeBytes)
	// bb := Stringtobyte(Result)
	// fmt.Println(bb)

	Write3(JY, file2)

}
