package main

import (
	"fmt"
	"strings"
)

func main() {

	// year := 0
	// fmt.Print("Enter the number: ")
	// fmt.Scanf("%d", &year)

	// if year == 0 {
	// 	fmt.Println("teacher says 0 is not a year")
	// 	fmt.Printf("\ntry again\n")

	// 	fmt.Print("Enter the number: ")
	// 	fmt.Scan(&year)

	// 	if year == 0 {
	// 		fmt.Println("")
	// 		fmt.Println("Yo a dumbass")
	// 		fmt.Println("Its Game Over for ya")
	// 		fmt.Println("")
	// 		// fmt.Print("Enter the number: ")
	// 		// fmt.Scanf("%d", &year)
	// 	}
	// } else {
	// 	if (year%4 == 0 && year%100 != 0) || (year%4 == 0 && year%100 == 0 && year%400 == 0) {
	// 		fmt.Println("Entered year is leap year")
	// 	} else {
	// 		fmt.Println("Entered year is not leap year")
	// 	}
	// }

	// x := 8
	// y := 10

	// if x != y {
	// 	if x < y {
	// 		fmt.Println("x is less then y")
	// 	} else if x > y {
	// 		fmt.Println("x is greater then y")
	// 	}
	// } else {
	// 	fmt.Println("x is equal to y")
	// }

	// var time float32
	// fmt.Println("enter the time: ")
	// fmt.Scan(&time)
	// if(time > 0 ){
	// 	if time <= 10 {
	// 		fmt.Println("good morning")
	// 	} else if time <= 20 {
	// 		fmt.Println("good day")
	// 	} else {
	// 		fmt.Println("good evening")
	// 	}
	// }

	// thisMonth := 5
	// switch thisMonth {
	// case 1:
	// 	fmt.Println("Jan")

	// case 2:
	// 	fmt.Println("Feb")

	// case 3:
	// 	fmt.Println("Mar")

	// case 4:
	// 	fmt.Println("April")

	// case 5:
	// 	fmt.Println("May")

	// case 6:
	// 	fmt.Println("June")

	// default:
	// 	fmt.Println(" nope ")
	// }

	// today := time.Now()
	// switch today.Day() {
	// case 5:
	// 	fmt.Println("clean the house")
	// case 10:
	// 	fmt.Println("Buy some wine")
	// case 15:
	// 	fmt.Println("Visit a doctor")
	// case 20:
	// 	fmt.Println("buy some food")
	// case 31:
	// 	fmt.Println("Party tonight")
	// default:
	// 	fmt.Println("kuch toh kar lo")
	// }

	// var month int = 12

	// month := time.Now().Month()
	// switch month {
	// case 1, 3, 5, 7, 8, 10, 12:
	// 	fmt.Println("31 days")
	// case 4, 6, 9, 11:
	// 	fmt.Println("30 days")
	// case 2:
	// 	fmt.Println("28 or 29 days")
	// }

	// Fallthrough
	// x := 1
	// switch x {
	// case 1:
	// 	fmt.Print("1") // output: 135
	// 	fallthrough
	// case 3:
	// 	fmt.Print("3") // output: 35
	// 	fallthrough
	// case 5:
	// 	fmt.Print("5") // output: 5
	// // 	fallthrough
	// // default:
	// // 	fmt.Println("Not in above cases")
	// }

	// var x interface{}
	// // var x interface{} = "RKNEC" // try it
	// switch i := x.(type) {
	// case nil:
	// 	fmt.Printf("type of x: %T", i)
	// case int:
	// 	fmt.Println("x is int")
	// case float64:
	// 	fmt.Println("x is float64")
	// case func(int) float64:
	// 	fmt.Println("x is func(int)")
	// case bool, string:
	// 	fmt.Println("x is bool / string")
	// default:
	// 	fmt.Println("don't know")
	// }

	//goto
	// i := 0
	// loop:
	// fmt.Println(i)
	// i++
	// if i < 5 {
	// 	goto loop
	// }
	// fmt.Println("Loop ends here")

	// 	i := 0
	// pattern:
	// 	fmt.Print("/\\/\\/\\/\\")
	// 	i++
	// 	if i <= 10000 {
	// 		goto pattern
	// 	}

	var str string = "Hello World"
	var substr string = "wor"

	if strings.Contains(str, substr) {
		fmt.Printf("String (%s) contains sub-string (%s)\n", str, substr)
	} else {
		fmt.Printf("String (%s) does not contains sub-string (%s)\n", str, substr)
	}
	fmt.Println(strings.ToUpper(str))    // opt: HELLO WORLD
	fmt.Println(strings.ToLower(str))    // opt: hello world
	fmt.Println(strings.Index(str, "W")) // opt: 6

}
