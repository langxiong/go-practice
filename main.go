package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan struct {
	*Person
	Power int
}

func (s *Saiyan) Introduce() {
	fmt.Printf("Hi, this is %s\n", s.Name)
}

func main() {
	scores := make([]int, 5)

	scores = append(scores, 3)

	fmt.Println(scores)

	names := []string{
		"xiongl",
		"fangy",
	}

	fmt.Println(names)

	xl := Saiyan {
		&Person {"xiongl",},
		9999,
	}

	fy := Saiyan {
		&Person {"fangyi",},
		10000,
	}

	{
		saiyans := []*Saiyan{};
		
		saiyans = append(saiyans, &xl)
		saiyans = append(saiyans, &fy)

		powers := extractPower(saiyans)
		fmt.Println(powers)
	}

	{
		saiyans := []Saiyan{};
		
		saiyans = append(saiyans, xl)
		saiyans = append(saiyans, fy)

		powers := extractPowerPoint(&saiyans)
		fmt.Println(powers)
	}
}

func extractPowerPoint(saiyans *[]Saiyan) []int {
	powers := make([]int, len(*saiyans))

	for i, s := range(*saiyans) {
		powers[i] = s.Power
	}

	return powers
}

func extractPower(saiyans []*Saiyan) []int{
	powers := make([]int, len(saiyans))

	for i, s := range(saiyans) {
		powers[i] = s.Power
	}

	return powers
}