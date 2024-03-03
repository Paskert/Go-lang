package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	//Adding two numbers static
	var n1, n2, sum int
	n1 = 5
	n2 = 10
	sum = n1 + n2
	fmt.Println("Sum of", n1, "and", n2, "is: ", sum)

	//Adding two numbers
	a1 := 10
	a2 := 20
	suum := a1 + a2
	fmt.Println(a1, "+", a2, "=", suum)

	//Affing and assagining
	var b1, b2 int
	b1, b2 = 15, 31
	sm := b1 + b2
	fmt.Println(b1, "+", b2, "=", sm)

	//Variables are declared but not defined
	// var c1, d1, int
	// var city string
	// var name string
	// fmt.Println("c1,d1,name", c1, d1, name)

	//sum
	aa := 10
	bb := 30
	sum = aa + bb

	// subtraction
	sub := bb - aa

	// Multiplication
	mul := aa * bb

	// Division
	div := bb / aa

	//Print
	fmt.Println("Sum is", sum)
	fmt.Println("Substraction is", sub)
	fmt.Println("Multiplication is", mul)
	fmt.Println("Division is", div)
}
