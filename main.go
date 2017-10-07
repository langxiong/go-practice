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

	lookup := make(map[string]int, 1000)

	lookup["xiongl"] = 9001

	power, exists := lookup["fangyi"]

	fmt.Println(power, exists)
	fmt.Println(lookup, " ", len(lookup))
	lookup["fangyi"] = 0
	fmt.Println(lookup, " ", len(lookup))
	delete(lookup, "fangyi")
	fmt.Println(lookup, " ", len(lookup))

	{
		xiongl := &Saiyan {
			Name: "xiongl",
			Friends: make(map[string]*Saiyan),
		}

		xiongl.Friends["fangyi"] = &Saiyan {
			Name: "fangyi",
			Friends: make(map[string]*Saiyan),
		}

		fmt.Println(xiongl)
	}
}

type Saiyan struct {
	Name string
	Friends map[string]*Saiyan
}


