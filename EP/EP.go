package main

import (
	npb "NPB-Golang/commons"
	"fmt"
	"math"
	"os"
)

var M int

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Invalid number of arguments")
		os.Exit(1)
	}

	args := os.Args[1:]
	getNPBClass(args[0])

	const MK = int(16)

	var MM = M - MK
	var NN = 1 << MM

	const NK = 1 << MK
	const NQ = 10
	const EPSILON = 1.0e-8
	const A = 1220703125.0
	const S = 271828183.0
	const NK_PLUS = (2 * NK) + 1

	var size float64
	var mops, t1, t2, t3, t4, x1, x2 float64
	var sx, sy, tm, an, tt, gc float64
	var sx_verify_value, sy_verify_value, sx_err, sy_err float64
	var np int32
	var ik, kk, l, k, nit int32
	var k_offset, j int32
	var verified bool
	dum := [3]float64{1.0, 1.0, 1.0}
	size = math.Pow(2.0, float64(M+1))
	var x = [NK_PLUS]float64{}
	var q = [NQ]float64{}

	fmt.Println("\n\n NAS Parallel Benchmarks 4.1 Parallel Golang version - EP Benchmark\\n\\n")
	fmt.Println(" Number of random numbers generated:", size)

	verified = false

	np = NN

	temp := []float64{dum[2]}
	npb.Vranlc(0, &dum[0], dum[1], temp)
	dum[0] = npb.Randlc(&dum[1], dum[2])
	for i := 0; i < NK_PLUS; i++ {
		x[i] = -1.0e99
	}
	mops = math.Log(math.Sqrt(math.Abs(math.Max(1.0, 1.0))))

	//TODO: add timers

	t1 = A
	npb.Vranlc(0, &t1, A, x[:])

	t1 = A

	for i := 0; i < MK+1; i++ {
		t2 = npb.Randlc(&t1, t1)
	}

	an = t1
	tt = S
	gc = 0.0
	sx = 0.0
	sy = 0.0

	for i := 0; i <= NQ-1; i++ {
		q[i] = 0.0
	}

	k_offset = -1
}
