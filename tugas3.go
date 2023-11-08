package main

import "fmt"

func main() {
	jarak := 63.0
	konsumsiBensin := 17.0
	hargaBensin := 8350.0
	totalBensin := jarak / konsumsiBensin
	biayaBensin := totalBensin * hargaBensin

	fmt.Printf("%.2f Rupiah\n", biayaBensin)
}
