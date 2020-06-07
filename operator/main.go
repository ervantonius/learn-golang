package main

import "fmt"

func main() {
	var value = (((2 + 6) % 3) * 4 - 2) / 3
	
	var isEqual = (value == 2)
	var isGreaterEqual = (value >= 2)
	var isSmallerEqual = (value <= 2)
	var isSmaller = (value < 2)
	var isGreater = (value > 2)
	var isNotEqual = (value != 2)

	var left = false
	var right = true
	var leftAndRight = left && right
	var leftOrRight = left || right
	var leftReserve = !left
	var rightReserve = !right

	
	fmt.Println(value);

	fmt.Printf("nilai %d isEqual (%t) \n", value, isEqual);
	fmt.Printf("nilai %d isGreaterEqual (%t) \n", value, isGreaterEqual);
	fmt.Printf("nilai %d isSmallerEqual (%t) \n", value, isSmallerEqual);
	fmt.Printf("nilai %d isSmaller (%t) \n", value, isSmaller);
	fmt.Printf("nilai %d isGreater (%t) \n", value, isGreater);
	fmt.Printf("nilai %d isNotEqual (%t) \n", value, isNotEqual);

	fmt.Printf("nilai %d isNotEqual (%t) \n", value, isNotEqual);
	fmt.Printf("left && right \t (%t) \n", leftAndRight)
	fmt.Printf("left || right \t (%t) \n", leftOrRight)
	fmt.Printf("!left \t\t (%t) \n", leftReserve)
	fmt.Printf("!right \t\t (%t) \n", rightReserve)
}