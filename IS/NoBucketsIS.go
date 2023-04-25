package IS

import (
	npb "NPB-Golang/commons"
	"fmt"
	"runtime"
	"sync"
)

func RankNoBuckets(iteration int64) {
	var (
		key_buff_ptr, key_buff_ptr2 []int64
		num_procs                   = runtime.NumCPU()
		group                       sync.WaitGroup
		//barrierGroup                []sync.WaitGroup
	)

	key_array[iteration] = iteration
	key_array[iteration+MAX_ITERATIONS] = max_key - iteration
	key_buff_ptr2 = key_array[:]
	key_buff_ptr = key_buff1[:]
	for i := 0; i < TEST_ARRAY_SIZE; i++ {
		partial_verify_vals[i] = key_array[test_index_array[i]]
	}

	for i := 0; i < num_procs; i++ {
		for j := int64(0); j < max_key; j++ {
			key_buff1_aptr[i][j] = 0
		}
	}

	npb.ParallelFor(
		num_keys,
		int64(num_procs),
		func(id int64, i int64) {
			key_buff1_aptr[id][key_buff_ptr2[i]]++
		},
		npb.STATIC,
	)
	group.Add(num_procs)
	for id := 0; id < num_procs; id++ {
		go func(id int64) {
			defer group.Done()
			for i := int64(0); i < max_key-1; i++ {
				key_buff1_aptr[id][i+1] += key_buff1_aptr[id][i]
			}
		}(int64(id))
	}
	group.Wait()
	for k := 1; k < num_procs; k++ {
		npb.ParallelFor(
			max_key,
			int64(num_procs),
			func(id int64, i int64) {
				key_buff_ptr[i] += key_buff1_aptr[k][i]
			},
			npb.STATIC,
		)
	}

	PartialVerify(iteration, key_buff_ptr[:])

	if iteration == MAX_ITERATIONS {
		key_buff_ptr_global = key_buff_ptr[:]
	}
}

func FullVerifyNoBuckets() {
	var (
		numProcs = runtime.NumCPU()
		group    sync.WaitGroup
	)
	npb.ParallelFor(
		num_keys,
		int64(numProcs),
		func(_ int64, i int64) {
			key_buff2[i] = key_array[i]
		},
		npb.STATIC,
	)
	group.Add(numProcs)
	for id := 0; id < numProcs; id++ {
		go func(myid int64) {
			defer group.Done()
			var (
				i, j, k, k1, k2 int64
			)
			j = int64(numProcs)
			j = (max_key + j - 1) / j
			k1 = j * myid
			k2 = k1 + j
			if k2 > max_key {
				k2 = max_key
			}
			for i = 0; i < num_keys; i++ {
				if key_buff2[i] >= k1 && key_buff2[i] < k2 {
					key_buff_ptr_global[key_buff2[i]]--
					k = key_buff_ptr_global[key_buff2[i]]
					key_array[k] = key_buff2[i]
				}
			}
		}(int64(id))
	}
	group.Wait()

	j := 0
	for i := int64(1); i < num_keys; i++ {
		if key_array[i-1] > key_array[i] {
			j++
		}
	}
	if j != 0 {
		fmt.Println("Full_verify: number of keys out of sort: ", j)
	} else {
		passed_verification += 1
	}
}
