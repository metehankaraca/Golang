package main

import "fmt"

var slc_1 []int
var mySlc1 []int

func main() {

	mySlc := []int{1, 2, 3} // LITERAL
	fmt.Println(mySlc)
	fmt.Println(len(mySlc))

	mySlc1 = make([]int, 4) // MAKE METHOD
	fmt.Println(mySlc1)
	// mySlc1[0] = 10
	fmt.Println("")

	underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //
	fmt.Println(underArray)

	mySlice := underArray[2:5]
	fmt.Println(mySlice)
	mySlice1 := underArray[:6]
	fmt.Println(mySlice1)
	mySlice2 := underArray[3:]
	fmt.Println(mySlice2)
	mySlice3 := underArray[:]
	fmt.Println(mySlice3)
	fmt.Println("")

	slc_2 := make([]int, 0)

	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 1)

	fmt.Println(slc_1, slc_2)
	fmt.Printf("slc_1 len:%d cap:%d \n", len(slc_1), cap(slc_1))
	fmt.Printf("slc_2 len:%d cap:%d \n", len(slc_2), cap(slc_2))
	fmt.Println("")

}
