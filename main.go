package main

import (
	"fmt"
	"go_project/calculation"
	"strings"
	"time"
)

type Shop struct {
	name    string
	item    string
	Address Address
}

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

// palindrome

func plaindrome(s string) bool {
	s = strings.ToLower(s)
	reversed := ""

	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return s == reversed
}

//printing diff languages using go routines

func greetEnglish() {
	time.Sleep(50 * time.Millisecond)
	fmt.Println("English → Hello")
}

func greetSpanish() {
	time.Sleep(70 * time.Millisecond)
	fmt.Println("Spanish → Hola")
}

func greetHindi() {
	time.Sleep(40 * time.Millisecond)
	fmt.Println("Hindi → Namaste")
}

func main() {

	go greetEnglish()
	go greetSpanish()
	go greetHindi()
	time.Sleep(200 * time.Millisecond)

	// printing count of the characters
	str := "go run run yes go"
	count := map[string]int{}
	for _, w := range strings.Fields(str) {
		count[w]++
	}
	for w, c := range count {
		fmt.Printf("%s-%d\n", w, c)
	}

	word := "DAD"
	if plaindrome(word) {
		fmt.Println("Yes its a plaindrome", word)
	} else {
		fmt.Println("Not a plaindrome", word)
	}
	//factorial program
	Number := 10
	Factorial := 1

	for i := 1; i < Number; i++ {
		Factorial *= i
	}
	fmt.Println(Number, Factorial)

	//Fibonacci series
	num := 10
	e, d := 0, 1
	fmt.Printf("Fibonaci series %d terms:", num)
	for i := 0; i < num; i++ {
		fmt.Printf("%d ", e)
		e, d = d, e+d
	}
	fmt.Println()

	//maps concept

	plan := make(map[string]string)

	plan["p1"] = "cricket"
	plan["p2"] = "fishing"
	plan["p3"] = "badminton"

	fmt.Println("number of plans:", plan)

	plan["p4"] = "dancing"
	fmt.Println("After adding p4 number of plans:", plan)

	plan["p2"] = "Hunting"
	fmt.Println("After Updating p2 number of plans:", plan)

	if v, exits := plan["p2"]; exits {
		fmt.Println("plan is still there:", v)
	} else {
		fmt.Println("plan is still not there:", v)
	}

	for i, k := range plan {
		fmt.Println(i, k)
	}

	sh := Shop{
		name:    "walmart",
		item:    "coconut",
		Address: Address{Street: "8769 tax street", City: "frisco", State: "tx", Zip: "20286"},
	}

	fmt.Println("Name:", sh.name)
	fmt.Println("Name:", sh.item)
	fmt.Println("Address:", sh.Address.Street, sh.Address.City, sh.Address.State, sh.Address.Zip)

	//slices concept
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s)

	s = make([]int, 3, 5)
	s = append(s, 4, 5)
	s = append(s, 6)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	//Array concept
	c := [10]string{"cow", "apple", "biscute", "banana"}

	for i, v := range c {
		fmt.Println(i, v)
	}

	//pointers

	i := 5
	j := &i
	fmt.Println("the value of i:", i)
	fmt.Println("the address of i:", &i)
	fmt.Println("the value of j:", j)
	fmt.Println("the value of j:", *j)
	*j = 12
	fmt.Println("the value of i:", i)

	//switch case in go
	product := 2
	switch product {
	case 1:
		fmt.Println("chocolate")

	case 2:
		fmt.Println("chips")
	case 3:
		fmt.Println("icecream")

	}
	a := 5

	//if else if
	if a > 10 {
		fmt.Println(a, "is positive")
	} else if a < 10 {
		fmt.Println(a, "is negative")
	} else {
		fmt.Println(a, "is zero")
	}

	// if else

	if a%2 == 0 {

		fmt.Println("the number is even", a)

	} else {
		fmt.Println("The number is odd", a)
	}

	//for loop

	for i = 0; i < 5; i++ {
		fmt.Println(i)
	}

	a, b := 10, 20
	fmt.Println("Addition:", calculation.Add(a, b))
	fmt.Println("Multiplication:", calculation.Mul(a, b))
	fmt.Println("Sub:", calculation.Sub(a, b))
	fmt.Println("Dividison:", calculation.Div(float32(a), float32(b)))

}
