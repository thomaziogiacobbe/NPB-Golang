package IS

import (
	npb "NPB-Golang/commons"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func CreateSeq(
	seed float64,
	a float64,
	id int,
	group *sync.WaitGroup,
) {
	var (
		x, s            float64
		k               int64
		k1, k2          int64
		an              = a
		myid, num_procs int
		mq              int64
	)
	defer (*group).Done()

	myid = id
	num_procs = runtime.NumCPU()

	mq = (num_keys + int64(num_procs) - 1) / int64(num_procs)
	k1 = mq * int64(myid)
	k2 = k1 + mq
	if k2 > num_keys {
		k2 = num_keys
	}
	s = FindMySeed(myid, num_procs, int32(4*num_keys), seed, an)
	k = max_key / 4
	for i := k1; i < k2; i++ {
		x = npb.Randlc(&s, an)
		x += npb.Randlc(&s, an)
		x += npb.Randlc(&s, an)
		x += npb.Randlc(&s, an)
		key_array[i] = int64(float64(k) * x)
	}
}

func Rank(iteration int64) {
	var (
		i                           int64
		key_buff_ptr, key_buff_ptr2 []int64
		shift                       = max_key_log_2 - num_buckets_log_2
		num_bucket_keys             = int64(1) << shift
		num_procs                   = runtime.NumCPU()
		group                       sync.WaitGroup
	)
	key_array[iteration] = iteration
	key_array[iteration+MAX_ITERATIONS] = max_key - iteration
	key_buff_ptr2 = key_buff2[:]
	key_buff_ptr = key_buff1[:]
	for i = 0; i < TEST_ARRAY_SIZE; i++ {
		partial_verify_vals[i] = key_array[test_index_array[i]]
	}

	bucket_size = make([][]int64, 0, num_procs)
	for iter := 0; iter < num_procs; iter++ {
		temp := make([]int64, num_buckets)
		bucket_size = append(bucket_size, temp)
	}

	npb.ParallelFor(
		num_keys,
		int64(num_procs),
		func(myid int64, it int64) {
			bucket_size[myid][key_array[it]>>shift]++
		},
		npb.STATIC,
	)
	group.Add(num_procs)
	for myid := 0; myid < num_procs; myid++ {
		go func(myid int) {
			bucket_ptrs[myid][0] = 0
			for k := 0; k < myid; k++ {
				bucket_ptrs[myid][0] += bucket_size[k][0]
			}
			for it := int64(1); it < num_buckets; it++ {
				bucket_ptrs[myid][it] = bucket_ptrs[myid][it-1]
				for k := 0; k < num_procs; k++ {
					if k < myid {
						bucket_ptrs[myid][it] += bucket_size[k][it]
					} else {
						bucket_ptrs[myid][it] += bucket_size[k][it-1]
					}
				}
			}
			defer group.Done()
		}(myid)
	}
	group.Wait()
	npb.ParallelFor(
		num_keys,
		int64(num_procs),
		func(myid int64, it int64) {
			k := key_array[it]
			key_buff2[bucket_ptrs[myid][k>>shift]] = k
			bucket_ptrs[myid][k>>shift]++
		},
		npb.STATIC,
	)
	group.Add(num_procs - 1)
	for myid := 0; myid < num_procs-1; myid++ {
		go func(myid int) {
			for it := int64(0); it < num_buckets; it++ {
				for k := myid + 1; k < num_procs; k++ {
					bucket_ptrs[myid][it] += bucket_size[k][it]
				}
			}
			defer group.Done()
		}(myid)
	}
	group.Wait()
	npb.ParallelFor(
		num_buckets,
		int64(num_procs),
		func(myid int64, i int64) {
			var (
				k1, k2 int64
				m      int64
			)
			k1 = i * num_bucket_keys
			k2 = k1 + num_bucket_keys
			for k := k1; k < k2; k++ {
				key_buff_ptr[k] = 0
			}
			if i > 0 {
				m = bucket_ptrs[myid][i-1]
			} else {
				m = 0
			}
			for k := m; k < bucket_ptrs[myid][i]; k++ {
				key_buff_ptr[key_buff_ptr2[k]]++
			}
			key_buff_ptr[k1] += m
			for k := k1 + 1; k < k2; k++ {
				key_buff_ptr[k] += key_buff_ptr[k-1]
			}
		},
		npb.DYNAMIC,
	)

	for i = 0; i < TEST_ARRAY_SIZE; i++ {
		k := partial_verify_vals[i]
		if 0 < k && k <= num_keys-1 {
			keyRank := key_buff_ptr[k-1]
			failed := false
			switch npb.Class {
			case "S":
				if i <= 2 {
					if keyRank != (test_rank_array[i] + iteration) {
						failed = true
					} else {
						passed_verification++
					}
				} else {
					if keyRank != (test_rank_array[i] - iteration) {
						failed = true
					} else {
						passed_verification++
					}
				}
				break
			case "W":
				if i < 2 {
					if keyRank != (test_rank_array[i] + iteration - 2) {
						failed = true
					} else {
						passed_verification++
					}
				} else {
					if keyRank != (test_rank_array[i] - iteration) {
						failed = true
					} else {
						passed_verification++
					}
				}
				break
			case "A":
				if i <= 2 {
					if keyRank != (test_rank_array[i] + iteration - 1) {
						failed = true
					} else {
						passed_verification++
					}
				} else {
					if keyRank != (test_rank_array[i] - (iteration - 1)) {
						failed = true
					} else {
						passed_verification++
					}
				}
				break
			case "B":
				if i == 1 || i == 2 || i == 4 {
					if keyRank != (test_rank_array[i] + iteration) {
						failed = true
					} else {
						passed_verification++
					}
				} else {
					if keyRank != (test_rank_array[i] - iteration) {
						failed = true
					} else {
						passed_verification++
					}
				}
				break
			case "C":
				if i <= 2 {
					if keyRank != (test_rank_array[i] + iteration) {
						failed = true
					} else {
						passed_verification++
					}
				} else {
					if keyRank != (test_rank_array[i] - iteration) {
						failed = true
					} else {
						passed_verification++
					}
				}
				break
			case "D":
				if i < 2 {
					if keyRank != (test_rank_array[i] + iteration) {
						failed = true
					} else {
						passed_verification++
					}
				} else {
					if keyRank != (test_rank_array[i] - iteration) {
						failed = true
					} else {
						passed_verification++
					}
				}
				break
			}
			if failed {
				fmt.Println("Failed partial verification: iteration ", iteration, ", test key ", i, "\n")
			}
		}
	}

	if iteration == MAX_ITERATIONS {
		key_buff_ptr_global = key_buff_ptr[:]
	}
}

func FullVerify() {
	var (
		j        int64
		numProcs = runtime.NumCPU()
	)
	npb.ParallelFor(
		num_buckets,
		int64(numProcs),
		func(myid int64, j int64) {
			var k, k1 int64
			if j > 0 {
				k1 = bucket_ptrs[myid][j-1]
			} else {
				k1 = 0
			}
			for i := k1; i < bucket_ptrs[myid][j]; i++ {
				key_buff_ptr_global[key_buff2[i]] -= 1
				k = key_buff_ptr_global[key_buff2[i]]
				key_array[k] = key_buff2[i]
			}
		},
		npb.DYNAMIC,
	)

	j = 0
	npb.ParallelFor(
		num_keys-1,
		int64(numProcs),
		func(_ int64, i int64) {
			if key_array[i] > key_array[i+1] {
				atomic.AddInt64(&j, 1)
			}
		},
		npb.STATIC,
	)
	if j != 0 {
		fmt.Println("Full_verify: number of keys out of sort: ", j)
	} else {
		passed_verification += 1
	}
}
