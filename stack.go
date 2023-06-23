package main

import "errors"

type Stack struct {
	MaxSize int
	Top     int
	arr     [10]interface{}
}

func initStack() (stack *Stack) {
	stack = &Stack{
		MaxSize: 10,
		Top:     -1,
		arr:     [10]interface{}{},
	}
	return stack
}

//入栈
func (s *Stack) push(v interface{}) error {
	if s.MaxSize-1 == s.Top {
		return errors.New("栈已满，无法插入数据")
	}
	s.Top++
	s.arr[s.Top] = v
	return nil
}

//出栈
func (s *Stack) pop() (interface{}, error) {
	if s.Top == -1 {
		return 0, errors.New("error:栈为空")
	}
	v := s.arr[s.Top]
	s.Top--
	return v, nil
}
