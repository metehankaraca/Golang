package main

import "fmt"

// Classic
var degisken_1 string
var degisken_2 = "deger 2"

func main() {
	degisken_1 = "deger 1"
	degisken_3 := "deger 3"

	fmt.Println(degisken_1, degisken_2, degisken_3)

	// Multiple
	var (
		dilAdi     = "go"
		konum      = "turkiye"
		kisiSayisi int
	)

	a, b, c, d := 1, 2, "go", true
	fmt.Println(a, b, c, d)

	a, b, c, d = 2, 1, "go turkiye", false
	fmt.Println(a, b, c, d)

	var i, j int = 1, 2
	fmt.Println(i, j)
	i, j = 2, 1
	fmt.Println(i, j)

	fmt.Println(dilAdi, konum, kisiSayisi)

}
