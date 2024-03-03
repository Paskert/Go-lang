package main

import (
	"fmt"
)

// aa:=10
// var bb int = 20

//var can be declared anywhere but ":=" cannout be used outside the function i.e here
// aa:=10 //syntax error: non-declaration statement outside function body

func main() {
	var a, b, c, d = 1, 2, 3, true // after the variable d there is no ':' so this is infered declaration
	e, f := 29, "DOOM"             // this is also infered declaration without var but ':=' operator

	var (
		g int
		h int    = 1
		i string = "hello"
	)
	var name1, name2 string

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Println(g, "\n", h, "\n", i)
	fmt.Printf("%d\n", g)
	fmt.Printf("%d\n", h)
	fmt.Printf("%s\n", i)
	fmt.Printf("%d\n%d\n%s\n", g, h, i)

	//input = James Moriarty
	fmt.Println("First name")
	// fmt.Scanf("%s", &name1) // it takes the string until it encounters a space coz it is implies two words
	fmt.Println("Last name")
	//fmt.Scanln(&name2)// here i have written two scan statements bcoz after the Print "last name" the compiler ignores this statement
	// fmt.Scanln(&name2)
	fmt.Println("Hello", name1)
	fmt.Println("hello", name2)

	var j, k int = 10, 20
	fmt.Println(j, k)
	fmt.Println(j + k)

	var l, m string = "James", "Moriarty"
	fmt.Println("Hello", l, m)

	// var fname, Age, Brach, cname string
	// fmt.Print("Name: ")
	// fmt.Scan(&fname)
	// fmt.Printf("Age: ")
	// fmt.Scan(&Age)
	// fmt.Printf("Branch: ")
	// fmt.Scan(&Brach)
	// fmt.Printf("College Name: ")
	// fmt.Scan(&cname)

	var n string = "Rajesh"
	// fmt.Println("")
	fmt.Println("Value of a is", "'", n, "'")
	fmt.Print("Value of a is ", "'", n, "'")
	// fmt.Println("")
	const name, dept = "GeeksforGeeks", "CS"

	fmt.Printf("\n%s is a %s portal.\n", name, dept)
	fmt.Printf("\nThe string is %s \n", name)
	var num1 int = 21
	fmt.Printf("\nThe decimal value is %d \n", num1)
	var num2 float32 = 7.283
	fmt.Printf("\nThe float is %f \n", num2)
	var num3 int = 14
	fmt.Printf("\nThe binary value of num3 is %b \n", num3)
	var num4 = 4 + 11
	fmt.Printf("\nThe scientific notation of num4: %e \n", num4)

	var o int = 10
	var p float32 = 10.6
	var nam string = "Rajesh"
	fmt.Printf("%d %f %s", o, p, nam)
}
