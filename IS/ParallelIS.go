package IS

import (
	npb "NPB-Golang/commons"
	"runtime"
)

func CreateSeq(
	seed float64,
	a float64,
	id int,
) {
	var (
		x, s            float64
		k               int64
		k1, k2          int64
		an              = a
		myid, num_procs int
		mq              int64
	)

	myid = id
	num_procs = runtime.NumCPU()

	mq = (num_keys + int64(num_procs) - 1) / int64(num_procs)
	k1 = mq * int64(myid)
	k2 = k1 + mq
	if k2 > num_keys {
		k2 = num_keys
	}
	FindMySeed(myid, num_procs, int32(4*num_keys), seed, an)
	k = max_key / 4
	for i := k1; i < k2; i++ {
		x = npb.Randlc(&s, an)
		x += npb.Randlc(&s, an)
		x += npb.Randlc(&s, an)
		x += npb.Randlc(&s, an)
		key_array[i] = int64(float64(k) * x)
	}
}

func Rank(iteration int) {
	var (
		i, k                        int64
		key_buff_ptr, key_buff_ptr2 *int64
		shift                       = max_key_log_2 - num_buckets_log_2
		num_bucket_keys             = int32(1) << shift
	)
	key_array[iteration] = int64(iteration)
	key_array[iteration+MAX_ITERATIONS] = max_key - int64(iteration)
	//TODO: converter []int64 pra *int64, ou alterar o tipo da variavel aqui
	//key_buff_ptr2 = key_buff2
	//key_buff_ptr = key_buff1

	//TODO: partial verify
}
