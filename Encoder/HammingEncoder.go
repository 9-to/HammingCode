package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	flag.Parse() //input n
	m_size, _ := strconv.Atoi(flag.Arg(0))
	n := int(math.Pow(2, float64(m_size))) - 1 //127
	k := n - m_size                            //120
	fmt.Printf("I'll make (%d,%d) Hamming code.\n", n, k)
	H := make([][]int, n-k) //7*127
	for i := range H {
		H[i] = make([]int, n)
	}
	tmp := 0
	row := 0
	i_n := 0
	for i := 1; i <= n; i++ {
		tmp = 0
		row = 0
		i_n = i
		column := i - 1
		//fmt.Println(i_n, tmp, row)
		for i_n != 0 {
			tmp = i_n % 2
			i_n /= 2
			H[row][column] = tmp
			row++
		}
	}
	fmt.Println("Parity check matrix H =")
	for i := 0; i < n-k; i++ {
		fmt.Println(H[i])
	}
	fmt.Println("----------------------------")
	code_size := int(math.Pow(2, float64(n)))
	c := make([]string, n)   //符号語
	s := make([]int, m_size) //シンドローム
	C := 0
	s0 := make([]int, m_size) //比較用 [0,0,...]が収納されている
	counter := 0
	out_fl, e2 := os.Create("output.txt")
	if e2 != nil {
		panic(e2)
	}
	defer out_fl.Close()
	writer := bufio.NewWriter(out_fl)

	header := "label "
	for i := 0; i < n; i++ {
		header = header + "c_" + strconv.Itoa(i) + " "
	}
	header2 := strings.TrimSpace(header)
	if _, err := fmt.Fprintln(writer, header2); err != nil {
		panic(err)
	}

	for i := 0; i < code_size; i++ {
		tmp = 0
		i_n = i
		row = 0
		s = make([]int, m_size)
		for i := 0; i < n; i++ {
			c[i] = "0 "
		}
		for i_n != 0 { //符号語Cを作る
			tmp = i_n % 2
			i_n /= 2
			c[row] = strconv.Itoa(tmp) + " "
			row++
		}
		//fmt.Printf("%d | code is %s\n", i, c)
		for i := 0; i < m_size; i++ { //シンドロームsをHc^t=sより作成
			for j := 0; j < n; j++ {
				str3 := strings.TrimSpace(c[j])
				C, _ = strconv.Atoi(str3)
				s[i] += H[i][j] * C
			}
			s[i] %= 2
			/*if s[i] != 0 {
				continue
			}*/
		}
		//fmt.Printf("%d | syndrome is %d\n", i, s)
		str := strings.Join(c, "")
		if reflect.DeepEqual(s, s0) { //HammingCodeのラベルは0、それ以外は1
			str = "0 " + str
			counter++
		} else {
			str = "1 " + str
		}
		str2 := strings.TrimSpace(str)
		if _, e3 := fmt.Fprintln(writer, str2); e3 != nil {
			panic(e3)
		}
	} //*/
	fmt.Printf("code_size is %d\n", counter)
	writer.Flush()
}
