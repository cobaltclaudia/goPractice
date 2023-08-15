package main

import "fmt"

// var name string = "Claudia"

func main() {
	// int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uintptr
	//byte
	//rune
	//float32 float64
	// complex64 complex128

	//Using var
	// var name string = "Claudia"

	// still valid
	// var name = "Claudia"

	// unused var results in error
	var age int32 = 100
	var isCool = false
	var size float32 = 2.3

	//short hand
	name := "Claudia"
	// size := 1.3
	name, email := "Claudia", "fakeemail@gmail.com"

	fmt.Println(name, age)
	fmt.Printf("%T\n", name)
	fmt.Println(age)
	fmt.Println(isCool)
	fmt.Println(size)
	fmt.Println(email)

}
