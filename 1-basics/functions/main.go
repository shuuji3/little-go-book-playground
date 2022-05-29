package main

import (
	"fmt"
	"os"
)

func log(message string) {
	fmt.Println(message)
}

func add(a, b int) int {
	return a + b
}

func power(name string) (int, bool) {
	if name == "Mochi" {
		return add(65535, 1), true
	}
	return 0, false
}

func main() {
	//name := "Mochi"
	name := "Dr.Mochi"
	power, exists := power(name)
	if !exists {
		log(fmt.Sprintf("Cannot find a power value for %s...", name))
		os.Exit(1)
	}
	fmt.Printf("%s's power is over %d!\n", name, power)
}
