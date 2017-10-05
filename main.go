package main

import (
	"fmt"
)

type Student struct {
	Name string
	Power int
	Classmate *Student
}

func main() {
	xiongl := Student {
		"xiongl",
		10000,
		nil,
	}

	xiongl1 := Student {

	}

	xiongl2 := Student {
		Name: "xiongl",
	}

	xiongl.Super()
	xiongl1.Super()
	xiongl2.Super()

	fmt.Printf("%s's power is %d\n", xiongl.Name, xiongl.Power)
	fmt.Printf("%s's power is %d\n", xiongl1.Name, xiongl1.Power)
	fmt.Printf("%s's power is %d\n", xiongl2.Name, xiongl2.Power)

	xl := NewStudent("xl", 99999)
	xl.Super()
	fmt.Printf("%s's power is %d, classmate %s\n", xl.Name, xl.Power, xl.Classmate.Name)
}

func getPower() int {
	return 9001
}

func (s *Student) Super() {
	s.Power += 10000
}

func NewStudent(name string, power int) *Student {
	return &Student {
		Name: name,
		Power: power,
		Classmate: nil,
	}
}