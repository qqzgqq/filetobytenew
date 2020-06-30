package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"os/exec"
	"os"
	"flag"
)

func main(){
	// var FILEDIR1 = "D:\\go_workplace\\releases\\filetobyte3.0\\server\\pictures\\"
	flag.Parse()
	var PYFILE1 = flag.Arg(0)
	var FILEDIR1 = flag.Arg(1)
	var SERVERFILE1 = flag.Arg(2)
	var PYCMDSTRING string
	FileSorT := []int{}
	var FNNUM int
FileS,_:=filepath.Glob(FILEDIR1+"*")
for _, i := range FileS {
	FNNUM, _ = strconv.Atoi(strings.Replace(filepath.Base(i), ".png", "", -1))
	FileSorT = append(FileSorT, FNNUM)
}
sort.Ints(FileSorT)
for i := range FileSorT{
	PYCMDSTRING="python "+PYFILE1+" "+FILEDIR1+strconv.Itoa(i)+".png"+">"+SERVERFILE1+strconv.Itoa(i)+".png.log"
	c := exec.Command("cmd", "/C", PYCMDSTRING)
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}  
	fmt.Println(strconv.Itoa(i)+".png","decode success")
}

}