package main

import (
	"fmt"
)

func nilaiTerbesar(angka []int) int {
	if len(angka) == 0 {
		return 0
	}
	max := angka[0]

	for _, num := range angka {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	angka := []int{11, 6, 31, 201, 99, 861, 1, 7, 14, 79}
	terbesar := nilaiTerbesar(angka)
	fmt.Printf("%d\n", terbesar)
}
