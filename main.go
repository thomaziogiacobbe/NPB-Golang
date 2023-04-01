package main

import (
	ep "NPB-Golang/EP"
	is "NPB-Golang/IS"
	"NPB-Golang/commons"
	"fmt"
	"os"
)

type classFunc func(int) string

// required params: {benchmark} CLASS={class}
// optional param: {-f file_name}
func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		_, _ = fmt.Fprintln(os.Stderr, "Invalid number of required arguments")
		os.Exit(1)
	}

	var benchmark func()

	commons.Benchmark, benchmark = getBenchmark(args[0])
	commons.Class = getClass(args[1])
	if len(args) > 2 && args[2] == "-f" {
		var fileArg string
		if len(args) > 3 {
			fileArg = args[3]
		}
		commons.File = getFile(commons.Benchmark, commons.Class, fileArg)
	}
	benchmark()
}

func getBenchmark(benchmark_arg string) (benchmark string, benchmarkFunc func()) {
	switch benchmark_arg {
	case "EP":
		benchmark = "EP"
		benchmarkFunc = ep.ExecEP
	case "IS":
		benchmark = "IS"
		benchmarkFunc = is.ExecIS
	default:
		fmt.Println("Incorrect benchmark argument")
		os.Exit(1)
	}
	return benchmark, benchmarkFunc
}

func getClass(class_arg string) (class string) {
	switch class_arg {
	case "CLASS=S":
		class = "S"
	case "CLASS=W":
		class = "W"
	case "CLASS=A":
		class = "A"
	case "CLASS=B":
		class = "B"
	case "CLASS=C":
		class = "C"
	case "CLASS=D":
		class = "D"
	case "CLASS=E":
		class = "E"
	default:
		_, _ = fmt.Fprintln(os.Stderr, "Invalid class argument")
		os.Exit(1)
	}
	return class
}

func getFile(benchmark string, class string, file_arg string) (file *os.File) {
	var fileName string
	if len(file_arg) != 0 {
		fileName = file_arg
	} else {
		fileName = benchmark + "_" + class + ".txt"
	}
	var err error
	file, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "file could not be open/created")
		return nil
	}
	return file
}
