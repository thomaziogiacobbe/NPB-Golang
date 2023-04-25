package IS

import (
	"fmt"
	"runtime"
	"sync"
)

func RankNoBuckets(iteration int64) {
	var (
		i                           int64
		key_buff_ptr, key_buff_ptr2 []int64
		num_procs                   = runtime.NumCPU()
		group                       sync.WaitGroup
	)
	_, _, _ = key_buff_ptr2, num_procs, group

	key_array[iteration] = iteration
	key_array[iteration+MAX_ITERATIONS] = max_key - iteration
	key_buff_ptr2 = key_array[:]
	key_buff_ptr = key_buff1[:]
	for i = 0; i < TEST_ARRAY_SIZE; i++ {
		partial_verify_vals[i] = key_array[test_index_array[i]]
	}

	//TODO: implementation of no buckets on rank

	partialVerify(iteration, key_buff_ptr[:])

	if iteration == MAX_ITERATIONS {
		key_buff_ptr_global = key_buff_ptr[:]
	}
}

func FullVerifyNoBuckets() {
	var (
		j        int64
		numProcs = runtime.NumCPU()
	)
	_ = numProcs

	//TODO: implementation of no buckets on full verify

	j = 0
	for i := int64(1); i < num_keys; i++ {
		if key_array[i-1] > key_array[i] {
			j++
		}
	}
	// Com channels
	//
	// ************
	//jReduction := make(chan int64, num_keys-1)
	//for id := int64(0); id < int64(numProcs); id++ {
	//	go func(id int64) {
	//		if id == 0 {
	//			id += int64(numProcs)
	//		}
	//		for i := id; i < num_keys; i += int64(numProcs) {
	//			if key_array[i-1] > key_array[i] {
	//				jReduction <- 1
	//			} else {
	//				jReduction <- 0
	//			}
	//		}
	//	}(id)
	//}
	//for i := int64(1); i < num_keys; i++ {
	//	j += <-jReduction
	//}
	//
	// **********************
	//
	// Com função ParallelFor
	//
	// **********************
	//npb.ParallelFor(
	//	num_keys-1,
	//	int64(numProcs),
	//	func(_ int64, i int64) {
	//		if key_array[i] > key_array[i+1] {
	//			atomic.AddInt64(&j, 1)
	//		}
	//	},
	//)
	if j != 0 {
		fmt.Println("Full_verify: number of keys out of sort: ", j)
	} else {
		passed_verification += 1
	}
}
