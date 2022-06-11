package main

import "fmt"

// 1-) 1-10 arasındaki sayıları if yapısıyla tek - çift olarak yazdırınız.

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even number")
		} else {
			fmt.Println(i, "odd number")
		}
	}
	fmt.Println("")

	main2()
	main3()
	main4()
	main5()

}

// 2-) for yapısı kullanarak Go'da olmayan while döngüsüne örnek veriniz.

func main2() {
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("")
}

// 3-) Switch fallthrough ifadesini açıklayın.
func main3() {
	switch x := 25; {
	case x < 20:
		fmt.Printf("%d küçüktür 20\n", x)
		fallthrough
	case x < 50:
		fmt.Printf("%d küçüktür 50\n", x)
		fallthrough
	case x < 100:
		fmt.Printf("%d küçüktür 100\n", x)
		fallthrough
	case x < 200:
		fmt.Printf("%d küçüktür 200\n", x)
	}
	fmt.Println("")
}

// 4-) Aşağıdaki if döngüsünü daha idiomatic hale getiriniz.
// func main() {
// 	if x := 20; x%2 == 0 {
// 		fmt.Println(x, "Çifttir.")
// 	} else {
// 		fmt.Println(x, "Tektir.")
// 	}
// }
func main4() {
	x := 20
	if x%2 == 0 {
		fmt.Println(x, "Çifttir.")
		return
	}
	fmt.Println(x, "Tektir.")
	fmt.Println("")
}

// 5-) 1-50 arasındaki asal sayıları gösteren bir program yazınız.
func main5() {
	var x, y int

	for x = 2; x < 50; x++ {
		for y = 2; y < (x / y); y++ {
			if x%y == 0 {
				break
			}
		}
		if y > (x / y) {
			fmt.Printf("%d bir asal sayidir.\n", x)
		}
	}
}
