package main

import "fmt"

func main() {
	lookup := make(map[string]int)
	lookup["goku"] = 9999
	lookup["mochi"] = 99999
	fmt.Println(lookup)

	power, found := lookup["vegeta"]
	fmt.Println(power, found)

	total := len(lookup)
	fmt.Println("total:", total)

	delete(lookup, "goku")
	fmt.Println(lookup)

	type Saiyan struct {
		Name    string
		Friends map[string]*Saiyan
	}

	saiyan := &Saiyan{Name: "goku", Friends: make(map[string]*Saiyan)}
	fmt.Println(saiyan)
	saiyan.Friends["krillin"] = &Saiyan{
		Name:    "krillin",
		Friends: make(map[string]*Saiyan),
	}
	fmt.Println(saiyan)

	addr := map[string]int{
		"Tokyo":   13,
		"Chiba":   999,
		"Ibaraki": 1000,
	}
	for pref, val := range addr {
		fmt.Println(pref, val)
	}
}
