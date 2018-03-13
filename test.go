package main

import (
	"fmt"
)

func qsort(data []int) {
	if len(data) <= 1 {
		return 
	}
	mid, i := data[0], 1
	head, tail := 0, len(data) - 1
	for i = 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	data[head] = mid
	qsort(data[:head])
	qsort(data[head+1:])
}

func main() {
	a := []int{4,5,2,8,9,11}
	qsort(a)
	fmt.Println(a)
}
