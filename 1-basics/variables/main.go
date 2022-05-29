package main

import "fmt"

func main() {
	name, power := "Mochi", getPower()
	fmt.Printf("%s's power is over %d!\n", name, power) // Mochi's power is over 65536!
}

func getPower() int {
	return 65536
}
