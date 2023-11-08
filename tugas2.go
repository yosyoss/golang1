package main

import "fmt"

func main() {
	absensi := 16
	tugasHarian := 70
	uts := 80
	uas := 70

	bobotAbsensi := 0.10
	bobotTugasHarian := 0.20
	bobotUTS := 0.30
	bobotUAS := 0.40

	nilaiAbsensi := float64(absensi) * bobotAbsensi
	nilaiTugasHarian := float64(tugasHarian) * bobotTugasHarian
	nilaiUTS := float64(uts) * bobotUTS
	nilaiUAS := float64(uas) * bobotUAS

	nilaiTotal := nilaiAbsensi + nilaiTugasHarian + nilaiUTS + nilaiUAS

	fmt.Printf("%.2f", nilaiTotal)
}
