package main

import (
	"fmt"
)

const PI = 3.1415290

// there are two types of constant
// Typed Constant
const A int = 10

// Untyped
const B = 1 //type is infered

const (
	a int = 1
	b     = 2
	c     = "Hi !!!"
)

func main() {

	// PI = PI + 1 cant do this because it is a constant
	fmt.Println(PI, A, B, a, b, c)

	const Length int = 10
	const Width int = 5
	var area int
	area = Length * Width
	fmt.Printf("area of rectangle: %d", area)

	//
	// Type and Descryption
	var g uint8 = 255
	fmt.Printf("\n%d\n", g)
	fmt.Println(g)

	var student1 string = "James"
	var student2 = "Bond"
	x := 2

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Printf("Type of the variable  : %T %T\n", x, g)

	var aaa string
	var bbb int
	var c bool

	fmt.Println(aaa, bbb, c)

	fmt.Printf("a is of type %T \n", aaa)
	fmt.Printf("b is of type %T \n", bbb)
	fmt.Printf("c is of type %T \n", c)
	// Slicing Arraysfmt.Printf("a is of type %T \n", a)

	// fmt.Println(aaa, "\t", reflect.TypeOf(aaa))
	// fmt.Println(bbb, "\t", reflect.TypeOf(bbb))
	// fmt.Println(c, "\t", reflect.TypeOf(c))
}
