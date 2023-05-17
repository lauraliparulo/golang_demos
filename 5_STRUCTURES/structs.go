package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
	Phone Phone
}

type Phone struct {
	AreaCode string
	Prefix   string
	Suffix   string
}

//uupper case functions are public!
func (p Person) SayHello() {
	fmt.Println("Hello, %s\n", p.First)
}

// Constructor
// the same type can be declared once
func NewPerson(first, last string) *Person {
	return &Person{
		First: first,
		Last:  last,
	}
}

func main() {

	person := Person{
		First: "John",
		Last:  "Doe",
		Age:   30,
		Phone: Phone{
			AreaCode: "+49",
			Prefix:   "177",
			Suffix:   "82739847",
		},
	}

	fmt.Println(person)

	me := NewPerson("laura", "liparulo")

	fmt.Println(me)

}
