package main

import "fmt"

func main() {
	// for <init statement>;  <condition>; <post statement> { code }

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("")

	// for { // Infinite Loop
	// 	fmt.Println("")
	// }

	i := 10
	for i >= 0 {
		fmt.Println(i)
		i--
	}
	fmt.Println("")

	for i := 0; i <= 10; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}
	slc_1 := []int{1, 2, 3}
	slc_2 := []int{}

	for i, value := range slc_1 {
		fmt.Printf("index: %d value: %d \n", i, value)
	}

	for i := 0; i < 10; i++ {
		slc_2 = append(slc_2, i)
	}

	fmt.Println(slc_1)
	fmt.Println(slc_2)
	fmt.Printf("slc_1 len:%d cap:%d \n", len(slc_1), cap(slc_1))
	fmt.Printf("slc_2 len:%d cap:%d \n", len(slc_2), cap(slc_2))

}
