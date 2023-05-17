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

type Employee struct {
	Person
	Salary int
	Boss   *Manager
}

type Manager struct {
	Employee
}

//a Person interface to pass to SayHello - to make it accept generic person, without specifying if it's a manager or employee

type IPerson interface {
	GetName() string
}

func (e Employee) GetName() string {
	return e.First
}

func (m Manager) GetName() string {
	return m.First
}

//uupper case functions are public!
func SayHello(p IPerson) {
	fmt.Println("Hello,", p.GetName())
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

	m := &Manager{
		Employee{
			Person: Person{
				First: "Donald",
				Last:  "Trump",
			},
			Salary: 150000,
			Boss:   nil,
		},
	}

	e := Employee{
		Person: Person{
			First: "Mike",
			Last:  "Pence",
		},
		Salary: 40000,
		Boss:   nil,
	}

	//to print values you don't have to dig into the fields
	fmt.Println(e.First)
	fmt.Println(m.First)

	me := NewPerson("laura", "liparulo")

	fmt.Println(me)

	SayHello(m)
	SayHello(e)

}
