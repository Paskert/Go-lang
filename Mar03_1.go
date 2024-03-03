package main

import (
	"fmt"
)

// var a = 10 // this is allowed
// b := 12 // syntax error: non-declaration statement outside function body
func test(x int) {
	fmt.Println("hello !", x)
}

// syntax error: non-declaration statement outside function body
var test3 = func(x int) {
	fmt.Println(x)
}

var test5 = func(x int) int {
	return x * -2
}(8)

func calculate(x int, y int) (int, int) {
	return x + y, x - y
}

// assigning a function to a variable
func main() {

	// x := test
	// x(5)

	// test2 := func(x int) {
	// 	fmt.Println(x)
	// }
	// test2(5)
	// test3(5)
	// // fmt.Println(a)
	// test4 := func(x int) int {
	// 	return x * -1
	// }(8)
	// fmt.Println(test4)
	// fmt.Println(test5)

	// sum, difference := calculate(10, 20)
	// fmt.Println("sum and difference is", sum, "and", difference)

	// yyyy, mm, dd := time.Now().Date()

	// fmt.Println(dd, "/", mm, "/", yyyy)

	// currentDateTime := time.Now()

	// day := currentDateTime.Day()
	// month := currentDateTime.Month()
	// year := currentDateTime.Year()
	// hour := currentDateTime.Hour()
	// min := currentDateTime.Minute()
	// sec := currentDateTime.Second()

	// fmt.Printf("\nReport generated on Day %d/%d/%d and on Time %d:%d:%d", day, month, year, hour, min, sec)

	// fmt.Println("hello world")
	// time.Sleep(500 * time.Millisecond)
	// fmt.Println("Hello India!")
	// time.Sleep(8 * time.Second)
	// fmt.Println("hello welcome to golang")

	colors := []string{"red", "yellow", "green"}

	for index, val := range colors {
		fmt.Println(index, "-", val)
	}
	for _, val := range colors {
		fmt.Println("-", val)
	}
	for index, _ := range colors {
		fmt.Println(index, "-")
	}

	var array = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	fmt.Printf("%T \n", array) // opt:- [5]int

	arrayOfIntegers := [5]int{0: 7, 3: 9}
	fmt.Println(arrayOfIntegers)

	var arrayOfString = [...]string{"Hello", "Rajesh"}
	fmt.Println(arrayOfString)

}
