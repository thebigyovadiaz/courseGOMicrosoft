package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/**
	* Basic Loop
	**/
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println("sum of 1..100 is", sum)

	/**
	* Empty pre-statements and post-statements
	**/
	var num int64
	rand.NewSource(time.Now().Unix())
	for num != 5 {
		num = rand.Int63n(15)
		fmt.Println(num)
	}

	/**
	* Infinity Loop and Break statement
	**/
	var num1 int32
	sec := time.Now().Unix()
	rand.NewSource(sec)

	for {
		fmt.Print("Writing inside the loop...")
		if num1 = rand.Int31n(10); num1 == 5 {
			fmt.Println("finish!")
			break
		}
		fmt.Println(num1)
	}

	/**
	* Continue statement
	**/
	sum1 := 0
	for num := 1; num <= 100; num++ {
		if num%5 == 0 {
			continue
		}
		sum1 += num
	}
	fmt.Println("The sum of 1 to 100, but excluding numbers divisible by 5, is", sum1)
}
