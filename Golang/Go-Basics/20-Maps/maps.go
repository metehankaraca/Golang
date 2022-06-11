package main

import "fmt"

var map_1 map[int]string

func main() {
	myMap := map[string]int{
		"Ahmet": 40,
		"Ayşe":  17,
		"Selim": 14,
		"Melis": 32,
	}
	fmt.Println(myMap)
	fmt.Println(myMap["Ahmet"], myMap["Selim"])
	fmt.Println("")

	isMarried := map[string]bool{
		"Ahmet": true,
		"Ayşe":  true,
		"Ali":   false,
	}
	fmt.Println(isMarried)
	isMarried["Mete"] = true
	fmt.Println(isMarried)
	delete(isMarried, "Ayşe")
	fmt.Println(isMarried)
	fmt.Println(len(isMarried))

	fmt.Println("")

	map_1 = make(map[int]string)
	map_2 := make(map[int]int)
	map_3 := make(map[int]int, 3)

	map_1[0] = "1"
	map_2[0] = 2
	map_3[0] = 3
	map_3[1] = 3
	map_3[2] = 3
	map_3[3] = 3

	fmt.Println(map_1, map_2, map_3)
	fmt.Printf("map_1 len:%d \n", len(map_1))
	fmt.Printf("map_2 len:%d \n", len(map_2))
	fmt.Printf("map_3 len:%d \n", len(map_3))
	fmt.Println("")

}
