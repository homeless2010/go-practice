package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 1
	var num2 int32
	fmt.Printf("%d\n", num)
	fmt.Printf("%d", num2)
	atoi, err := strconv.Atoi("wdgdfgfgfd")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ffffff", atoi)
}
