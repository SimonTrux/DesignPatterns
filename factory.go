package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func Greet(p Person) {
	fmt.Printf("Hi ! My name is %s and I'm %d years old.", p.Name, p.Age)
	fmt.Println()
}

// Factory function to make Persons
func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age: age,
	}
}


func main() {
	fmt.Println("Greeting a couple of people and using a factory method :")

// Making new person manually
	Bob := Person{
	Name: "Bob",
	Age: 42}
	Greet(Bob)

	Greet(Person{
		Name:"Alice",
		Age: 24})

// Using factory in Greet method scope only
	Greet(NewPerson("Jane", 18))

// Or declaring first for a broader scope
	var Lee Person = NewPerson("Lee", 99)
	Greet(Lee)

	Mike := NewPerson("Mike", 10)
	Greet(Mike)
}

