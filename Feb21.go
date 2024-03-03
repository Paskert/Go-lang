package main

import (
	"fmt"
)

// func add(x,y){
// 	z:=x + Y
// 	fmt.Println("Absolute is :",z)
// }

func main() {
	// fmt.Println("Fback")
	// var num1, num2 int = 10, 5
	// var avg float64
	// avg = (float64(num1) + float64(num2)) / 2
	// int(avg) = (num1 + num2) / 2
	// fmt.Printf("Average of %d and %d is %.2f\n", num1, num2, (num1+num2)/2)

	// IF Statement
	// Syntax
	// if condition {
	// 	// code
	// }
	// if 20 > 10 {
	// 	fmt.Println(" > ")
	// }
	// if 20 != 10 {
	// 	fmt.Println(" != ")
	// }
	// if !(20 < 10) {
	// 	fmt.Println(" !(20 < 10) ")
	// }
	// if true {
	// 	fmt.Println(" true ")
	// }

	// x := 8
	// if y := 10; x < y {
	// 	fmt.Println(" y := 10; x < y ")
	// }

	// if _ , err := 2312{
	// 	if( err!= nil ){
	// 		fmt.Println()
	// 	}
	// }
	// var name string
	// var age int
	// if _, err := fmt.Scan(&name, &age); err != nil {
	// 	fmt.Println("err =", err)
	// 	os.Exit(1)
	// }
	// // The Scan function returns two values: the first is a value of type (void, error), represented in the
	// // code as _, which is ignored. The second is an error value, which is stored in the variable err.
	// var abc string
	// if abc, err := fmt.Scan(&name, &age); err != nil {
	// 	fmt.Println("err =", err)
	// 	fmt.Println("abc =", abc)
	// 	os.Exit(1)
	// }
	// // Pratham
	// // 21
	// // abc
	// // sadf
	// // err = expected integer
	// // abc = 1
	// // exit status
	// fmt.Printf("Your name is: %s\n", name)
	// fmt.Printf("Your age is: %d\n", age)
	// fmt.Println("abc =", abc)

	// var n1, n2, n3 int //= 33, 25, 15
	// fmt.Scan(&n1) // scan ln takes input in one line so if we input 11 22 33 then incluing all it will take 11 and print 11 as the biggest
	// fmt.Scan(&n2) // Scan can take 11 22 33 and give 33 as the biggest
	// fmt.Scan(&n3)

	// if n1 > n2 {
	// 	if n1 > n3 {
	// 		fmt.Println("n1", n1)
	// 	}
	// 	if n1 < n3 {
	// 		fmt.Println("n3", n3)
	// 	}
	// }
	// if n1 < n2 {
	// 	if n2 > n3 {
	// 		fmt.Println("n2", n2)
	// 	}
	// 	if n2 < n3 {
	// 		fmt.Println("n3", n3)
	// 	}
	// }

	//Time function
	// startTime := time.Now()
	// here comes the code
	// endTime := time.Now()
	// extime := endTime.Sub(startTime)
	// fmt.Printf("Function executed in %v\n", extime)

	x, y := 18, 10
	if x < y {
		fmt.Println("X is greater then y")
	} else {
		fmt.Println("Y is greater then OR equal to x")
	}

	// "}else{" is allowed
	// "}else" is also allowed if { of else is on next line
	// "else" is not allowed when } of if is on previous line //output:-  syntax error: unexpected else, expected }

	// x, y := 18, 10
	// var f int
	// fmt.Print("enter the number: ")
	// fmt.Scan(&f)
	//Odd or Even
	// if f%2 == 0 {
	// 	fmt.Println("Even")
	// } else {
	// 	fmt.Println("Odd")
	// }
	// //abs value
	// if f < 0 {
	// 	fmt.Println("Absolute value is:", f*-1)
	// } else {
	// 	fmt.Println("Absolute value is:", f)
	// }

}
