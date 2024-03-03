package main

import "fmt"

// here the function is calculate with return type(annonyoms function which returns an int)
func calculate() func() int {
	num := 1
	return func() int {
		num += 2
		return num
	}
}

func greet() func() string {
	name := "John"
	return func() string {
		name = "Hi " + name
		return name
	}
}

//here the displayNumber() is returning [func() int] and not an int
func displayNumber() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {

	// var odd int // cannot use calculate() (value of type func() int) as int value in assignment
	// since we are assingning calculate function to odd now odd is a function
	odd := calculate()
	fmt.Println(odd()) //3
	fmt.Println(odd()) //5
	fmt.Println(odd()) //7
	// the values keep increasing bcoz even after the calculate is executed odd is a func()
	// if we wanted to do this in c then we have to define static int = 1 which will declare only once and
	// every time we call it will give us next number and will not get formated until the whole code is executed

	odd2 := calculate()
	fmt.Println(odd2()) // here data is isolated so it will not affect each other

	message := greet()
	fmt.Println(message()) //
	fmt.Println(message())

	num1 := displayNumber()
	fmt.Println(num1())
	fmt.Println(num1())
	fmt.Println(num1())

	num2 := displayNumber()
	fmt.Println(num2())
	fmt.Println(num2())
	fmt.Println(displayNumber())
	fmt.Println(displayNumber())
	fmt.Println(displayNumber())
	fmt.Println(displayNumber()())
	fmt.Println(displayNumber()())
	fmt.Println(displayNumber()())

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
	//7
	//6 5
	//5 4 , 4 3
	//4 3 , 3 2, 3 2, 2 1
	//3 2 2 1 ,2 1, ., 3 2 .,. .

}
