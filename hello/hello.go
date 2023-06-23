package hello

import "fmt"

func init() {
	fmt.Println("imp-init() come here")
}

func Print() {
	fmt.Println("Hello!")
}

func main() {
	fmt.Printf("%d\n", 1100)
}
