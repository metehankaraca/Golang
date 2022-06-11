// Function vs Method(Parameters, arguments,multiple return)

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// hello("Mete", 25) // Arguments run function

	var moment time.Time = time.Now() // Now --> method
	fmt.Println(moment)
	fmt.Println("")

	timer(time.After(2 * time.Second))
	fmt.Println("")

	fmt.Print("Please enter your exam result: ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') // _ blank identifier
	fmt.Println(value)
	fmt.Println("")

	section, remainder := divide(104, 5)
	fmt.Println(section, remainder)
}

// func hello(name string, age int) { // Parameters write function
// 	fmt.Printf("My name is %s, my age %d\n", name, age)
// }

func timer(c <-chan time.Time) {
	for {
		select {
		case <-c:
			return
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}
	}
}

// 104 / 5 =====> 20 - 4
func divide(dividing, divisor int) (section, remainder int) {
	section = dividing / divisor
	remainder = dividing % divisor

	return section, remainder
}
