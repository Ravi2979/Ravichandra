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
type City struct {
	Name       string
	Population int
	areacode   float64
}

func (c City) Info() {
	fmt.Println(c.Name, "has population", c.Population, "and areacode", c.areacode)
}

func (c *City) Grow(extra int) {
	c.Population += extra
}

func main() {

	mathmatics := 70
	fmt.Println("Maths before finalexam :", mathmatics)
	Marks(&mathmatics, 20)
	fmt.Println("Maths after finalexam :", mathmatics)

	Groceries()

	c1 := City{Name: "Ravulapalem", Population: 150000, areacode: 05}
	c1.Info()
	c1.Grow(5000)
	fmt.Println("After growth:")
	c1.Info()
}
