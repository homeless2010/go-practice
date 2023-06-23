package algorithm

import "testing"

func TestSort2(t *testing.T) {
	arr := Sort1()
	for i := range arr {
		t.Log(arr[i])
	}
}
