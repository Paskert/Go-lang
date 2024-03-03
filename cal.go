// PLACE THIS CODE WITH NAME "calculator" IN PACKAGES OF GO 
// package calculator

// func Add(a, b int) int {
// 	return a + b
// }
// func Subtract(n1, n2 int) int {
// 	return a - b
// }

package main

import {
	"fmt"
	"Packages/calculator"
}

func main(){
	number1 := 9
	number2 := 5

	// use the add function of calculator package
	fmt.Println(calculator.Add(number1, number2))   // output: 14

	// use the subtract function of calculator package
	fmt.Println(calculator.Subtract(number1,number2))
}