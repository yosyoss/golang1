package main

import "fmt"

func bubbleSort(arr [10]int) [10]int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := [10]int{99, 2, 64, 8, 111, 33, 65, 11, 102, 50}
	arr = bubbleSort(arr)
	fmt.Println(arr)
}
