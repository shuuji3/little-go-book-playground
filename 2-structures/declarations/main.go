package main

import "fmt"

type Mochi struct {
	Name   string
	Power  int
	Father *Mochi
}

func NewMochi(name string, power int) *Mochi {
	return &Mochi{
		Name:  name,
		Power: power,
	}
}

func (m *Mochi) Super() {
	m.Power += 10000
}

func main() {
	emptyMochi := new(Mochi)
	emptyMochi.Name = "empty"
	normalMochi := NewMochi("Normal", 255)
	drMochi := NewMochi("Professor", 65536)
	fmt.Println(emptyMochi, normalMochi, drMochi)

	normalMochi.Super()
	fmt.Println(normalMochi)

	grandMochi := new(Mochi)
	grandMochi.Name = "Grand father"
	grandMochi.Power = 10000

	drMochi.Father = grandMochi
	fmt.Println(emptyMochi, normalMochi, drMochi)
	fmt.Printf("drMochi's father is %v\n", grandMochi)
}
