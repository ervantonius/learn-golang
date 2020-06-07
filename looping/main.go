package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	fmt.Print("\n")

	var j = 0
	for j < 5 {
		fmt.Println("Angka", j)
		j++
	}

	fmt.Print("\n")

	var k = 0
	for {
		fmt.Println("Angka", k)

		k++
		if k == 5 {
			break
		}
	}

	fmt.Print("\n")

	for i := 0; i <= 10; i++ {
		if i % 2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	fmt.Print("\n")

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	fmt.Print("\n")

	outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}