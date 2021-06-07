package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func replace() {
	fp, e1 := os.Open("output.txt")
	if e1 != nil {
		panic(e1)
	}
	defer fp.Close()
	b, e11 := ioutil.ReadAll(fp)
	if e11 != nil {
		panic(e1)
	}
	str := string(b)
	// 普通に表示
	fmt.Println(str)

	//置換作業
	str = strings.Replace(str, "[", "", -1)
	str = strings.Replace(str, "]", "", -1)
	str = strings.Replace(str, " ", "", -1)
	fmt.Printf("replaced :\n%s\n", str)
	out_fl, e2 := os.Create("output2.txt")
	if e2 != nil {
		panic(e2)
	}
	defer out_fl.Close()
	writer := bufio.NewWriter(out_fl)
	if _, e3 := fmt.Fprintln(writer, str); e3 != nil {
		panic(e3)
	}
	writer.Flush()
}
