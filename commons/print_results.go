package commons

import (
	"fmt"
	"runtime"
)

// TODO: better print results, maybe use https://github.com/jedib0t/go-pretty
func Print_results(
	name string,
	classNpb string,
	epSize int,
	niter int32,
	t float64,
	mops float64,
	optype string,
	passedVerification bool,
) {
	fmt.Println("\n\n", name, " Benchmark Completed")
	fmt.Println(" Class =", classNpb)
	//TODO: others benchmarks to be defined, for now just printing size of EP
	fmt.Println(" Size =", epSize)
	fmt.Println(" Number of available threads =", runtime.NumCPU())
	fmt.Println(" Number of iterations =", niter)
	fmt.Println(" Time in seconds =", t)
	fmt.Println(" Mop/s total =", mops)
	fmt.Println(" Operation type =", optype)
	if passedVerification {
		fmt.Println(" Verification =", "SUCCESSFUL")
	} else {
		fmt.Println(" Verification =", "UNSUCCESSFUL")
	}
	fmt.Println(" NPB Version = 4.1")
}
