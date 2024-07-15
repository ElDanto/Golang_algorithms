package controller

import (
	"algorithms/console"
	fibonacci "algorithms/fibonacci"
	"fmt"
)

func ViewActions() (test int) {
	var action int
	fmt.Println("Choose the action: [1]view algorithm, [2]exit.")
	fmt.Scanln(&action)

	var msg string
	switch action {
	case 1:
		msg = viewAlgorithm()

	case 2:
		console.ClearScreen()
		fmt.Println("See you space cowboy...")
		return 0
	default:
		msg = "Action not recognized"

	}

	if msg != "" {
		fmt.Println(msg)
	}
	ViewActions()
	return 0
}

func viewAlgorithm() (msg string) {

	fmt.Println("Enter algorithm name: ")
	var algorithm string
	fmt.Scanln(&algorithm)

	switch algorithm {
	case "fibonacci":
		fibonacci.View()
	default:
		return "This algorithm doesn't support"
	}
	return ""
}
