package commons

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"runtime"
	"strconv"
	"time"
)

func Print_results(
	name string,
	classNpb string,
	epSize int,
	niter int32,
	tt *time.Duration,
	mops float64,
	optype string,
	passedVerification bool,
	fileToWrite string,
) {
	var verifyString string
	if passedVerification {
		verifyString = "SUCCESSFUL"
	} else {
		verifyString = "UNSUCCESSFUL"
	}

	tableHeader := table.Row{name + " Benchmark Completed"}
	//TODO: others benchmarks to be defined, for now just printing size of EP
	tableRows := []table.Row{
		{"Class", classNpb},
		{"Size", epSize},
		{"Number of available threads", runtime.NumCPU()},
		{"Number of iterations", niter},
		{"Time in seconds", *tt},
		{"Mop/s total", mops},
		{"Operation type", optype},
		{"Verification", verifyString},
		{"NPB Version", "4.1"},
	}
	tw := table.NewWriter()
	tw.SetStyle(table.StyleLight)
	tw.SetOutputMirror(os.Stdout)
	tw.AppendHeader(tableHeader)
	tw.AppendRows(tableRows)
	tw.Render()

	if len(fileToWrite) != 0 {
		file, err := os.OpenFile(fileToWrite, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			fmt.Fprintln(os.Stderr, "file ", fileToWrite, " could not be open/created")
		} else {
			file.WriteString(strconv.FormatFloat((*tt).Seconds(), 'g', -1, 64) + "\n")
		}
	}
}

func PrintEPResults(
	tm *time.Duration,
	m int,
	gc float64,
	sx float64,
	sy float64,
	nq int,
	q []float64,
) {
	tableHeader := table.Row{"EP Benchmark Results"}
	tableRows := []table.Row{
		{"CPU Time", *tm},
		{"N ", "2^" + strconv.Itoa(m)},
		{"No. Gaussian Pairs", gc},
		{"Sums", strconv.FormatFloat(sx, 'g', -1, 64) + " " + strconv.FormatFloat(sy, 'g', -1, 64)},
	}

	var tableRowsSums []table.Row

	for i := 0; i < nq-1; i++ {
		tableRowsSums = append(tableRowsSums, []interface{}{i, q[i]})
	}

	tw := table.NewWriter()
	tw.SetStyle(table.StyleLight)
	tw.SetOutputMirror(os.Stdout)
	tw.AppendHeader(tableHeader)
	tw.AppendRows(tableRows)
	tw.AppendSeparator()
	tw.AppendRow([]interface{}{"Counts"})
	tw.AppendRows(tableRowsSums)
	tw.Render()
}
