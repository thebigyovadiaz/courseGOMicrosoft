package main

import (
	"fmt"
)

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		panic("highlow() low greater than high")
	}
	defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")

	highlow(high, low+1)
}

func main() {
	/**
	* Using Defer function
	**/
	/* newFile, error := os.Create("learnGO.txt")
	if error != nil {
		fmt.Println("Error: Could not create file.")
		return
	}
	defer newFile.Close()

	if _, error := io.WriteString(newFile, "Learning GO with Microsoft!"); error != nil {
		fmt.Println("Error: Could not write to file.")
		return
	}

	newFile.Sync() */

	/**
	* Using Panic and Recover function
	**/
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("main(): recover", handler)
		}
	}()

	highlow(2, 0)
	fmt.Println("Program finished successfully!")

	/**
	* Exercises using control flow in Go
	**/
	// Write a FizzBuzz program

}
