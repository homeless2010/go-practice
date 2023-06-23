package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)

	fmt.Printf("len: %v\n", l.Len())
	fmt.Printf("first: %#v\n", l.Front())
	fmt.Printf("second: %#v\n", l.Front().Next())
}
