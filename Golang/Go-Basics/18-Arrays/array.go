package main

import "fmt"

var arr_1 [3]int
var arr_2 = [5]int{1, 2, 3, 4, 5}

func main() {
	arr_3 := make([]int, 3)
	arr_3[1] = 2

	fmt.Println(arr_1, arr_2, arr_3)
	fmt.Printf("arr_1 len:%d \n", len(arr_1))
	fmt.Printf("arr_2 len:%d \n", len(arr_3))
	fmt.Println("")

	// city1 := "İstanbul"
	// city2 := "Roma"
	// city3 := "belgrad"
	// city4 := "Los Angeles"
	// fmt.Println(city1, city2, city3, city4)

	cities := [4]string{"İstanbul", "Roma", "Belgrad", "San Diego"}
	fmt.Println(cities)
	fmt.Println(cities[0]) // zero based index
	fmt.Println(cities[3])
	fmt.Println("")

	cities1 := []string{"Ankara", "Washington DC", "Rochester", "Beverly Hills"}
	fmt.Println(cities1)
	fmt.Println(cities1[0]) // zero based index
	fmt.Println(cities1[3])
	fmt.Println("")

	cities2 := [...]string{"Londra", "Dublin", "Kopenhag", "Amsterdam"}
	fmt.Println(cities2)
	fmt.Println(len(cities2))
	fmt.Println("")

	var myArr [5]int
	fmt.Println(myArr)

	myArr[0] = 5
	myArr[2] = 30
	myArr[len(myArr)-1] = 500
	fmt.Println(myArr)
	fmt.Println("")

	countries := [4]string{"USA", "Englang", "Netherlands", "Norway"}

	for i := 0; i < len(countries); i++ {
		fmt.Println(i, countries[i])
	}
	fmt.Println("")

	countries[1] = "Germany"
	for i := 0; i < len(countries); i++ {
		fmt.Println(i, countries[i])
	}
	fmt.Println("")

	numberArr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	numberArr = mySquare(numberArr) // First Class Functions
	fmt.Println(numberArr)
	fmt.Println("")

	// For -- Range
	// for index, value := range myArr

	cars := [4]string{"bmw", "ford", "fiat", "honda"}
	for i, car := range cars {
		fmt.Println(i, car)
	}
}

func mySquare(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}
