package main

func main() {

	//io.WriteString()function
	// var mm, dd, yy int = 5, 2, 2003
	// var str string
	// str = fmt.Sprintf("%02d-%02d-%04d", dd, mm, yy)
	// io.WriteString(os.Stdout, str) // it prints the output on console because of os.Stdout

	// var str1 float32
	// var ok int
	// fmt.Printf("\nEnter String: ") // input:-  dgddf 342.43 3445
	// fmt.Scan(&str, &str1, &ok)
	// fmt.Printf("\nResult: %s %f %d\n", str, str1, ok) //output:-  dgddf 342.429993 3445

	// var a int = -123
	// var b uint8 = 0
	// b = a
	// fmt.Printf("a = %d , b = %d\n", a, b) // cannot use a (variable of type int) as uint value in assignment
	// type conversion is not automic so it cannot be done
	// b = uint8(a)                          //input :- a = 123
	// fmt.Printf("a = %d , b = %d\n", a, b) // output:- a = 123 , b = 123
	// a = -64
	// b = uint8(math.Abs(float64(a)))
	// fmt.Printf("a = %d , b = %d\n", a, b) // output:- a = -64 , b = 64
	// c := 34.65
	// a = int(c)
	// fmt.Println(a)

	// var x int = 255
	// var r float32
	// // r = math.Sqrt(float64(x))
	// r = float32(math.Sqrt(float64(x)))
	// fmt.Printf("Square root of %d is %.2f\n", x, r) //opt:- Square root of 255 is 15.97

	// var x int = 42
	// var y float64 = float64(x)
	// var z uint = uint(y)
	// fmt.Printf("Value of x is %d and type is %T\n", x, x) // Value of x is 42 and type is int
	// fmt.Printf("Value of y is %.2f and type is %T\n", y, y) // Value of y is 42.00 and type is float64
	// fmt.Printf("Value of z is %d and type is %T\n", z, z) // Value of z is 42 and type is uint

}
