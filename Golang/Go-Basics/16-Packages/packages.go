package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println(time.Now())

	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	fmt.Println("")

	fmt.Println(strings.Count("animatrix", "a"))
	fmt.Println(strings.ToUpper("Gopher"))
}
