package algorithm

import "math/rand"

var arr = []int{3, 6, 1, 7, 8, 9, 7, 2}

// 冒泡 稳定
func Sort1() []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			temp := arr[j+1]
			if arr[j] > arr[j+1] {
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
	for i := range arr {
		print(arr[i])
	}
	return arr
}

// 选择 不稳定
func Sort2() []int {

	return nil
}

// 快速 不稳定
func Sort4(arr []int, lo int, hi int) []int {
	if lo >= hi {
		return arr
	}
	p := partition(arr, lo, hi)
	Sort4(arr, lo, p-1)
	Sort4(arr, p+1, hi)
	return arr
}

func partition(arr []int, lo int, hi int) int {
	swap(arr, lo+rand.Intn(hi-lo), hi)
	var i, j int
	for i, j = lo, lo; j < hi; j++ {
		if arr[j] <= arr[hi] {
			swap(arr, i, j)
			i++
		}
	}
	swap(arr, i, j)
	return i
}

func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

// 归并 稳定
func sort5(arr []int) []int {
	len := len(arr)
	if len == 1 {
		return arr
	}
	m := len / 2
	left := sort5(arr[:m])
	right := sort5(arr[m:])
	return merge(left, right)
}

func merge(l []int, r []int) []int {
	lenL := len(l)
	lenR := len(r)
	res := make([]int, 0)
	lIndex, rIndex := 0, 0
	for lIndex < lenL && rIndex < lenR {
		if l[lIndex] > r[rIndex] {
			res = append(res, r[rIndex])
			rIndex++
		} else {
			res = append(res, l[lIndex])
			lIndex++
		}
	}
	if lIndex < lenL { //左边的还有剩余元素
		res = append(res, l[lIndex:]...)
	}
	if rIndex < lenR {
		res = append(res, r[rIndex:]...)
	}
	return res
}
