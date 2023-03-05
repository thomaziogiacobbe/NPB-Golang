package EP

import (
	npb "NPB-Golang/commons"
	"fmt"
	"math"
	"os"
	"time"
)

var (
	M  int
	MM int
	NN int
)

const (
	MK      = int(16)
	NK      = 1 << MK
	NQ      = 10
	A       = 1220703125.0
	S       = 271828183.0
	NK_PLUS = (2 * NK) + 1
)

func ExecEP() {
	if len(os.Args) <= 1 {
		fmt.Println("Invalid number of arguments")
		os.Exit(1)
	}

	args := os.Args[1:]
	getNPBClass(args[0])

	MM = M - MK
	NN = 1 << MM

	var size float64
	var mops, t1 float64
	var sx, sy, an, gc float64
	var np int
	var nit int32
	var verified bool
	dum := [3]float64{1.0, 1.0, 1.0}
	size = math.Pow(2.0, float64(M+1))
	var x = [NK_PLUS]float64{}
	var q = [NQ]float64{}
	var tt time.Duration

	fmt.Println("\n\n NAS Parallel Benchmarks 4.1 Parallel Golang version - EP Benchmark\n")
	fmt.Println(" Number of random numbers generated:", size, "\n")

	verified = false

	np = NN

	temp := []float64{dum[2]}
	npb.Vranlc(0, &dum[0], dum[1], temp)
	dum[0] = npb.Randlc(&dum[1], dum[2])
	for i := 0; i < NK_PLUS; i++ {
		x[i] = -1.0e99
	}
	mops = math.Log(math.Sqrt(math.Abs(math.Max(1.0, 1.0))))

	t1 = A
	npb.Vranlc(0, &t1, A, x[:])

	t1 = A

	for i := 0; i < MK+1; i++ {
		npb.Randlc(&t1, t1)
	}

	an = t1
	gc = 0.0
	sx = 0.0
	sy = 0.0

	ParallelEP(np, an, &sx, &sy, q[:], &tt)

	for i := 0; i < NQ-1; i++ {
		gc = gc + q[i]
	}

	nit = 0

	verified = verify(sx, sy)
	mops = math.Pow(2.0, float64(M+1)) / tt.Seconds() / 1000000.0

	npb.PrintEPResults(&tt,
		M,
		gc,
		sx,
		sy,
		NQ,
		q[:],
	)

	npb.Print_results("EP",
		args[0],
		int(size),
		nit,
		&tt,
		mops,
		"Random numbers generated",
		verified)
}
