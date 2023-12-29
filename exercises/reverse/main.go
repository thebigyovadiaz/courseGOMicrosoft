package main

import "fmt"

func reversedString(data []string) []string {
	var newArr []string
	//newARR := make([]string, 0)

	if len(data) > 0 {
		for i := len(data) - 1; i >= 0; i-- {
			newArr = append(newArr, data[i])
		}
	}

	return newArr
}

func middleFor(data []int) bool {
	middle := len(data) / 2
	fmt.Println("Middle >>", middle)
	lenData := len(data)

	firstArr := make([]int, 0)
	secondArr := make([]int, 0)

	for i := 0; i < middle+1; i++ {
		for j := 0; j < middle; j++ {
			if i == j {
				firstArr = append(firstArr, data[j]+1)
			}
		}

		for k := lenData - 1; k > 248; k-- {
			if k == lenData-i {
				secondArr = append(secondArr, data[k]+1)
			}
		}
	}

	fmt.Printf("firstData >> %v, len >> %v\n\n", firstArr, len(firstArr))
	fmt.Printf("secondData >> %v, len >> %v", secondArr, len(secondArr))

	return true
}

func main() {
	tapp := []string{"o", "l", "l", "e", "H"}
	result := reversedString(tapp)
	fmt.Println("result >>", result)

	arr := make([]int, 0)
	number := 500
	for i := 0; i < number; i++ {
		arr = append(arr, i)
	}
	middleFor(arr)
}
