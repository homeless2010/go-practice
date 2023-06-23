package main

import "fmt"

func printAnything(thing interface{}) {
	fmt.Println(thing)
}

type Singer string

func (w Singer) Sing() {
	fmt.Println(string(w))
}

func MakeSound(thing interface{}) {
	//使用类型断言，将空接口转换成具体类型
	singer, ok := thing.(Singer)
	fmt.Printf("%v\n", ok)
	if ok {
		singer.Sing()
	}
}

//类型断言

func main() {
	printAnything(1)
	printAnything(true)
	printAnything("dddddddddd")
	MakeSound(Singer("Go Away"))
}
