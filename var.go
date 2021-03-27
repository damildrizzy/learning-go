package main

import ("fmt")

// var (
// 	name string
// 	age  int
// 	location string
// )

// var (name, location string = "Sam", "Ibadan"
// 	age int = 24
// )

func main() {
	name := "Sam"
	age, location := 24, "Ibadan"
	fmt.Printf("I am %s, %d years old from %s ", name, age, location)
}