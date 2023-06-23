package main

import (
	"fmt"
	"strings"
)

/**
闭包
*/

//工厂函数
func MackAddSuffix(suffix string) func(string2 string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

//
func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

func main() {
	addBmp := MackAddSuffix(".bmp")
	addJpeg := MackAddSuffix(".jpeg")

	fmt.Println("addBmp", addBmp("file"))
	fmt.Println("addJpeg", addJpeg("file"))

	adder := Adder()
	fmt.Println("adder1", adder(1))
	fmt.Println("adder20", adder(20))
	fmt.Println("adder300", adder(300))
}
