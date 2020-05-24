package main

import "fmt"

func main() {
	var positiveNumber uint8 = 89
	var negativeNumber = -12123123232
	var decimalNumber = 2.62
	var exist bool = true
	// var message string = "halo"
	var message string = `Nama saya "Ervan Antonius"
	Salam kenal.
	Mari belajar 'Golang'
	`

	const konstanta string = "Ervan"
	const konstanta2 = "Antonius"

	fmt.Printf("bilangan positif: %d \n", positiveNumber)
	fmt.Printf("bilangan negatif: %d \n", negativeNumber)
	fmt.Printf("bilangan desimal: %f \n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f \n", decimalNumber)
	fmt.Printf("exist? %t \n", exist)
	fmt.Printf("message: %s \n", message)
	fmt.Print("konstanta: ", konstanta, "! \n")
	fmt.Print("konstanta2: ", konstanta2, "! \n")

	fmt.Println("Ervan Antonius")
	fmt.Println("Ervan", "Antonius")
	fmt.Print("Ervan Antonius\n")
	fmt.Print("Ervan ", "Antonius\n")
	fmt.Print("Ervan", " ", "Antonius\n")
}