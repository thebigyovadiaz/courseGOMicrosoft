package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(num int) string {
	switch {
	case num%15 == 0:
		return "FizzBuzz"
	case num%3 == 0:
		return "Fizz"
	case num%5 == 0:
		return "Buzz"
	}
	return strconv.Itoa(num)
}

func findprimes(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	if number > 1 {
		return true
	} else {
		return false
	}
}

func main() {
	/**
	* Exercises using control flow in Go
	**/
	// Write a FizzBuzz program
	/* for i := 1; i <= 100; i++ {
		result := fizzBuzz(i)
		fmt.Println("Resultado: ", result)
	} */

	// Find the primes
	/* fmt.Println("Prime numbers less than 20:")

	for number := 1; number <= 20; number++ {
		if findprimes(number) {
			fmt.Printf("%v ", number)
		}
	} */

	// Ask a number
	val := 0

	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)

		switch {
		case val < 0:
			panic("You entered a negative number!")
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		default:
			fmt.Println("You entered:", val)
		}
	}
}
