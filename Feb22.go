package main

import "fmt"

//Functions in GO

//func functionName(){
//code
//}
var suum int

func greet() {
	fmt.Println("Good Morning")
}

func addNumbers(n1 int, n2 int) {
	// a := 12
	// b := 8
	suum = n1 + n2
	// fmt.Println("Sum:", sum)
	// fmt.Println("Sum:", n1+n2)
	// return suum
	// fmt.Println("After return statement") // missing return
}

func calc(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func getSquare(num int) {
	square := num * num
	fmt.Printf("Square of %d is %d\n", num, square)

}

func countDown(n int) {

	if n >= -10 {
		fmt.Println(n)
		countDown(n - 1)
	}
}

func fibonacci(n int, f int, s int) {

	if n > 0 {
		fmt.Print(s, " ")
		temp := f + s
		f = s
		s = temp

		n--
		fibonacci(n, f, s)

	}
}

func main() {
	// greet()
	// fmt.Println("hello")
	// var x, y int
	// fmt.Println("Enter x and y: ")
	// fmt.Scan(&x, &y)
	// var suum int
	// suum := addNumbers(x, y)
	// fmt.Println("Sum:", suum)

	// sum, diff := calc(21, 13)
	// fmt.Println("Sum:", sum, "Difference:", diff)

	//call function 3 times
	// getSquare(3)
	// getSquare(5)
	// getSquare(7)
	// addNumbers(23, 32)
	// fmt.Println("Sum is:", suum)
	// countDown(3)

	var n int
	fmt.Print("Enter the no of fibonacci items : ")
	fmt.Scan(&n)
	f, s := 0, 1
	fibonacci(n, f, s)

}
