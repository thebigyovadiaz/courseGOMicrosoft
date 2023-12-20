package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func giveMeANumber(num int) int {
	return num
}

func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
		region, continent = "India", "Asia"
	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"
	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}

	return region, continent
}

func main() {
	/***
	* Conditional IF-ELSE
	**/

	x := 27

	// if x%2 == 0 {
	// 	fmt.Println(x, "is even")
	// } else {
	// 	fmt.Println(x, "is not even")
	// }

	// digit, _ := strconv.Atoi(os.Args[1])

	if num := giveMeANumber(x); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	/***
	* Switch Case
	**/

	sec := time.Now().Unix()
	rand.NewSource(sec)
	i := rand.Int31n(10)

	switch i {
	case 0:
		fmt.Print("zero...")
	case 1:
		fmt.Print("one...")
	case 2:
		fmt.Print("two...")
	default:
		fmt.Print("no match...")
	}

	fmt.Println("ok")

	// Using function location
	region, continent := location("kochi")
	fmt.Printf("John works in %s, %s\n", region, continent)

	// Invoke a function
	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn some Go.")
	default:
		fmt.Println("It's the weekend, time to rest!")
	}

	fmt.Println(time.Now().Weekday().String())

	// Using RegExp
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	var phone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)

	contact := "foo@bar.com"

	switch {
	case email.MatchString(contact):
		fmt.Println(contact, "is an email")
	case phone.MatchString(contact):
		fmt.Println(contact, "is a phone number")
	default:
		fmt.Println(contact, "is not recognized")
	}

	// Omit a condition
	rand.NewSource(time.Now().Unix())
	r := rand.Float64()
	switch {
	case r > 0.1:
		fmt.Println("Common case, 90% of the time")
	default:
		fmt.Println("10% of the time")
	}

	// Fallthrough to the next case
	switch num := 15; {
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200", num)
	}
}
