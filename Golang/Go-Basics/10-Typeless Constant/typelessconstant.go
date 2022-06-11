package main

import "fmt"

func main() {

	const m = 5
	fmt.Printf("%T, %v\n", m, m)

	const x = 3
	var y int16 = 12
	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", x+y, x+y) // int16(x) + y
	fmt.Printf("%T, %v\n", x, x)

	const j = int16(5.2 + 4.8)
	fmt.Printf("%T, %v\n", j, j)

	const k = 3
	const l = 5.6
	fmt.Printf("%T, %v\n", k, k)
	fmt.Printf("%T, %v\n", l, l)
	fmt.Printf("%T, %v\n", k+l, k+l)

	const q = 3
	const w = 5.6
	const e = true
	const r = "test"

	fmt.Printf("%T, %v\n", q, q)
	fmt.Printf("%T, %v\n", w, w)
	fmt.Printf("%T, %v\n", e, e)
	fmt.Printf("%T, %v\n", r, r)

}
