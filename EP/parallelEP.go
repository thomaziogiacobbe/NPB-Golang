package main

import (
	npb "NPB-Golang/commons"
	"math"
)

func parallelEP(
	np int,
	an float64,
	sx float64,
	sy float64,
) {
	var t1, t2, t3, t4, x1, x2 float64
	var kk, ik, l int
	var qq, x = [NQ]float64{}, [NK_PLUS]float64{}
	var k_offset = -1

	for k := 1; k <= np; k++ {
		go func(k int) {
			kk = k_offset + k
			t1 = S
			t2 = an

			//TODO: thread id is missing timer

			for i := 0; i <= 100; i++ {
				ik = kk / 2
				if (2 * ik) != kk {
					t3 = npb.Randlc(&t1, t2)
				}
				if ik == 0 {
					break
				}
				t3 = npb.Randlc(&t2, t2)
				kk = ik
			}

			//TODO: timer_start(2)
			npb.Vranlc(2*NK, &t1, A, x[:])
			//TODO: timer_stop(2)

			//TODO: timer_start(1)
			for i := 0; i < NK; i++ {
				x1 = 2.0*x[2*i] - 1.0
				x2 = 2.0*x[2*i+1] - 1.0
				t1 = math.Pow(x1, 2) + math.Pow(x2, 2)
				if t1 <= 1.0 {
					t2 = math.Sqrt(-2.0 * math.Log(t1) / t1)
					t3 = x1 * t2
					t4 = x2 * t2
					l = int(math.Max(math.Abs(t3), math.Abs(t4)))
					qq[l] += 1.0
					//TODO: reduction
					sx = sx + t3
					sy = sy + t4
				}
			}
			//TODO: timer_stop(1)
			//TODO: missing mutex for array
		}(k)
	}
}
