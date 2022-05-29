package main

import "fmt"

type Mochi struct {
	*Person
	Power  int
	Father *Mochi
}

type Person struct {
	Name string
}

func NewMochi(name string, power int) *Mochi {
	return &Mochi{
		Person: &Person{name},
		Power:  power,
	}
}

func (m *Mochi) Super() {
	m.Power += 10000
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s.\n", p.Name)
}

func (m *Mochi) Introduce() {
	fmt.Printf("Hi, I'm %s Mochi.\n", m.Name)
}

func main() {
	emptyMochi := new(Mochi)
	emptyMochi.Person = new(Person)
	emptyMochi.Name = "empty"

	normalMochi := NewMochi("Normal", 255)
	drMochi := NewMochi("Professor", 65536)
	fmt.Println(emptyMochi, normalMochi, drMochi)

	normalMochi.Super()
	fmt.Println(normalMochi)

	grandMochi := new(Mochi)
	grandMochi.Person = new(Person)
	grandMochi.Name = "Grand father"
	grandMochi.Power = 10000

	drMochi.Father = grandMochi
	fmt.Println(emptyMochi, normalMochi, drMochi)
	fmt.Printf("drMochi's father is %v\n", grandMochi)

	drMochi.Person.Introduce() // Hi, I'm Professor.
	drMochi.Introduce()        // Hi, I'm Professor Mochi.
	grandMochi.Introduce()     // Hi, I'm Grand father Mochi.
}
