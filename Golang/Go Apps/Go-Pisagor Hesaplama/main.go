// Pisagor Hesaplama

package main

import (
	"fmt"
	"math"
)

func main() {
	var kenar1, kenar2, hipotenus float64

	fmt.Println("Dik ucgen Hipotenus hesaplama programina hosgeldiniz...")
	fmt.Print("1. kenari giriniz: ")
	fmt.Scan(&kenar1)

	fmt.Print("2. kenari giriniz: ")
	fmt.Scan(&kenar2)

	hipotenus = math.Sqrt(math.Pow(kenar1, 2) + math.Pow(kenar2, 2))
	fmt.Printf("\nHipotenus: %f", hipotenus)

}
