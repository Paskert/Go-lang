package main

import "fmt"

func sum(n int) int {
	if n <= 0 {
		return 0
	} else {
		return n + sum(n-1)
	}
}

func factorial(n int) int {
	if n < 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

//this is annonymus function
var greetme = func() string {
	// fmt.Println("SUPERRRR!!!")
	return "SUPERRRR!!!"
}

func findSquare(n int) int {
	return n * n
}

func main() {
	// fmt.Println("type the number to sum it:")
	// var num int
	// fmt.Scan(&num)
	// if num <= 0 {
	// 	fmt.Println("Invalid input please try again")
	// } else {
	// 	result := sum(num)
	// 	fmt.Println("Result is: ", result)
	// }

	//anonymous function
	var greet = func() string {
		// fmt.Println("Hello,how are you")
		return "And then there were none"
	}
	greet()
	// greetme()
	franky := greetme()
	OnePiece := greet()
	fmt.Println(franky, OnePiece)
	var twoSum = func(n1, n2 int) {
		result := n1 + n2
		fmt.Println("Sum is:", result)
	}
	twoSum(13, 56)

	var threeSum = func(n1, n2 int) int {
		res := n1 + n2
		// fmt.Println("Sum is:", result)
		return res
	}
	fmt.Println(threeSum(23, 32))

	rest := findSquare(threeSum(6, 9))
	fmt.Println("Result is:", rest)
	// fmt.Println(greeting())
}
