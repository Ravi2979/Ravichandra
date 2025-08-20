package main

import (
	"fmt"
	"ravichandra/simplecal"
)

func main() {

	//add,sub
	a, b := 10, 7
	fmt.Println(simplecal.Add(a, b))
	fmt.Println(simplecal.Subb(a, b))

	//short hand declaration
	d := 30
	fmt.Println("The value of d is:", d)

	//for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//var char
	var r int
	fmt.Println(r)

	//Basic hello world
	fmt.Println("Hello World")

	//If-else
	num := 20
	if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	//greater or less
	rav := 6
	if rav > 10 {
		fmt.Println("greater")
	} else if rav == 10 {
		fmt.Println("equal")
	} else {
		fmt.Println("Less")
	}

	//switch
	week := "Sunday"
	switch week {
	case "monday":
		fmt.Println("First weekday")
	case "Tuesday":
		fmt.Println("Second weekday")
	default:
		fmt.Println("weekend")
	}

	//defer
	defer fmt.Println("Varma")
	fmt.Println("Go")
	defer fmt.Println("Learn")

	//pointer
	c := 15
	p := &c
	fmt.Println(c)
	fmt.Println(p)
	fmt.Println(*p)
	*p = 40
	fmt.Println(c)

}
