package main

import (
	"fmt"
)

// structure using full details
type All struct {
	Fullname   string
	DOB        string
	education  string
	Add        string
	Street     string
	City       string
	State      string
	postalcode int
}

func main() {
	// Create a struct variable
	details := All{
		Fullname:   "Ravichandra",
		DOB:        "10/23",
		education:  "Masters",
		Add:        "1201 W Whitestone Blvd Apt 101",
		Street:     "Whitestone Crossing",
		City:       "Austin",
		State:      "TX",
		postalcode: 78613,
	}

	fmt.Println("Full Details:")
	fmt.Println("Fullname:", details.Fullname)
	fmt.Println("DOB:", details.DOB)
	fmt.Println("education:", details.education)
	fmt.Println("Add:", details.Add)
	fmt.Println("Street:", details.Street)
	fmt.Println("City:", details.City)
	fmt.Println("State:", details.State)
	fmt.Println("PostalCode:", details.postalcode)

	//array
	a, b := 20, 10
	for i := 1; i <= 4; i++ {
		if i == 1 {
			fmt.Println("Addition :", a+b)
		}
		if i == 2 {
			fmt.Println("Subtraction :", a-b)
		}
		if i == 3 {
			fmt.Println("multiplication :", a*b)
		}
		if i == 4 {
			fmt.Println("Division :", a/b)
		}
	}

	//slices
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("slice Value :", nums)

	nums = append(nums, 10, 12)
	fmt.Println("slice after appending 10 and 12:", nums)

	nums = append(nums, 20, 40, 60)
	fmt.Println("slice after appending 20, 40, 60:", nums)

}
