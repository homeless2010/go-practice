package main

import "fmt"

func main() {
	//var mapInit := map[string]string{"xiaoli": "湖南", "xiaoliu": "天津"}
	var mapTemp map[string]string
	mapTemp = make(map[string]string, 10)
	mapTemp["xiaoming"] = "北京"
	mapTemp["xiaowang"] = "河北"
	v1, ok := mapTemp["xiaoming"]
	fmt.Println(ok, v1)
	for k, v := range mapTemp {
		fmt.Println(k, v)
	}
	delete(mapTemp, "xiaoming")
	l := len(mapTemp)
	fmt.Println(l)
}
