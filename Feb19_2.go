package main

func main() {

	// Operators

	// a := 5
	// b := 3
	// c := a & b  //Bitwise AND
	// d := a | b  //Bitwise OR
	// e := a ^ b  //Bitwise XOR ('^' is called carrot || carret )
	// f := a >> b //Bitwise  Right shift
	// g := a << b //Bitwise  Left shift
	// fmt.Printf("%d\n", c) // 1
	// fmt.Printf("%d\n", d) // 7
	// fmt.Printf("%d\n", e) // 6
	// fmt.Printf("%d\n", f) // 0
	// fmt.Printf("%d\n", g) // 40
	// a = 10
	// fmt.Printf("%t\n", a > 5) // true
	// fmt.Printf("%t\n", a == 10) // true
	// fmt.Printf("%t\n", a < 5) // false
	// // fmt.Printf("%t\n", a && 5)
	// // fmt.Printf("%t\n", a || 5)

	//import "unsafe"
	// var num1 int = 10
	// var num2 byte = 20
	// var num3 float32 = 30
	// var num4 float64 = 30
	// var name string = "Rajesh"
	// fmt.Printf("\nSize of Num1: %d", unsafe.Sizeof(num1))
	// fmt.Printf("\nSize of Num2: %d", unsafe.Sizeof(num2))
	// fmt.Printf("\nSize of Num3: %d", unsafe.Sizeof(num3))
	// fmt.Printf("\nSize of Num4: %d", unsafe.Sizeof(num4))
	// fmt.Printf("\nSize of Name: %d", unsafe.Sizeof(name))

	//import "reflect"
	// a := 10
	// b := 10.2
	// c := "Hello"
	// d := true
	// e := []string{"India", "USA", "UK"}
	// fmt.Println("Type of a = ", reflect.TypeOf(a)) //Type of a =  int
	// fmt.Println("Type of b = ", reflect.TypeOf(b)) //Type of b =  float64
	// fmt.Println("Type of c = ", reflect.TypeOf(c)) //Type of c =  string
	// fmt.Println("Type of d = ", reflect.TypeOf(d)) //Type of d =  bool
	// fmt.Println("Type of e = ", reflect.TypeOf(e)) //Type of e =  []string
	// fmt.Println("Type of a = ", reflect.TypeOf(a).Kind()) //Type of a =  int
	// fmt.Println("Type of b = ", reflect.TypeOf(b).Kind()) //Type of b =  float64
	// fmt.Println("Type of c = ", reflect.TypeOf(c).Kind()) //Type of c =  string
	// fmt.Println("Type of d = ", reflect.TypeOf(d).Kind()) //Type of d =  bool
	// fmt.Println("Type of e = ", reflect.TypeOf(e).Kind()) //Type of e =  slice

	//Keyword defer
	// defer fmt.Println("hello")
	// fmt.Println("Good morning")
	// defer fmt.Println("hii")
	// // Output:-
	// // Good morning
	// // hii
	// // hello

	// var radius, perameter float32
	// var area float32
	// fmt.Printf("Enter Radius: ")
	// fmt.Scanf("%f", &radius)
	// const PI = 3.14159
	// area = PI * radius * radius
	// perameter = 2 * PI * PI
	// fmt.Printf("Area of Circle: %.2f", area)
	// fmt.Printf("Parameter of Circle: %.2f", perameter)

	//Fahrenheit to Celcius
	// var ftemp, ctemp float64
	// fmt.Printf("Enter temperature in Fahrenheit: ")
	// fmt.Scanf("%f", &ftemp)
	// ctemp = (ftemp - 32) / 1.8
	// fmt.Printf("Temperature in Celsius: %.2f", ctemp)

	//Celcius to Fahrenheit
	// fmt.Printf("Enter temperature in Celcius: ")
	// fmt.Scanf("%f", &ctemp)
	// ftemp = (ctemp * 1.8) + 32
	// fmt.Printf("Temperature in Fahrenheit: %.2f", ftemp)

}
