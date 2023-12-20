package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	employee, err := getInformation(1001)
	if err != nil {
		// Something is wrong. Do something.
		fmt.Println("Error:", err.Error())
	} else {
		log.SetOutput(file)
		log.SetPrefix("[CourseMicrosoftGO] ")
		log.Println("Employee:", employee)
		fmt.Print(employee)
	}
}

func getInformation(id int) (*Employee, error) {
	for tries := 0; tries < 3; tries++ {
		employee, err := apiCallEmployee(1000)
		if err == nil {
			return employee, nil //fmt.Errorf("Got an error when getting the employee information: %v", err)
		}

		fmt.Println("Server is not responding, retrying ...")
		time.Sleep(time.Second * 2)
	}
	return nil, fmt.Errorf("Server has failed to respond to get the employee information.")
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}
