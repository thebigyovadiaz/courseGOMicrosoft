package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/**
	* Using Arrays
	**/
	var ar [3]int // init, length and type of element in array
	ar[1] = 20

	for i := 0; i < len(ar); i++ {
		fmt.Println(ar[i])
	}

	cities := [5]string{"Lithuania", "Paris", "Berlin", "Warsaw"}
	fmt.Println(cities)
	citiesII := [...]string{"Lithuania", "Paris", "Berlin", "Warsaw"} // Ellipsis in array
	fmt.Println(citiesII)

	numbers := [...]int{99: -1}
	fmt.Println("First Position:", numbers[0])
	fmt.Println("Last Position:", numbers[99])
	fmt.Println("Length:", len(numbers))
	fmt.Println()

	// Multidimensional arrays
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("ROW: ", i, twoD[i])
	}
	fmt.Println("\nAll at once: ", twoD)

	var threeD [3][5][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 2; k++ {
				threeD[i][j][k] = (i + 1) * (j + 1) * (k + 1)
			}
		}
	}
	fmt.Println("\nAll at once: ", threeD)
	fmt.Println()
	fmt.Println()

	/**
	* Using Slices
	**/
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(months)
	fmt.Println("Length: ", len(months))
	fmt.Println("Capacity: ", cap(months))

	// Slice items
	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))

	// Append items
	var appendNumbers []int
	for i := 0; i < 10; i++ {
		appendNumbers = append(appendNumbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(appendNumbers), appendNumbers)
	}
	fmt.Println()

	// Remove items
	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2

	if remove < len(letters) {
		fmt.Println("Before:", letters, "Remove:", letters[remove])
		letters = append(letters[:remove], letters[remove+1:]...)
		fmt.Println("After:", letters)
	}
	fmt.Println()

	// Create copies of slices
	lettersII := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before:", lettersII)

	slice1 := lettersII[:2]

	/* slice2 := lettersII[1:4]
	slice1[1] = "Z" // Rewrite item in all slices */

	slice2 := make([]string, 3)
	copy(slice2, lettersII[1:4])

	slice1[1] = "Z" // Rewrite item only in slice2

	fmt.Println("After:", lettersII)
	fmt.Println("Slice2: ", slice2)
	fmt.Println()
	fmt.Println()

	/**
	* Using Maps
	**/
	studentsAge := map[string]int{
		"john": 32,
		"bob":  31,
	}

	// Add items
	studentsAge["peter"] = 28
	studentsAge["kate"] = 30
	fmt.Println(studentsAge)
	fmt.Println()

	// Access item
	fmt.Println("Bob's age is", studentsAge["bob"])

	age, exist := studentsAge["christy"]
	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found")
	}
	fmt.Println()

	// Remove item
	delete(studentsAge, "bob")
	fmt.Println(studentsAge)
	fmt.Println()

	// Loop
	for name, age := range studentsAge {
		fmt.Printf("Name: %s\tAge: %d\n", name, age)
	}

	for _, age := range studentsAge {
		fmt.Printf("Age: %d\n", age)
	}

	// Return only key of map
	for name := range studentsAge {
		fmt.Printf("Names: %s\n", name)
	}
	fmt.Println()
	fmt.Println()

	/**
	* Using Struct
	**/
	var adele Employee
	adele = Employee{23, "Kate", "Adele", "United Kingdom"}
	fmt.Println(adele.ID, adele.Address)

	employees := Employee{ID: 1, FirstName: "John", LastName: "Doe"}
	fmt.Println(employees)

	employeeAdele := &employees
	employeeAdele.ID = 2
	employeeAdele.FirstName = "Kate"
	employeeAdele.LastName = "Adele"
	employeeAdele.Address = "London"
	fmt.Println(employees)
	fmt.Println()

	// Embedding
	var miggi Person = Person{
		ID:        24,
		FirstName: "Miguel",
		LastName:  "Cabrera",
		Address:   "Detroit - USA",
	}

	companyEmployee := Employees{
		Person:    miggi,
		ManagerID: 1,
	}

	fmt.Println(companyEmployee)
	fmt.Println(companyEmployee.FirstName)
	fmt.Println()

	// Encode - Decode structs with JSON
	companyEmployees := []Employees{
		{
			Person: miggi,
		},
		{
			Person: Person{
				ID:        27,
				FirstName: "José",
				LastName:  "Altuve",
				Address:   "Houston - USA",
			},
		},
	}

	data, _ := json.Marshal(companyEmployees)
	fmt.Println("Encoded: ", data)

	var decodedEmployees []Employees
	json.Unmarshal(data, &decodedEmployees)
	fmt.Printf("Decoded: %v", decodedEmployees)
}

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

type Employees struct {
	Person
	ManagerID int
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

type PersonJSON struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string `json:"lastname"`
	Address   string `json:"address,omitempty"`
}
