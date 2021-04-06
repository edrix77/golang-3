package main

import (
	"fmt"
)

func main() {
	var i, n int
	var j int
	fmt.Print("Jumlah = ")
	fmt.Scan(&n)

	fmt.Print("fibonacci : ")
	for i = 1; i <= n; i++ {
		fmt.Print(fibonaci(j), ",")
		j++
	}
	fmt.Println()
}

func fibonaci(j int) int {
	if j < 2 {
		return j
	}
	return fibonaci(j-1) + fibonaci(j-2)
}
