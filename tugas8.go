package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [10]int{55, 31, 20, 1, 10, 89, 90, 66, 11, 21}

	sort.Ints(arr[:])
	fmt.Println(arr)
}
