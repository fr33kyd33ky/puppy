package puppy

import (
	"fmt"

	"github.com/fr33kyd33ky/dog"
)

func Bark() string {
	return "Bark"
}

func Barks() string {
	return "Bark Bark Bark"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From11() {
	fmt.Println("v11")
}

func From12() {
	fmt.Println("v12")
}

func From13() {
	fmt.Println("v13")
}
