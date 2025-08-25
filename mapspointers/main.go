package main

import (
	"fmt"
)

// pointers
func Marks(mark *int, change int) {
	*mark += change
	if *mark > 100 {
		*mark = 100
	}
	if *mark < 0 {
		*mark = 0
	}
}

// Maps
func Groceries() {
	groceries := map[string]float64{
		"Chicken": 11.29,
		"Mutton":  29.20,
		"Fish":    18.39,
	}

	groceries["vegetables"] = 45.99

	fmt.Println("\nGroceries:")
	for name, price := range groceries {
		fmt.Println(name, "-> $", price)
	}
}

// Struct + Methods
type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Resize(newWidth, newHeight float64) {
	r.width = newWidth
	r.height = newHeight
}

func main() {

	mathmatics := 70
	fmt.Println("Maths before finalexam :", mathmatics)
	Marks(&mathmatics, 20)
	fmt.Println("Maths after finalexam :", mathmatics)

	Groceries()

	rect := Rectangle{width: 10, height: 5}
	fmt.Println("\nRectangle area:", rect.Area())

	rect.Resize(20, 10)
	fmt.Println("Rectangle area rebuilt :", rect.Area())
}
