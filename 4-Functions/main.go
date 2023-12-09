package main

import (
	"fmt"

	"github.com/thebigyovadiaz/calculator"
)

/* func sum(numb1 string, numb2 string) (result int) {
	number1, _ := strconv.Atoi(numb1)
	number2, _ := strconv.Atoi(numb2)
	result = number1 + number2
	return
} */

/* func calculator(numb1 string, numb2 string) (sum int, mult int, div int, subs int) {
	number1, _ := strconv.Atoi(numb1)
	number2, _ := strconv.Atoi(numb2)
	sum = number1 + number2
	mult = number1 * number2
	div = number1 / number2
	subs = number1 - number2
	return
} */

func updateName(name *string) *string {
	*name = "Pepito"
	return name
}

func main() {
	/***
	* Main Functions
	**/
	// number1, _ := strconv.Atoi(os.Args[1])
	// number2, _ := strconv.Atoi(os.Args[2])
	// fmt.Println(number1 + number2)

	/***
	* Custom Functions
	**/
	//sum := sum(os.Args[1], os.Args[2]) // Sum function
	// sum, mult, div, subs := calculator(os.Args[1], os.Args[2]) // Calculator function
	// fmt.Println(sum, subs)
	// fmt.Println(mult, div)

	/***
	* Change Function Parameter
	**/
	firstName := "Peter"
	updateName(&firstName)
	fmt.Println(firstName)

	/***
	* Using package locally
	**/
	fmt.Println(calculator.Version)
	sum := calculator.Sum(5, 20)
	subs := calculator.Subs(25, 20)
	mul := calculator.Mul(5, 20)
	div := calculator.Div(500, 20)
	fmt.Println("Sum: ", sum)
	fmt.Println("Subs: ", subs)
	fmt.Println("Mul: ", mul)
	fmt.Println("Div: ", div)
}
