package main

import "fmt"

//工厂模式
type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &hellAPI{}
	}
	return nil
}

type hiAPI struct {
}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi,%s", name)
}

type hellAPI struct {
}

func (*hellAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
