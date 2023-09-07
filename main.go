package main

import "fmt"

func main() {
	sam := person{name: "Samuel", age: 25, hobbies: []string{"coding", "climbing", "playing guitar"}}

	sam.listHobbies()

	sam.addHobby("hiking")

	sam.listHobbies()

	sam.print()
	sam.fakeChangeAge(80)
	sam.print()

	sam.clearHobbies()

	sam.listHobbies()
}

type person struct {
	name    string
	age     int
	hobbies []string
}

func (p person) listHobbies() {
	fmt.Println(p.hobbies)
}

func (p person) print() {
	fmt.Println(p)
}

func (p *person) addHobby(h string) {
	p.hobbies = append(p.hobbies, h)
}

func (p *person) clearHobbies() {
	clear(p.hobbies)
}

func (p person) fakeChangeAge(age int) {
	p.age = age
}
