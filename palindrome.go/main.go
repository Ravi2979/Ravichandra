package main

import (
	"fmt"
	"time"
)

func main() {
	var name string
	fmt.Print("Enter a word: ")
	fmt.Scanln(&name)

	distance := len(name)

	a, b := 1, 2
	isFibonacci := false
	for a <= distance {
		if a == distance {
			isFibonacci = true
			break
		}
		a, b = b, a+b
	}

	if isFibonacci {
		rev := ""
		for i := distance - 1; i >= 0; i-- {
			rev = rev + string(name[i])
		}
		fmt.Println("Reversed:", rev)

		if name == rev {
			fmt.Println(name, "is a palindrome")
		} else {
			fmt.Println(name, "is not a palindrome")
		}
	} else {
		fmt.Println("Name is not Fibonacci sequence It has no reversed")
	}

	// Go routine I was printing different lang
	languages := []string{
		"Hola (Spanish)",
		"Namaste (Telugu)",
		"Namaskar (Hindi)",
		"Hello (English)",
		"Ravi (I'm trying to print different lang)",
	}

	for i := 0; i < len(languages); i++ {
		go fmt.Println(languages[i])
	}

	time.Sleep(1 * time.Millisecond)

	//factorial program

	var f int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&f)

	fact := 1
	for i := 1; i <= f; i++ {
		fact = fact * i
	}

	fmt.Println("Factorial of", f, "is", fact)

	// count the character's

	letters := []string{"go", "run", "fast", "go"}

	goCount := 0
	runCount := 0
	fastCount := 0

	for i := 0; i < len(letters); i++ {
		if letters[i] == "go" {
			goCount++
		} else if letters[i] == "run" {
			runCount++
		} else if letters[i] == "fast" {
			fastCount++
		}
	}

	// print results
	fmt.Println("go -", goCount)
	fmt.Println("run -", runCount)
	fmt.Println("fast -", fastCount)
}
