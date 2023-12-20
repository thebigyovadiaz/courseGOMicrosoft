package main

import (
	"fmt"
	"math"
	"strings"
)

type triangle struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

// Pointer in methods
func (t *triangle) doubleSize() {
	t.size *= 2
}

type coloredTriangle struct {
	triangle
	color string
}

func (ct coloredTriangle) perimeter() int {
	return ct.triangle.size * 3
}

type triangleE struct {
	size int
}

func (tE *triangleE) SetSize(size int) {
	tE.size = size
}

// Encapsulation
func (tE *triangleE) perimeter() int {
	tE.doubleSize()
	return tE.size * 3
}

func (tE *triangleE) doubleSize() {
	tE.size *= 2
}

type square struct {
	size int
}

func (s square) perimeter() int {
	return s.size * 4
}

func (s *square) doubleSize() {
	s.size *= 2
}

type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

/**
* Using Interfaces
**/
type Shape interface {
	Perimeter() float64
	Area() float64
}

type Square struct {
	size float64
}

func (s Square) Perimeter() float64 {
	return s.size * 4
}

func (s Square) Area() float64 {
	return s.size * s.size
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printInformation(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

func main() {
	t := triangle{3}
	t.doubleSize() // using pointer in function
	s := square{4}
	s.doubleSize()
	fmt.Println("Perimeter (triangle):", t.perimeter())
	fmt.Println("Perimeter (square):", s.perimeter())
	fmt.Println()

	// Methods other types
	str := upperstring("Learning GO with Microsoft!")
	fmt.Println("Before:", str, "\nAfter:", str.Upper())
	fmt.Println()

	// Embed Methods
	t2 := triangle{5}
	tC := coloredTriangle{t2, "purple"}
	fmt.Println("Size:", tC.size)
	fmt.Println("Perimeter:", tC.perimeter())
	fmt.Println("Color:", tC.color)
	tC.doubleSize()
	fmt.Println("DoubleSize:", tC.size)
	fmt.Println()

	// Encapsulation methods
	tE := triangleE{}
	tE.SetSize(10)
	fmt.Println("BeforeSize:", tE.size)
	fmt.Println("Perimeter:", tE.perimeter())
	fmt.Println("AfterSize:", tE.size)
	fmt.Println()
	fmt.Println()

	// Implement an interface
	var sShape Shape = Square{5}
	printInformation(sShape)

	c := Circle{6}
	printInformation(c)
}
