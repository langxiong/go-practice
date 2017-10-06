package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	scores := make([]int, 100)

	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}

	sort.Ints(scores)

	worst := make([]int, 5)
	copy(worst[2:5], scores[:5])
	fmt.Println(worst)
}
