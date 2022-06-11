package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// 1-) İki rakam arasında toplama, çıkarma, çarpma işleminin yapıldığı bir fonksiyon yazınız.
func main() {
	x, y := 10, 4

	sum, dif, prod := calculation(x, y)
	fmt.Println("Toplam : ", sum)
	fmt.Println("Fark : ", dif)
	fmt.Println("Çarpim : ", prod)
	fmt.Println("")

	//
	main2()
	main3()
}

func calculation(num1, num2 int) (int, int, int) {
	sum := num1 + num2
	dif := num1 - num2
	prod := num1 * num2

	return sum, dif, prod
}

// 2-) Kullanıcı tarafından girilen nota göre geçtiniz veya kaldınız geri dönüşü yazdırınız.
func main2() {
	fmt.Print("Lütfen notunu giriniz: ")
	grade, _ := getGrade()
	var result string

	if grade >= 50 {
		result = "Geçti"
	} else {
		result = "Kaldi"
	}

	fmt.Println(result)
	fmt.Println("")

}

func getGrade() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}
	return num, nil
}

// 3-) 1 ile 100 arasındaki bir sayıyı tahmin etme uygulaması yazınız. Toplam hakkınız 10 olsun.
func main3() {
	target := numRand(1, 100)
	fmt.Println("1 ile 100 arasindaki sayiyi tahmin ediniz: ")

	reader := bufio.NewReader(os.Stdin)

	for attempts := 0; attempts < 10; attempts++ {
		fmt.Println(10-attempts, "hakkiniz kaldi.")
		fmt.Print("Lütfen tahmininizi yapiniz: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}

		if num > target {
			fmt.Println("Tahmininiz çok oldu, daha küçük bir sayi giriniz.")
		} else if num < target {
			fmt.Println("Tahmininiz az oldu, daha büyük bir sayi giriniz.")
		} else {
			fmt.Println("Doğru Tahmin, hedef sayi ", target, "idi", attempts, ". seferde buldunuz.")
			break
		}
	}

}

func numRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
