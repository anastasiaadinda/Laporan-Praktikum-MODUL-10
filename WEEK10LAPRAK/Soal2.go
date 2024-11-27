package main

import "fmt"

func main() {
	var nam float64

	fmt.Print("Nilai Akhir Mata Kuliah: ")
	fmt.Scan(&nam)

	if nam <= 40 {
		fmt.Println("E")
	} else if nam > 40 && nam <= 50 {
		fmt.Println("D")
	} else if nam > 50 && nam <= 57.5 {
		fmt.Println("C")
	} else if nam > 57.5 && nam <= 65 {
		fmt.Println("BC")
	} else if nam > 65 && nam <= 72.5 {
		fmt.Println("B")
	} else if nam > 72.5 && nam <= 80 {
		fmt.Println("AB")
	} else if nam > 80 {
		fmt.Println("A")
	}
}
