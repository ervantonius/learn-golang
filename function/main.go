package main

import (
	"fmt"
	"strings"
	"math"
	"math/rand"
	"time"
)

func main() {
	var names = []string{"ervan", "antonius"}
	printMessage("halo", names)

	fmt.Println()

	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	fmt.Println()

	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)

	fmt.Println()

	var diameter float64 = 15
	// var area, circumference = calculate(diameter)
	var area, circumference = calculatePredefined(diameter)
	
	fmt.Printf("luas lingkaran \t\t : %.2f \n", area)
	fmt.Printf("keliling lingkaran \t : %.2f \n", circumference)

	fmt.Println()

	// var avg = calculateVariadic(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var avg = calculateVariadic(numbers...)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)

	fmt.Println(msg)

	fmt.Println()

	// yourHobbies("Ervan", "eating", "sleeping")
	var hobbies = []string{"eating", "sleeping"}
	yourHobbies("Ervan", hobbies...)

	fmt.Println()

	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}

		return min, max
	}

	var numbers1 = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers1)
	fmt.Printf("data: %v \nmin : %v \nmax : %v \n", numbers1, min, max)

	fmt.Println()

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers1 {
			if e < min {
				continue
			}

			r = append(r, e)
		}

		return r
	}(3)

	fmt.Println("original number: ", numbers1)
	fmt.Println("filtered number: ", newNumbers)

	fmt.Println()

	var max1 = 3
	var howMany, getNumbers = findMax(numbers1, max1)
	var theNumbers = getNumbers()

	fmt.Println("numbers \t :", numbers1)
	fmt.Printf("find \t : %d \n\n", max1)
	fmt.Println("found \t :", howMany)
	fmt.Println("value \t :", theNumbers)

	fmt.Println()

	var data = []string{"ervan", "antonius", "budi", "asep"}
	var dataContainsE = filter(data, func(each string) bool {
		return strings.Contains(each, "e")
	})
	var dataLenght5 = filter(data, func(each string) bool {
		return len(each) == 5
	})
	fmt.Println("data asli \t\t :", data)
	fmt.Println("filter ada huruf \"e\" \t :", dataContainsE)
	fmt.Println("filter jumlah huruf \"5\" \t :", dataLenght5)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	var value = rand.Int() % (max - min + 1) + min
	return value
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divide. %d cannot divided by %d \n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d \n", m, n, res)
}

func calculate(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d / 2, 2)

	var circumference = math.Pi * d

	return area, circumference
}

func calculatePredefined(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d / 2, 2)

	circumference = math.Pi * d

	return
}

func calculateVariadic(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s \n", name)
	fmt.Printf("My hobbies are: %s \n", hobbiesAsString)
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}

	return len(res), func() []int {
		return res
	}
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}