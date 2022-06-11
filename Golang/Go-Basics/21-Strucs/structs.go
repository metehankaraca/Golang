package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Name    string `json:"name"`
	Surname string `json:"-"`
	Likes   []Like `json:"like,omitempty"`
}

type Like struct {
	ID   string
	Date time.Time
}

func main() {
	var employe struct {
		name      string
		age       int
		isMarried bool
	}

	employe.name = "Ali"
	employe.age = 30
	employe.isMarried = true

	fmt.Printf("%#v\n", employe)
	fmt.Println(employe.name)
	fmt.Println(employe.age)
	fmt.Println(employe.isMarried)
	fmt.Println()

	var employe2 struct {
		name      string
		age       int
		isMarried bool
	}

	employe.name = "Ayşe"
	employe.age = 5
	employe.isMarried = false

	fmt.Println(employe2)
	fmt.Println()

	main2()
	main3()
	main4()

}

func main2() {
	type employee struct {
		name      string
		age       int
		isMarried bool
	}

	var e1 employee
	e1.name = "Mete"
	e1.age = 20
	e1.isMarried = false

	var e2 employee
	e2.name = "Tuğçe"
	e2.age = 45
	e2.isMarried = true

	e3 := employee{
		name:      "Murat",
		age:       17,
		isMarried: false,
	}

	fmt.Printf("%#v\n", e1)
	fmt.Printf("%#v\n", e2)
	fmt.Printf("%#v\n", e3)
	fmt.Println()
}

func main3() {
	u := User{
		Name:    "go",
		Surname: "turkiye",
		Likes: []Like{
			{
				Date: time.Now(),
			},
			{
				Date: time.Now(),
			},
			{
				Date: time.Now(),
			},
			{
				Date: time.Now(),
			},
		},
	}

	arr, _ := json.Marshal(u)
	fmt.Println(string(arr))
	fmt.Println()

	fmt.Println(u)
	fmt.Println()
}

type Employe1 struct {
	name      string
	age       int
	isMarried bool
}

type Manager struct {
	Employe1
	hasDegree bool
}

func main4() {
	m1 := Manager{
		Employe1: Employe1{
			name:      "Zeynep",
			age:       35,
			isMarried: false,
		},
		hasDegree: true,
	}
	fmt.Println(m1)
	fmt.Println()
}
