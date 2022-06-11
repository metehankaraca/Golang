package main

import "fmt"

func main() {
	// The switch structure is a simple version of the if structure
	grade := 3

	switch grade {
	case 5: // if grade == 5 { fmt.println("perfect student") }
		fmt.Println("Congrats! Perfect Student.")
		break // If we don't add break it will be added by default.
	case 4: // if grade == 4 { fmt.println("perfect student") }
		fmt.Println("Congrats! Good Student.")
	case 3: // if grade == 3 { fmt.println("perfect student") }
		fmt.Println("Congrats! Normal Student.")
	case 2: // if grade == 2 { fmt.println("perfect student") }
		fmt.Println("Sorry! Bad Student.")
	case 1: // if grade == 1 { fmt.println("perfect student") }
		fmt.Println("Sorry! Very Bad Student.")
	default:
		fmt.Println("Not Found!")
	}

	// if grade == 5 {
	// 	fmt.Println("Congrats! Perfect Student.")
	// } else if grade == 4 {
	// 	fmt.Println("Congrats! Good Student.")
	// } else if grade == 3 {
	// 	fmt.Println("Congrats! Normal Student.")
	// } else if grade == 2 {
	// 	fmt.Println("Sorry! Bad Student.")
	// } else if grade == 1 {
	// 	fmt.Println("Sorry! Very Bad Student.")
	// } else {
	// 	fmt.Println("Not Found!")
	// }
	fmt.Println("")

	// Declaring variables in an switch structure
	switch note := 5; note {
	case 5:
		fmt.Println("Congrats! Perfect Student.")
		x := 100
		fmt.Println(x)
	case 4:
		fmt.Println("Congrats! Good Student.")
	case 3:
		fmt.Println("Congrats! Normal Student.")
		j := 70
		fmt.Println(j)
	case 2:
		fmt.Println("Sorry! Bad Student.")
		// fmt.Println(j) => error
	case 1:
		fmt.Println("Sorry! Very Bad Student.")
	default:
		fmt.Println("Not Found!")
	}
}
