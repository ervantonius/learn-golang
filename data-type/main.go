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

	fmt.Println()
	fmt.Println()
	fmt.Println("--ARRAY--")

	var names[4]string
	names[0] = "Ervan"
	names[1] = "Antonius"
	names[2] = "Agung"
	names[3] = "Pratama"

	fmt.Println(names[0], names[1], names[2], names[3])

	fmt.Println()

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}
	var fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	fmt.Println()

	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t", numbers)
	fmt.Println("jumlah element \t", len(numbers))

	fmt.Println()

	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1 \t", numbers1)
	fmt.Println("numbers2 \t", numbers2)

	fmt.Println()

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("element %d : %s \n", i, fruits[i])
	}

	fmt.Println()

	for i, fruit := range fruits {
		fmt.Printf("element %d : %s \n", i, fruit)
	}

	fmt.Println()
	
	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s \n", fruit)
	}

	fmt.Println()

	// for i, _ := range fruits {
	for i := range fruits {
		fmt.Printf("element %d \n", i)
	}

	fmt.Println()

	var fruits2 = make([]string, 2)
	fruits2[0] = "apple"
	fruits2[1] = "manggo"

	fmt.Println(fruits2)

	fmt.Println()
	fmt.Println()
	fmt.Println("--SLICE--")

	var fruitsSlice = []string{"apple", "grape", "banana", "melon"}
	var aFruits = fruitsSlice[0:3]
	var bFruits = fruitsSlice[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruitsSlice)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)

	baFruits[0] = "pineapple"

	fmt.Println(fruitsSlice)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)

	fmt.Println()
	
	fmt.Println(len(fruitsSlice))
	fmt.Println(cap(fruitsSlice))
	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))
	
	fmt.Println()

	var cFruits = append(fruitsSlice, "papaya")

	fmt.Println(cFruits);

	fmt.Println()

	dst := make([]string, 3)
	src := []string{"watermelon", "pineapple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	fmt.Println(fruitsSlice[0:2:3])
	fmt.Println(len(fruitsSlice[0:2:3]))
	fmt.Println(cap(fruitsSlice[0:2:3]))

	fmt.Println()
	fmt.Println()
	fmt.Println("--MAP--")

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40
	chicken["maret"] = 30
	chicken["april"] = 20

	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"])

	for key, val := range chicken {
		fmt.Println(key, " \t: ", val)
	}


	fmt.Println(chicken)
	delete(chicken, "april")
	fmt.Println(chicken)

	var value, isExist = chicken["april"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}
}