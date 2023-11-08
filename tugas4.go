package main

import "fmt"

func main() {
	tinggi := 5

	for i := tinggi; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
