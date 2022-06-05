package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Mochi struct {
	Power int
}

func main() {
	scores1 := []int{1, 3, 9, 16}
	fmt.Println(scores1, cap(scores1))

	scores2 := make([]int, 0, 10)
	fmt.Println(scores2, len(scores2))
	scores2 = append(scores2, 314, 22)
	fmt.Println(scores2, len(scores2))

	scores2 = scores2[0:10]
	scores2[7] = 700
	fmt.Println(scores2)

	scores3 := make([]int, 0, 5)
	c := cap(scores3)
	println(c)

	for i := 0; i < 2500; i++ {
		scores3 = append(scores3, i)

		if cap(scores3) != c {
			c = cap(scores3)
			println(c)
		}
	}

	scores4 := make([]int, 5)
	fmt.Println(scores4)
	scores4 = append(scores4, 999)
	fmt.Println(scores4)

	names := []string{"leto", "gato", "cat", "paul"}
	checks := make([]bool, 4)
	var text []string
	points := make([]int, 5, 10)
	fmt.Println(names, checks, text, points)

	mochis := make([]*Mochi, 2)
	mochis[0] = &Mochi{100}
	mochis[1] = &Mochi{200}
	fmt.Println("mochis:", mochis)

	powers := make([]int, len(mochis))
	for i, m := range mochis {
		powers[i] = m.Power
	}

	fmt.Println("powers:", powers)

	// slice
	scores5 := []int{1, 2, 3, 4, 5}
	slice := scores5[2:4]
	slice[0] = 999
	fmt.Println(slice)
	fmt.Println(scores5)

	// copy
	scores6 := make([]int, 100)
	for i := 0; i < len(scores6); i++ {
		scores6[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores6)

	worst := make([]int, 5)
	copy(worst, scores6[:5])

	fmt.Println(scores6)
	fmt.Println("worst:", worst)
}
