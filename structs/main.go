package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
}

func main() {
	krunal := person{
		firstName: "Krunal",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "krunal@gmail.com",
			zipCode: 560021,
		},
	}
	ts := "Test String"
	fmt.Println(*&ts)
	krunal.updateName("Kunnu")
	krunal.print(&ts)
	fmt.Println(ts)
}

func (p *person) updateName(firstName string) {
	p.firstName = firstName
}

func (p person) print(ts *string) {
	fmt.Printf("%+v\n", p)
	*ts = "asdsad"
	fmt.Println(*ts)
}
