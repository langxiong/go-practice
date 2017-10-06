package main

import (
	"fmt"
)

type Persion struct {
	Name string
}

func (p *Persion) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan struct {
	*Persion
	Power int
}

func (s *Saiyan) Introduce() {
	fmt.Printf("Hi, this is %s\n", s.Name)
}

func main() {
	xl := Persion{
		"xiongl",
	}

	fy := &Saiyan {
		Persion: &Persion{
			"fangyi",
		},
		Power: 9001,
	}

	xl.Introduce()
	fy.Introduce()
	fy.Persion.Introduce()

	var scores [4]int
	scores[0] = 33

	newscores := [4]int {
		3333,
		1111,
		2222,
		4444,
	}

	for index, value := range newscores {
		fmt.Println(index, value)
	}
}