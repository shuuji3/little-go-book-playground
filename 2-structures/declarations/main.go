package main

import "fmt"

type Mochi struct {
	Name  string
	Power int
}

func main() {
	emptyMochi := &Mochi{}
	normalMochi := &Mochi{
		Name:  "Normal",
		Power: 255,
	}
	drMochi := &Mochi{"Professor", 65536}
	fmt.Println(emptyMochi, normalMochi, drMochi)

	Super(normalMochi)
	fmt.Println(normalMochi)
}

func Super(m *Mochi) {
	m.Power += 10000
}
