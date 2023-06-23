package main

import "fmt"

func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}

func fibonaci2(s map[int]int, i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	value, ok := s[i]
	if ok {
		return value
	}
	s[i] = fibonaci2(s, i-1) + fibonaci2(s, i-2)
	return s[i]
}

func fibonaci3(n int) int {
	l1, l2 := 1, 1
	var temp int
	for i := 3; i <= n; i++ {
		temp = l1 + l2
		l1 = l2
		l2 = temp
	}
	return l2
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		m := make(map[int]int)
		fmt.Printf("%d\n", fibonaci2(m, i))
	}
}
