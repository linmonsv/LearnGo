package main

import (
	"fmt"
	"math"
)

//A method is a function with an implicit first argument, called a receiver."

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type Human struct {
	name  string
	phone string
}

type Student struct {
	Human
	school string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//method重写
func (s *Student) SayHi() {
	fmt.Printf("Hi, I am %s you can find me in %s\n", s.name, s.school)
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is:", r1.area())
	fmt.Println("Area of r2 is:", r2.area())
	fmt.Println("Area of c1 is:", c1.area())
	fmt.Println("Area of c2 is:", c2.area())

	//method继承
	sam := Student{Human{"Sam", "222-222-XXXX"}, "MIT"}
	sam.SayHi()

	//大写开头的为共有，小写开头的为私有)
}
