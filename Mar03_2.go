package main

import (
	"fmt"
)

func main() {

	//code to check if the date and time are same or not
	// date1 := time.Date(2020, 2, 5, 10, 5, 20, 0, time.UTC)
	// date2 := time.Date(2020, 2, 5, 10, 5, 20, 0, time.UTC)
	// date3 := time.Date(2020, 2, 7, 10, 5, 20, 0, time.UTC)

	// we use Equal fucntion to check dates
	// if date1.Equal(date2) {
	// 	fmt.Println("date1 = date2")
	// } else {
	// 	fmt.Println("date1 ≠ date2")
	// }
	// //writing "==true" is optional
	// if date1.Equal(date3) == true {
	// 	fmt.Println("date1 = date3")
	// } else {
	// 	fmt.Println("date1 ≠ date3")
	// }
	// Output:-
	// date1 = date2
	// date1 ≠ date3

	// FOR LOOP

	// For Loop working as for loop
	// for i := 1; i <= 3; i++ {
	// 	fmt.Println("for")
	// }

	// For Loop working as while loop
	// num1 := 1
	// for num1 <= 3 {
	// 	fmt.Println("while")
	// 	num1++
	// }

	// For Loop working as (do while / infinite) loop
	// number := 1
	// for {
	// 	fmt.Printf("do while %d < 5\n", number)
	// 	number++
	// 	if number > 5 {
	// 		break
	// 	}
	// }

	// code to print table of numbers using for loops
	// var table int
	// fmt.Println("enter the no of table you want:")
	// fmt.Scan(&table)
	// i := 1
	// for i <= 10 {
	// 	fmt.Print(table, " * ", i, " = ", table*i, "\n")
	// 	i++
	// }

	str := "Hello World!"

	//Printing character of a string using for loop
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}

	//code for printing ASCII value of characters of string
	fmt.Println("\nAscii value of stirng:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c - %d\n", str[i], str[i])
	}

	//Slice
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
	//Array
	var array = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	fmt.Printf("%T \n", array) // opt:- [5]int

	arrayOfIntegers := [5]int{0: 7, 3: 9}
	fmt.Println(arrayOfIntegers) // output:- [7 0 0 9 0]

	var arrayOfString = [...]string{"Hello", "Rajesh"}
	fmt.Println(arrayOfString) // output:- [Hello Rajesh]

	//create an array
	var arr = [...]int{1, 2, 3, 4, 5, 6, 7}
	//find the length of array using len()
	length := len(arr)
	fmt.Println("Length of Array is :", length)
	// opt:- Length of Array is : 7

	// loop through the array
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

}
