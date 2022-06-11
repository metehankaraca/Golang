package main

import "fmt"

/* var x = 10
var customerName
var Y = 100 */

func main() {
	// Print - Println - Printf

	fmt.Print("Print")
	fmt.Println("Println") // -> ln = new space
	fmt.Printf("Printf\n")
	fmt.Println("***************")

	name := "Arin"
	fmt.Print(name)
	fmt.Println(name)
	fmt.Printf(name)
	fmt.Println("*******************")

	fmt.Print("Benim Adim", name)
	fmt.Println("")
	fmt.Println("Benim Adim", name)
	fmt.Printf("Benim Adim %v %T", name, name)
	fmt.Println("*******************")

	x := 100
	y := 20
	z := 30
	fmt.Printf("%b %d %o", x, y, z)
	fmt.Println("*******************")

	name1, age := "Ali", 5
	//fmt.Print("Benim Adim ", name1, ", ve ben ", age, " yasindayim.")
	//fmt.Println("Benim Adim", name1, "ve ben", age, "yasindayim.")
	fmt.Printf("Benim Adim %v, ve ben %v yasindayim", name1, age)

	// Visibility

	/* 	fmt.Println(x)
	   	myFunc() */

	/* 	var coin string  // count - customer - coin
	   	// Go camel case isimlendirme kullanılır
	   	var coinType string
	   	var custName string
	   	// kısaltmalar büyük harflerle yazılır
	   	var URL // Url değil
	   	var HTTP // http değil "xyzHTTP"
	   	i , j , k */

	// benForDongusuDegiskeninİsmiyim

}
