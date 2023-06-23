package main

import (
	"fmt"
)

type A struct {
	name string
	age  int8
}

func main() {
	//http.HandleFunc("/", dddd.Redirect)
	//http.HandleFunc("/add", dddd.Add)
	//http.ListenAndServe(":8080", nil)
	ms := A{"ssss", 12}
	fmt.Printf("%v\n", ms)
	fmt.Printf("%+v\n", ms)
	fmt.Printf("%#v\n", ms)
	fmt.Printf("%T\n", ms)
	fmt.Printf("%%\n")
	fmt.Println(ms)

	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	if _, ok := scoreMap["张三"]; ok {
		fmt.Println("dddddddddddddddddd")
	}

	for s := range scoreMap {
		fmt.Println("ddd " + s)
	}

	for s, i := range scoreMap {
		fmt.Println(s, " ", i)
	}

	var mapSlice = make([]map[string]string, 3)
	for i, m := range mapSlice {
		fmt.Printf("index:%d value:%v\n", i, m)
	}
	fmt.Println("after init")
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	for i, m := range mapSlice {
		fmt.Printf("index:%d value:%v\n", i, m)
	}
	fmt.Println(1 << 3)
}
