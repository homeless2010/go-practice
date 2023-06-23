package main

import (
	"fmt"
	"reflect"
)

var x float64 = 3.4

type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
	typeOf := reflect.TypeOf(x)
	valueOf := reflect.ValueOf(x)
	fmt.Println(typeOf)
	fmt.Println(typeOf.Kind())
	fmt.Println(valueOf)
	fmt.Println(valueOf.Type())
	fmt.Println("////////////////////////")
	type MyInt int
	var m MyInt = 5
	value := reflect.ValueOf(m)
	fmt.Println(value.Kind())
	fmt.Println(value.Interface())
	//反射设值
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("settability of v:", v.CanSet())
	v = reflect.ValueOf(&x)
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())
	v = v.Elem()
	fmt.Println("The Elem of v is:", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.1415)
	fmt.Println(v.Interface())
	fmt.Println(v)
	//reflect struct
	fmt.Println("/////////////////////////////////")
	fmt.Println("////////reflect struct///////////")
	fmt.Println("/////////////////////////////////")
	val := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)
	fmt.Println(val)
	fmt.Println(typ)
	fmt.Println(val.Kind())
	for i := 0; i < val.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, val.Field(i))
		//结构中只有被到处字段才可以设置
		val.Field(i).SetString("C#")
	}
}
