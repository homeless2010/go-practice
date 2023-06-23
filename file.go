package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("test1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("hello world golang")
	s, err2 := file.WriteAt([]byte("Go语言中文网"), 10)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(s)

	bytes.NewReader([]byte("Go语言中文网")).WriteTo(os.Stdout)
	fmt.Println()

	reader := strings.NewReader("Go语言中文网")
	reader.Seek(-6, io.SeekEnd)
	r, _, _ := reader.ReadRune()
	fmt.Printf("%c\n", r)
	log.Writer()

	fmt.Println(len("谷歌中国"))
	fmt.Println(len("谷"))
	fmt.Println(len("歌"))
	fmt.Println(len("中"))
	fmt.Println(len("国"))

	fmt.Println(strings.Count("fivevev", "vev"))

	fmt.Printf("Fields are: %q\n", strings.Fields(" foo bar baz  "))
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
}
