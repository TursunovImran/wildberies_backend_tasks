package main

import (
	"fmt"
	"sync"
)

type SomeStruct struct {
	ID int 
	Numbers []int
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}


func main() {
	var wg sync.WaitGroup

	arr := []SomeStruct{
		{ID: 1, Numbers: []int{4, 2, 9, 1, 5}},
		{ID: 2, Numbers: []int{7, 3, 8, 6}},
		{ID: 3, Numbers: []int{10, 15, 12}}}

	for i, _ := range(arr) {
		wg.Add(1)
		k := i
		go func(){
			arr[k].Numbers = mergeSort(arr[k].Numbers)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(arr)
}
