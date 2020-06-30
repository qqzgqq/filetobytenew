package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	var FILEDIR = flag.Arg(0)
	// FILEDIR := "D:\\go_workplace\\releases\\filetobyte3.0\\server\\rar\\"
	var FNN = []int{}
	var GHNM, GQNM string
	var RARDIRSTNAME = "1.part01.rar"
	FileS, error := filepath.Glob(FILEDIR + "*")
	if error != nil {
		log.Fatal(error)
		fmt.Println("获取rar文件夹文件列表失败")
		os.Exit(0)
	}
	FileNum := len(FileS) - 1
	fmt.Println(FileNum)
	for i := 0; i <= FileNum; i++ {
		aa, _ := strconv.Atoi(strings.Replace(strings.Replace(filepath.Base(FileS[i]), "1.part", "", -1), ".rar", "", -1))
		FNN = append(FNN, aa)
	}
	sort.Ints(FNN)
	ZNUM := strings.Count(strconv.Itoa(FNN[FileNum]), "") - 1
	switch ZNUM {
	case 1:
		RARDIRSTNAME = FILEDIR + RARDIRSTNAME
	case 2:
		for i := 1; i < 10; i++ {
			GQNM = FILEDIR + "1.part0" + strconv.Itoa(i) + ".rar"
			GHNM = FILEDIR + "1.part00" + strconv.Itoa(i) + ".rar"
			error := os.Rename(GQNM, GHNM)
			if error != nil {
				fmt.Println(error)
				fmt.Println(GQNM, "rename bad")
				os.Exit(0)
			}
			fmt.Println(GQNM, "rename success")
		}
		RARDIRSTNAME = FILEDIR + "1.part001.rar"
	case 3:
		for i := 1; i < 100; i++ {
			if i <= 9 {
				GQNM = FILEDIR + "1.part0" + strconv.Itoa(i) + ".rar"
				GHNM = FILEDIR + "1.part000" + strconv.Itoa(i) + ".rar"
				error := os.Rename(GQNM, GHNM)
				if error != nil {
					fmt.Println(error)
					fmt.Println(GQNM, "rename bad")
					os.Exit(0)
				}
				fmt.Println(GQNM, "rename success")
			} else {
				GQNM = FILEDIR + "1.part0" + strconv.Itoa(i) + ".rar"
				GHNM = FILEDIR + "1.part00" + strconv.Itoa(i) + ".rar"
				error := os.Rename(GQNM, GHNM)
				if error != nil {
					fmt.Println(error)
					fmt.Println(GQNM, "rename bad")
					os.Exit(0)
				}
				fmt.Println(GQNM, "rename success")
			}

		}
		RARDIRSTNAME = FILEDIR + "1.part0001.rar"
	}
	c := exec.Command("cmd", "/C", "start", RARDIRSTNAME)
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
