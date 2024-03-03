package main

import (
	"fmt"
)

func main() {
	// integer := 23
	// fmt.Printf("%T %T\n", integer, &integer) //output:- int *int

	// fat := 1234.1234565
	// fmt.Printf("%5.2f\n", fat) //output:- 1234.12

	// fmt.Printf("%*s\n", 40, "text") //Output:-                                    text
	// fmt.Printf("%040s\n", "text")     //Output:-000000000000000000000000000000000000text
	// text := "Rajesh"
	// fmt.Printf("%*s\n", 40, text) //Output:-                                   Rajesh
	// fmt.Printf("%040s\n", text)     //Output:-0000000000000000000000000000000000Rajesh

	// var val byte = 'A'
	// ch := 'A'
	// bch := "B"

	// fmt.Printf("ch ' is %T and bch \" is %T and val byte is %T ", ch, bch, val)
	// output:- ch ' is int32 and bch " is string and val byte is uint8

	// var valf float64 = -19.25
	// var vald int = -19
	// fmt.Printf("Absolute value of %f is %f", valf, math.Abs(valf))
	// //output:- Absolute value of -19.250000 is 19.250000
	// fmt.Printf("Absolute value of %d is %d", vald, math.Abs(vald))
	// //output:- cannot use vald (variable of type int) as float64 value in argument to math.Abs

	// var num1, num2, large float64 = 10.25, 20.14, 0
	// // var nu1, nu2 int = 10, 20
	// large = math.Max(num1, num2)
	// fmt.Printf("Largest number is %f\n", large) //output:- Largest number is 20.140000
	// // large = math.Max(nu1, nu2)
	// // fmt.Printf("Largest number is %f", large)
	// //output:- cannot use nu1 (variable of type int) as float64 value in argument to math.Max
	// large = math.Max(10, 20)
	// fmt.Printf("Largest number is %f\n", large) //opt: Largest number is 20.000000
	// large = math.Min(num1, num2)
	// fmt.Printf("Largest number is %f\n", large) //opt: Largest number is 10.250000

	// var res float64 = 0
	// res = math.Ceil(num1)
	// fmt.Printf("Result is : %f\n", res) //opt:Result is : 11.000000
	// res = math.Ceil(num2)
	// fmt.Printf("Result is : %f\n", res) //opt:Result is : 21.000000
	// var six = 12342323.2323
	// fmt.Printf("%20.2f", six) //opt:         12342323.23

	//Printing without package
	// println("Using println insted of fmt.Println")
	// print("Using print insted of fmt.Print")

	// var x = 10
	// x += 5
	// println(x)
	var in int
	fmt.Printf("type input: ")
	fmt.Scan(&in)
	// in = in % 2
	fmt.Printf("\nint is %d and modulo is %d", in, in%2)

}
