# Ravichandra
<<<<<<< HEAD
This is go lang program
# Theory of Function and Methods

# Functions
It is a standalone block of code that perform a spefic task. It will take zero or more arguments as input and it will perform operations can optionally return one or more values. It helps to reuse code and keep ir organized.
It will declare with keyword func (function declaration - Interms of code maintainbility, reusability and flexibility)

**<pre>syntax</pre>**
func functionName(parameters) returnType{
      ------code------
      }

**<pre>Example</pre>**
=======
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
>>>>>>> c0fc344f4e65b816f3b8f2e0da0836af6b381904
func add(a int, b int) int {
    retuen a+b
    }
In real-Time function used for utility tasks-----calculation, DB query, API call.


<<<<<<< HEAD
# Method
Method is a function with a receiver like it is associated with a type called struct, alias and soon.
It is useful for behavior and encapsulation.

**<pre>syntax</pre>**
func(receiverName ReceiverType) MethodName(parameters) returnType.

**<pre>Example</pre>**
=======
**Method**
Method is a function with a receiver like it is associated with a type called struct, alias and soon.
It is useful for behavior and encapsulation.
**Syntax**
func(receiverName ReceiverType) MethodName(parameters) returnType.
**Example**
>>>>>>> c0fc344f4e65b816f3b8f2e0da0836af6b381904
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

<<<<<<< HEAD
# Example of Difference b/w Function vs Method
=======
**Example of Difference b/w Function vs Method**
>>>>>>> c0fc344f4e65b816f3b8f2e0da0836af6b381904
type Rectangle struct {
    width, height float64
}

// Function
<<<<<<< HEAD
It has no receiver
we can call directly like
add(3,5)
=======
>>>>>>> c0fc344f4e65b816f3b8f2e0da0836af6b381904
func area(r Rectangle) float64 {
    return r.width * r.height
}

// Method
<<<<<<< HEAD
It has receiver like c *city
it can called on an object c.PrintHello()
func (r Rectangle) Area() float64 {
    return r.width * r.height
}
=======
func (r Rectangle) Area() float64 {
    return r.width * r.height
}



>>>>>>> c0fc344f4e65b816f3b8f2e0da0836af6b381904
