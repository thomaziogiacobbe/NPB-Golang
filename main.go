package main

import (
	ep "NPB-Golang/EP"
	is "NPB-Golang/IS"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Println("Invalid number of arguments")
		os.Exit(1)
	}

	args := os.Args[1:2]
	switch args[0] {
	case "EP":
		ep.ExecEP()
	case "IS":
		is.ExecIS()
	default:
		fmt.Println("Incorrect benchmark argument")
		os.Exit(1)
	}

}
