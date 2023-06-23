package main

import "fmt"

type Website struct {
	Name string
}

var site = Website{Name: "studygolang"}

type X string

func (x X) String() string {
	return fmt.Sprintf("<%s>", string(x))
}

var xx X = "dddddd"

func main() {
	s := xx.String()
	fmt.Println(s)

	fmt.Printf("%v\n", site)
	fmt.Printf("%+v", site)
	fmt.Println()
	fmt.Printf("%#v", site)
	fmt.Println()
	fmt.Printf("%T", site)
	fmt.Println()
	fmt.Printf("%%")

}
