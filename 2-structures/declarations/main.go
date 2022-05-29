package main

import "fmt"

type Mochi struct {
	Name  string
	Power int
}

func main() {
	emptyMochi := Mochi{}
	mochi := Mochi{
		Name:  "Normal Mochi",
		Power: 255,
	}
	drMochi := Mochi{"Professor", 65536}
	fmt.Println(emptyMochi, mochi, drMochi)
}
