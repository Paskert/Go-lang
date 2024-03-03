package main

import (
	"fmt"
)

func main() {
	// 	the hex form (starts with a "0x" or "0X") is used to represent numbers in base 16.
	// 	the octal  form starts with a "0o" or "0O".
	// 	the binary form (starts with "0b" or "0B") represents numbers in base 2.
	// a := 10
	// fmt.Println(15 == 0o17)
	// fmt.Println(15 == 0xF)
	// a++
	// fmt.Println(a)
	// a--
	// fmt.Println(a)
	// b := 15
	// fmt.Printf("Binary of %d is: %b\n", b, b)
	// fmt.Printf("hexadecimal of %d is: %x\n", b, b)
	// fmt.Printf("Octal of %d is: %o\n", b, b)

	// 1.23e2 == 123
	// 123e-2 == 1.23
	//

	//v format value in default format
	//d format decimal integers
	//g format floating point numbers in general format
	//b format floating point numbers in base 2
	//o format octal numbers
	//t format 1 for true and 0 for false(boolean)
	//s format string values
	// fmt.Printf("Expression 15 == 0xf is %t", 15 == 0xf)

	// fname := "aditya"
	// lname := "kumar"
	// fmt.Printf("My name is", fname, "and My last name is", lname)
	// fmt.Printf("%s\n", fname)
	// fmt.Printf("%s\n", lname)
	// i := 10
	// j := "rajesh"
	// fmt.Printf("%s %s", i, j)
	// fmt.Print(i, j)

	// i := 15.5
	// var txt = "hello world!"
	// fmt.Println("")
	// fmt.Printf("%v\n", i)   //15.5
	// fmt.Printf("%#v\n", i)  //15.5
	// fmt.Printf("%v%%\n", i) //15.5%
	// fmt.Printf("%T\n", i)   //float64
	// fmt.Printf("%v\n", txt) // hello world!
	// fmt.Printf("%s", txt)
	// fmt.Printf("%#v\n", txt) // "hello world!"
	// fmt.Printf("%T\n", txt)  //string

	// fmt.Printf("%04d\n", i)

	// i := 15
	// fmt.Println("")
	// fmt.Println("")

	// fmt.Printf("%b\n", i)   //1111
	// fmt.Printf("%d\n", i)   //15
	// fmt.Printf("+%d\n", i)  //+15
	// fmt.Printf("%o\n", i)   //17
	// fmt.Printf("%O\n", i)   //0o17
	// fmt.Printf("%x\n", i)   //f
	// fmt.Printf("%X\n", i)   //F
	// fmt.Printf("%#x\n", i)  //0xf
	// fmt.Printf("%4d\n", i)  //  15 with 2 spaces in front of 15
	// fmt.Printf("%-4d\n", i) //15   whith 2 spaces after 15

	// fmt.Println("")
	// fmt.Println("")

	// var txt = "hello"

	// fmt.Printf("%s\n", txt)
	// fmt.Printf("%q\n", txt)
	// fmt.Printf("%8s\n", txt)
	// fmt.Printf("%-8s\n", txt)
	// fmt.Printf("%x\n", txt)
	// fmt.Printf("% x\n", txt)

	// HW
	// take the input from the keybord result 5 subject marks
	// output should be formated with total marks and percentage

	// w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight)
	var mat, sci, sst, eng, hin, total, percentage int
	fmt.Print("Enter the marks for Maths: ")
	// fmt.Scan(&mat)
	fmt.Scan(&mat)
	// fmt.Printf("%d", mat)
	// fmt.Scan(&mat)
	fmt.Print("Enter the marks for Science: ")
	fmt.Scan(&sci)
	// fmt.Scan(&sci)
	fmt.Print("Enter the marks for Social Science: ")
	fmt.Scan(&sst)
	// fmt.Scan(&sst)
	fmt.Print("Enter the marks for English: ")
	fmt.Scan(&eng)
	// fmt.Scan(&eng)
	fmt.Print("Enter the marks for Hindi: ")
	fmt.Scan(&hin)
	// fmt.Scan(&hin)

	// fmt.Println(mat, sci, sst, eng, hin)
	// w.Flush()

	total = mat + sci + sst + eng + hin
	percentage = total / 5

	// fmt.Println(mat, sci, sst, eng, hin, total, percentage)
	fmt.Print("\n\nName: Ajay devgan \t\t Brach: Cyber Security\n\n")
	fmt.Printf("------------------------------------------------------------------------------------------------------------------\n")
	fmt.Printf("|\tMaths\t|\tScience\t|\tS.Science |\tEnglish |      Hindi\t|  Total marks  |   Percentage   |\n")
	fmt.Printf("------------------------------------------------------------------------------------------------------------------\n")
	fmt.Printf("|\t%d\t|\t%d\t|\t%d\t  |\t%d\t|\t%d\t|\t%d\t|\t%d%%\t |\n", mat, sci, sst, eng, hin, total, percentage)
	fmt.Printf("------------------------------------------------------------------------------------------------------------------\n")

	if (mat >= 40) && (sci >= 40) && (sst >= 40) && (eng >= 40) && (hin >= 40) {
		fmt.Println("\nResult: Passed\n")
	} else {
		fmt.Println("\nResult: Failed\n")
	}
}
