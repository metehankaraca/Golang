package main

import "fmt"

func main() {

	//name := "Ali"

	/* 	   	fmt.Println(name)
	   	   	fmt.Println(&name) */ // & --> address operator

	/* 	fmt.Println(x)
	   	fmt.Println(&x)
	   	fmt.Println()
	   	fmt.Printf("%T, %v\n", x, x)
	   	fmt.Printf("%T, %v\n", &x, &x) */

	/* 	y := &x
	   	fmt.Printf("%T, %v\n", y, y)
	   	z := &name
	   	fmt.Printf("%T, %v\n", z, z) */

	/* 	x := 22 */
	/* 	fmt.Println(x)        // 22
	   	fmt.Println(&x)       // 22 nin adresi &x ---> *int
	   	fmt.Println(*(&x))    // dereferencing
	   	fmt.Println(&(*(&x))) // * --->> ilgili adresteki değeri gösteriri
	   	fmt.Println(*(&(*(&x))))
		   fmt.Println(3 * 5) */

	/* 	x1 := 10
	   	x2 := x1
	   	fmt.Println(x1, x2)
	   	x1 = 5
	   	fmt.Println(x1, x2) */

	/* 	x1 := 10
	   	x2 := &x1
	   	fmt.Println(x1, x2)
	   	fmt.Println(x1, *x2)
	   	*x2 = 3
	   	fmt.Println(x1, *x2)
	   	x3 := &x1
	   	*x3 = 5
	   	fmt.Println(x1, *x2, *x3) */

	// x1 := [4]int{1, 10, 100, 1000}  // array pass by value

	x1 := []int{1, 10, 100, 1000}
	x2 := x1

	fmt.Println(x1, x2)

	x2[0] = 3
	fmt.Println(x2)
	fmt.Println(x1) // slice pass by reference

	fmt.Println()

	main1()
	main2()
	main3()
	main4()
}

func main1() { // GO pass by value
	x := 5
	fmt.Println(x)
	double1(x)
	fmt.Println(x)
}
func double1(num int) {
	num *= 2
	fmt.Println(num)
}

func main2() {
	mySlc := []int{1, 10, 100}
	fmt.Println(mySlc)
	double2(mySlc)
	fmt.Println(mySlc)
}
func double2(slc []int) {
	for i := 0; i < len(slc); i++ {
		slc[i] *= 2
	}
	fmt.Println(slc)
}

func main3() {
	myArr := [3]int{1, 10, 100}
	fmt.Println(myArr)
	double3(myArr)
	fmt.Println(myArr)
}
func double3(arr [3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
	fmt.Println(arr)
}

func main4() { // GO pass by value
	x := 5
	fmt.Println(x)
	double5(&x)
	fmt.Println(x)

}
func double5(num *int) { // double --> pointer method
	fmt.Println(num)
	*num *= 2
	fmt.Println(*num)
}
