package main

import "fmt"

type Human struct {
	Name   string
	Gender bool
	Weight uint
	Growth uint
}

func (h *Human) AskQuestion(question string) string {
	if len(question) > 42 {
		return "Yes"
	}
	if len(question) < 42 {
		return "No"
	}
	return "I don`t know"
}

func (h *Human) Feed(foodNum uint) {
	h.Weight += foodNum
}

type Action struct {
	Human
	comment string
}

func main() {
	action := Action{
		Human: Human{
			Name:   "Oleg",
			Gender: true,
			Weight: 66,
			Growth: 175,
		},
		comment: "unnecessary actions",
	}
	action.Feed(23)
	fmt.Println(action.Weight)
}
