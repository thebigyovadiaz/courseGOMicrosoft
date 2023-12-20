package main

import (
	"fmt"
	"strconv"
)

func main() {
	/**
	* TODO: Integer Numbers
	*
	**/
	// var integer8 int8 = 127
	// var integer16 int16 = 32767
	// var integer32 int32 = 2147483647
	// var integer64 int64 = 9223372036854775807
	// fmt.Println(integer8, integer16, integer32, integer64)

	// // Challenge 1
	// var integer32 int = 2147483648
	// fmt.Println(integer32)

	// // Challenge 2
	// var integer uint = -10
	// fmt.Println(integer)

	/**
	* TODO: Float numbers
	*
	**/
	// var float32 float32 = 2147483647
	// var float64 float64 = 9223372036854775807
	// fmt.Println(float32, float64)
	// fmt.Println(math.MaxFloat32, math.MaxFloat64)

	/**
	* TODO: Booleans
	*
	**/
	// var featureFlag bool = true
	// fmt.Println(featureFlag)
	// fmt.Println(!featureFlag)

	/**
	* TODO: Strings
	*
	**/
	// fullName := "John Doe \t(alias \"\\Foo\")\n"
	// fmt.Println(fullName)

	/**
	* TODO: Type Conversions
	*
	**/
	// Integer casting
	var integer16 int16 = 127
	var integer32 int32 = 32767
	fmt.Println(int32(integer16) + integer32)

	// String casting
	i, _ := strconv.Atoi("-41")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)
}
