package EP

import (
	npb "NPB-Golang/commons"
	"math"
	"time"
)

type ResultData struct {
	sxResult float64
	syResult float64
	qResult  [NQ]float64
}

func ParallelEP(
	np int,
	an float64,
	sx *float64,
	sy *float64,
	q []float64,
	tt *time.Duration,
) {
	const k_offset = -1

	qTemp := [NQ]float64{}

	var result ResultData

	resultChn := make(chan ResultData, np)

	start := time.Now()
	defer getExecTime(tt, &start)
	for k := 1; k <= np; k++ {
		go func(k int) {
			var (
				sxThis, syThis         = 0.0, 0.0
				t1, t2, t3, t4, x1, x2 float64
				kk, ik, l              int
				qq, x                  = [NQ]float64{}, [NK_PLUS]float64{}
				resultThis             ResultData
			)

			kk = k_offset + k
			t1 = S
			t2 = an

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

			npb.Vranlc(2*NK, &t1, A, x[:])

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
					sxThis += t3
					syThis += t4
				}
			}
			resultThis.sxResult = sxThis
			resultThis.syResult = syThis
			resultThis.qResult = qq
			resultChn <- resultThis
		}(k)
	}
	for k := 1; k <= np; k++ {
		result = <-resultChn
		*sx += result.sxResult
		*sy += result.syResult
		qTemp = result.qResult
		for j := range q {
			q[j] += qTemp[j]
		}
	}
}

func getExecTime(tt *time.Duration, start *time.Time) {
	*tt = time.Since(*start)
}
