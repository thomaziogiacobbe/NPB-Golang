package IS

import (
	npb "NPB-Golang/commons"
	"runtime"
)

func CreateSeq(
	seed float64,
	a float64,
) {
	var (
		x, s            float64
		k               int64
		k1, k2          int64
		an              = a
		myid, num_procs int
		mq              int64
		id              int
	)

	myid = id
	num_procs = runtime.NumCPU()

	mq = int64((num_keys + num_procs - 1) / num_procs)
	k1 = mq * int64(myid)
	k2 = k1 + mq
	if k2 > int64(num_keys) {
		k2 = int64(num_keys)
	}
	FindMySeed(myid, num_procs, int32(4*num_keys), seed, an)
	k = int64(max_key / 4)
	for i := k1; i < k2; i++ {
		x = npb.Randlc(&s, an)
		x += npb.Randlc(&s, an)
		x += npb.Randlc(&s, an)
		x += npb.Randlc(&s, an)
		key_array[i] = float64(k) * x
	}
}
