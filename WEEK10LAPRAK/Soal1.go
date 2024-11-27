package main

import "fmt"

func main() {

	//untuk input
	var berat int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	//perhitungan berart dalam kg dan sisa gram
	kg := berat / 1000
	sisa := berat % 1000

	//biaya dalam kg
	biayaPerKg := 10000
	biayaTambahan := 0

	//kondisi untuk sisa berat
	if kg > 10 {
		biayaTambahan = 0
	} else if sisa >= 500 {
		biayaTambahan = sisa * 5
	} else {
		biayaTambahan = sisa * 15
	}

	//hitung totl biaya
	total := (kg * biayaPerKg) + biayaTambahan

	//menampilkan hasil
	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", kg*biayaPerKg, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", total)
}
