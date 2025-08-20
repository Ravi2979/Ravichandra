# Ravichandra
This is go lang program.

//Day - 4 Assignment
**Functions**
It is a standalone block of code that perform a spefic task. It will take zero or more arguments as input and it will perform operations can optionally return one or more values. It helps to reuse code and keep ir organized.
It will declare with keyword func (function declaration - Interms of code maintainbility, reusability and flexibility)
**syntax**
func functionName(parameters) returnType{
      ------code------
      }
**Example**
func add(a int, b int) int {
    retuen a+b
    }
In real-Time function used for utility tasks-----calculation, DB query, API call.


**Method**
Method is a function with a receiver like it is associated with a type called struct, alias and soon.
It is useful for behavior and encapsulation.
**Syntax**
func(receiverName ReceiverType) MethodName(parameters) returnType.
**Example**
type person struct{
name string
age int
}
//method//
type Rectangle struct{
width float64
height float 64
}
func (r Rectangle) area() float64{
return r.width*r.height
}

**Example of Difference b/w Function vs Method**
type Rectangle struct {
    width, height float64
}

// Function
func area(r Rectangle) float64 {
    return r.width * r.height
}

// Method
func (r Rectangle) Area() float64 {
    return r.width * r.height
}



