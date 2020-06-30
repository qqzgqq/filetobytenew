package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// Read3 read file to string
func Read3(s string) string {
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

	return string(fd[:])
}
func main() {
	flag.Parse()
	var FILEDIR = flag.Arg(0)
	// var FILEDIR = "D:\\go_workplace\\releases\\filetobyte3.0\\server\\decode\\"
	var File1 string
	FILEMAP := make(map[int]string)
	var ZSZ = []int{}
	var ZSZNEW = []int{}
	var FNNUM int
	FileSorT := []int{}
	FileS, error := filepath.Glob(FILEDIR + "*")
	if error != nil {
		log.Fatal(error)
		fmt.Println("获取decode文件夹文件列表失败")
		os.Exit(0)
	}
	for _, i := range FileS {
		FNNUM, _ = strconv.Atoi(strings.Replace(filepath.Base(i), ".png.log", "", -1))
		FileSorT = append(FileSorT, FNNUM)
	}
	sort.Ints(FileSorT)
	fmt.Println("file name to sort shuzu ", FileSorT)
	FileSnum := len(FileS)
	fmt.Println("file name number is ", FileSnum)
	for i := 0; i < FileSnum; i++ {
		File1 = Read3(FILEDIR + strconv.Itoa(FileSorT[i]) + ".png.log")
		FILEMAP[FileSorT[i]] = File1
	}

	for i, k := range FILEMAP {
		if strings.Contains(k, "Can not recognize.") {
			os.RemoveAll(FILEDIR + strconv.Itoa(i) + ".png.log")
			delete(FILEMAP, i)
			fmt.Println("clear error eweima pic ", i)
		}
	}
	for i := range FILEMAP {
		ZSZ = append(ZSZ, i)
	}
	sort.Ints(ZSZ)
	fmt.Println("will srot zhuzu", ZSZ)
	ZSZ2 := len(ZSZ) - 1
	for i := 0; i < ZSZ2; i++ {

		if FILEMAP[ZSZ[i]] == FILEMAP[ZSZ[i+1]] {

			fmt.Println("start duibi shuzu ", ZSZ[i], ZSZ[i+1])
			os.RemoveAll(FILEDIR + strconv.Itoa(ZSZ[i]) + ".png.log")
			delete(FILEMAP, ZSZ[i])
			fmt.Println("delete chongfu eweima pic ", ZSZ[i])
		}
	}

	for i := range FILEMAP {

		ZSZNEW = append(ZSZNEW, i)
	}
	fmt.Println("shengyu shuzu ", ZSZNEW)
	sort.Ints(ZSZNEW)
	fmt.Println("sort new shuzu ", ZSZNEW)
	fmt.Println("文件筛选成功")
}
