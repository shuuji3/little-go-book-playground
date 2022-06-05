package main

import "fmt"

func main() {
	var scores [10]int
	scores[0] = 339
	scores[8] = 10
	fmt.Println("scores:", scores, len(scores))

	scores2 := [4]int{10, 200, 3000, 4000}
	fmt.Println("scores2:", scores2)
	for i, v := range scores2 {
		fmt.Println(i, v)
	}
}
