package main

import "fmt"

type People interface {
	Eat() string
}

type Teacher struct {
	name string
}

func (t Teacher) Eat() string {
	return "老师吃"
}

type Student struct {
	name string
}

func (s *Student) Eat() string {
	return "学生吃"
}

func PeopleEat(p People) {
	fmt.Println(p.Eat())
}

func main() {
	t := Teacher{"ZYY"}
	s := Student{"LCY"}

	PeopleEat(t)
	PeopleEat(&s)
}