package main

import "fmt"

// 1-) 5 string elemandan oluşan bir array ve 5 int elemandan oluşan slice oluşturup index numaralarıyla birlikte gösterelim.

// 2-) [0,1,2,3,4,5,6,7,8] slice dan [0,1,2,3],[4,5,6], [6,7,8],[2,3,6,7] slicelarını oluşturunuz.

// 3-) slicelar için copy methodunu ve assing ( = ) ile farkını açıklayın.

// 4-) map gösterimi ile yazar ve yazarlara ait kitapların isimlerini for range ile gösterin.

func main() {
	fmt.Println("Answer 1")
	myArr := [5]string{"bursa", "istanbul", "ankara", "izmir", "aydin"}

	for index, value := range myArr {
		fmt.Println(index, value)
	}

	fmt.Println()

	mySlc := []int{1, 2, 3, 4, 5}
	for i, v := range mySlc {
		fmt.Println(i, v)
	}

	fmt.Println()

	main2()
	main3()
	main4()
}

func main2() {
	fmt.Println("Answer 2")

	mySlc2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	myPart1 := mySlc2[:4]
	fmt.Println(myPart1)
	myPart2 := mySlc2[4:7]
	fmt.Println(myPart2)
	myPart3 := mySlc2[6:9]
	fmt.Println(myPart3)
	myPart4 := append(mySlc2[2:4], mySlc2[6:8]...)
	fmt.Println(myPart4)
	fmt.Println()
}

func main3() {
	fmt.Println("Answer 3")

	mySlc3 := []int{1, 2, 3}
	mySlc4 := make([]int, 2)

	fmt.Println(mySlc3)
	fmt.Println(mySlc4)

	copy(mySlc4, mySlc3)
	fmt.Println(mySlc3)
	fmt.Println(mySlc4)

	mySlc4 = mySlc3
	fmt.Println(mySlc3)
	fmt.Println(mySlc4)
	fmt.Println()
}

func main4() {
	fmt.Println("Answer 4")

	myMap := map[string][]string{
		"Yaşar Kemal":     {"İnce Memed,", "Yer demir gök bakir"},
		"Ahmet Ümit ":     {"Kardeşimin Hikayesi,", "Beyoğlu Rapsodisi,", "Bab-i Esrar"},
		"Haruki Murakami": {"1Q84,", "Dans Dans Dans,", "Kumandani Öldürmek"},
	}

	for key, value := range myMap {
		fmt.Println("Yazarimiz: ", key)
		fmt.Println("Bazi Kitaplari Aşağidadir: ")
		for i, v := range value {
			fmt.Println("\t", i+1, v)
		}
	}
}
