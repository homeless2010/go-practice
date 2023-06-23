package main

import (
	"fmt"
	"math/rand"
)

//冒泡 稳定
func sort1(arr []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	for i, v := range arr {
		fmt.Println(i, "----", v)
	}
}

//选择 不稳定
func sort2(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		temp := arr[i]
		var index = i
		for j := i + 1; j < l; j++ {
			if arr[index] > arr[j] {
				index = j
			}
		}
		if i == index {

		} else {
			arr[i] = arr[index]
			arr[index] = temp
		}
	}
	for _, v := range arr {
		fmt.Println(v)
	}
}

//插入 稳定
func sort3(arr []int) {
	len := len(arr)
	for i, j := 0, 0; i < len; i++ {
		curr := arr[i]
		for j = i - 1; j >= 0 && arr[j] > curr; j-- {
			arr[j+1] = arr[j]
			fmt.Println("aaaaaaaaaaaa", "  ", j)
		}
		arr[j+1] = curr
		fmt.Println("bbbbbbbbbb", "  ", j)
	}
	for _, v := range arr {
		fmt.Println(v)
	}
}

func sort3_(arr []int) {
	l := len(arr)
	for i, j := 0, 0; i < l; i++ {
		curr := arr[i]
		for j = i - 1; j >= 0 && arr[j] > curr; j-- {
			arr[j+1] = arr[j]
			fmt.Println("ccccccccccc", "  ", j)
		}
		arr[j+1] = curr
		fmt.Println("kkkkkkkkkk", "  ", j)
	}
	for _, v := range arr {
		fmt.Println(v)
	}
}

//快速 不稳定
func sort4(arr []int, lo int, hi int) []int {
	if lo >= hi {
		return arr
	}
	p := partition(arr, lo, hi)
	sort4(arr, lo, p-1)
	sort4(arr, p+1, hi)
	return arr
}

func partition(nums []int, lo int, hi int) int {
	swap(nums, lo+rand.Intn(hi-lo), hi)
	var i, j int
	for i, j = lo, lo; j < hi; j++ {
		if nums[j] <= nums[hi] {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, j)
	return i
}

//快速 不稳定

func quickSort(nums []int, lo int, hi int) {
	if lo >= hi {
		return
	}
	p := partition_(nums, lo, hi)
	quickSort(nums, lo, p-1)
	quickSort(nums, p+1, hi)
}

func partition_(nums []int, lo int, hi int) int {
	swap(nums, lo+rand.Intn(hi-lo), hi)
	var i, j int
	for i, j = lo, lo; j < hi; j++ {
		if nums[j] <= nums[hi] {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, j)
	return i
}

func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

//归并 稳定
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

//归并 稳定
func sort6(arr []int) []int {
	len := len(arr)
	if len == 1 {
		return arr
	}
	m := len / 2
	left := sort6(arr[:m])
	right := sort6(arr[m:])
	return merge(left, right)
}

func merge2(l []int, r []int) []int {
	lenL := len(l)
	lenR := len(r)
	res := make([]int, 0)
	indexL, indexR := 0, 0
	for indexL < lenL && indexR < lenR {
		if l[indexL] > r[indexR] {
			res = append(res, r[indexR])
			indexR++
		} else {
			res = append(res, l[indexL])
			indexL++
		}
	}
	if indexL < lenL { //左边的还有剩余元素
		res = append(res, l[indexL:]...)
	}
	if indexR < lenR {
		res = append(res, r[indexR:]...)
	}
	return res
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

func main() {
	arr := []int{9, 455, 34, 5, 1, 2, 5, 7, 8, 6}
	//sort1(arr)
	//sort2(arr)
	//sort3(arr)
	//ints := sort4(arr, 0, len(arr)-1)
	//rsult := sort5(arr)
	//for _, v := range rsult {
	//	fmt.Println("归并---", v)
	//}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
