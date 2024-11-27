package main

import "fmt"

func main() {

	var bil int

	//bikin input
	fmt.Print("Bilangan:")
	fmt.Scan(&bil)

	//rumus mencari faktor bilangan
	fmt.Print("Faktor:")
	jumlahFaktor := 0
	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			fmt.Print(i, "")
			jumlahFaktor++
		}
	}
	fmt.Println()

	isPrima := jumlahFaktor == 2
	//output hasil

	fmt.Println("prima", isPrima)
}
