package main

import (
	"sync"
)

type ArrayStack struct {
	array []string
	size  int
	lock  sync.Mutex
}

func main() {
}

//入栈
func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.array = append(stack.array, v)
	stack.size = stack.size + 1
}

//出栈

func (stack *ArrayStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.size == 0 {
		panic("empty")
	}
	v := stack.array[stack.size-1]
	//stack.array = stack.array[0 : stack.size-1]
	newArray := make([]string, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArray[i] = stack.array[i]
	}
	stack.array = newArray
	stack.size = stack.size - 1
	return v
}
