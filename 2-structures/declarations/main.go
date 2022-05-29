package main

import "fmt"

type Mochi struct {
	Name  string
	Power int
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
	emptyMochi := NewMochi("empty", 0)
	normalMochi := NewMochi("Normal", 255)
	drMochi := NewMochi("Professor", 65536)
	fmt.Println(emptyMochi, normalMochi, drMochi)

	normalMochi.Super()
	fmt.Println(normalMochi)
}
