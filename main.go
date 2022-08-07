package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	var k float64
	if len(arguments) == 1 {
		fmt.Println("Bir ve ya birden cox arqument girin zehhmet olmasa")
		os.Exit(1)
	}
	for i := 1; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		k+=n
	}
	fmt.Println(k)

}
