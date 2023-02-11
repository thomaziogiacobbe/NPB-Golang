package main

import (
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

	//var MM = M - MK
	//var NN = 1 << MM

	const NK = 1 << MK
	const NQ = 10
	const EPSILON = 1.0e-8
	const A = 1220703125.0
	const S = 271828183.0
	const NK_PLUS = (2 * NK) + 1

	var size float64
	//var mops, t1, t2, t3, t4, x1, x2 float64
	//var sx, sy, tm, an, tt, gc float64
	//var sx_verify_value, sy_verify_value, sx_err, sy_err float64
	//var np int32
	//var i, ik, kk, l, k, nit int32
	//var k_offset, j int32
	//var verified bool
	//dum := [3]float64{1.0, 1.0, 1.0}
	size = math.Pow(2.0, float64(M+1))

	fmt.Println(size)

	fmt.Println(verify(1, 1))
}
