package main

import "fmt"

type Mochi struct {
	Name  string
	Power int
}

func (m *Mochi) Super() {
	m.Power += 10000
}

func main() {
	emptyMochi := &Mochi{}
	normalMochi := &Mochi{
		Name:  "Normal",
		Power: 255,
	}
	drMochi := &Mochi{"Professor", 65536}
	fmt.Println(emptyMochi, normalMochi, drMochi)

	normalMochi.Super()
	fmt.Println(normalMochi)
}
