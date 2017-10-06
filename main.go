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
	scores := []int{1, 2, 3, 4, 5}
	slice := scores[2:4]
	slice[0] = 999
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
	{
		scores := []int{1, 2, 3, 4, 5}
		scores = scores[:len(scores)-1]
		fmt.Println(scores)
	}
	{
		scores := []int{1, 2, 3, 4, 5}
		scores = removeAtIndex(scores, 2)
		fmt.Println(scores)
	}
}

func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
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