package main

import "fmt"

// Distintos tipos de declarar variables
/* var firstName, lastName string
var age int */

/* var (
	firstName, lastName string
	age int
) */

/* var (
    firstName string = "Yova"
    lastName  string = "Diaz"
    age       int    = 37
) */

/* var (
	firstName, lastName, age = "Yova", "Diaz", 37
) */

// Global variables declared but these can modified
var (
	firstName = "Yova"
	lastName  = "Diaz"
	age       = 37
)

// Constants declaration
const HTTPStatusOK = 200

const (
	StatusOK              = 0
	StatusConnectionReset = 1
	StatusOtherError      = 2
)

func main() {
	firstName, lastName := "Yovanny", "Diaz"
	age := 30
	fmt.Println(firstName, lastName, age)
	fmt.Println(HTTPStatusOK)
	fmt.Println(StatusConnectionReset)
}
