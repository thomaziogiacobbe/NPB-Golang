package main

import (
	"fmt"
	"os"
)

func getNPBClass(class string) {
	switch class {
	case "S":
		M = 24
	case "W":
		M = 25
	case "A":
		M = 28
	case "B":
		M = 30
	case "C":
		M = 32
	case "D":
		M = 36
	case "E":
		M = 40
	default:
		fmt.Println("Incorrect class argument")
		os.Exit(1)
	}
}
