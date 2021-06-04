package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
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
	c := make([]int, n)
	s := make([]int, m_size)
	s0 := make([]int, m_size)
	//count := 0
	out_fl, e2 := os.Create("output.txt")
	if e2 != nil {
		fmt.Println(os.Stderr, e2)
		os.Exit(1)
	}
	defer out_fl.Close()
	writer := bufio.NewWriter(out_fl)
	for i := 0; i < code_size; i++ {
		tmp = 0
		i_n = i
		row = 0
		s = make([]int, m_size)
		for i_n != 0 { //符号語Cを作る
			tmp = i_n % 2
			i_n /= 2
			c[row] = tmp
			row++
		}
		//fmt.Printf("%d | code is %d\n", i, c)
		for i := 0; i < m_size; i++ { //シンドロームsをHc^t=sより作成
			for j := 0; j < n; j++ {
				s[i] += H[i][j] * c[j]
			}
			s[i] %= 2
			if s[i] != 0 {
				continue
			}
		}
		//fmt.Printf("%d | syndrome is %d\n", i, s)
		if reflect.DeepEqual(s, s0) {
			//fmt.Printf("%d code is %d\n", count, c)
			//count++
			if _, e3 := fmt.Fprintln(writer, c); e3 != nil {
				fmt.Println(os.Stderr, e3)
				os.Exit(1)
			}
		}
	} //*/
	writer.Flush()
}
