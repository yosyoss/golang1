package main

import "fmt"

func main() {
	totalBelanja := 13500
	uangDiberikan := 100000

	uangKembalian := uangDiberikan - totalBelanja
	//fmt.Printf("%d", uangKembalian)

	uangPecahan := []int{
		50000,
		20000,
		10000,
		5000,
		2000,
		1000,
		500,
	}
	for _, pecahan := range uangPecahan {
		jumlahPecahan := uangKembalian / pecahan
		uangKembalian %= pecahan
		fmt.Printf("%d \n", jumlahPecahan)
	}
}
