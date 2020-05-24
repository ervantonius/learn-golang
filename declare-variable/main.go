package main

import "fmt"

func main() {
	// var firstName string = "Ervan"
	var firstName = "Ervan"

	// var lastName string
	// lastName = "Antonius"
	lastName := "Antonius"

	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	var fourth, fifth, sixth string = "empat", "lima", "enam"

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	_ = "belajar variabel sampah"

	// pointer := new(string)
	pointer := new(int)

	fmt.Printf("halo %s %s! \n", firstName, lastName)
	fmt.Println("halo", firstName, lastName + "!")
	fmt.Println(first, second, third)
	fmt.Println(fourth, fifth, sixth)
	fmt.Println(seventh, eight, ninth)
	fmt.Println(one, isFriday, twoPointTwo, say)
	fmt.Println(pointer)
	fmt.Println(*pointer)
}