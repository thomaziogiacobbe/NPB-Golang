package IS

import (
	npb "NPB-Golang/commons"
	"runtime"
	"sync"
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

func Rank(iteration int) {
	var (
		i, k                        int64
		key_buff_ptr, key_buff_ptr2 []int64
		shift                       = max_key_log_2 - num_buckets_log_2
		num_bucket_keys             = int32(1) << shift
		num_procs                   = runtime.NumCPU()
		group                       sync.WaitGroup
	)
	key_array[iteration] = int64(iteration)
	key_array[iteration+MAX_ITERATIONS] = max_key - int64(iteration)
	key_buff_ptr2 = key_buff2[:]
	key_buff_ptr = key_buff1[:]
	for i := 0; i < TEST_ARRAY_SIZE; i++ {
		partial_verify_vals[i] = key_array[test_index_array[i]]
	}

	group.Add(int(int64(num_procs) * num_keys))
	for myid := 0; myid < num_procs; myid++ {
		for i := int64(0); i < num_keys; i++ {
			go func(myid int, it int64) {
				bucket_size[myid][key_array[it]>>shift]++
				defer group.Done()
			}(myid, i)
		}
	}
	group.Wait()
	group.Add(int(int64(num_procs) * num_buckets))
	for myid := 0; myid < num_procs; myid++ {
		func(bucketPtrs []int64) {
			for i := int64(0); i < num_buckets; i++ {
				go func(myid int, it int64) {
					for k := 0; k < num_procs; k++ {
						if k < myid {
							bucketPtrs[it] += bucket_size[k][it]
						} else {
							bucketPtrs[it] += bucket_size[k][it-1]
						}
					}
					defer group.Done()
				}(myid, i)
			}
		}(bucket_ptrs[:])
		//TODO: return bucket_ptrs to outside this lambda
	}
	group.Wait()
	group.Add(int(num_buckets))
	for it := int64(1); it < num_buckets; it++ {
		go func(it int64) {
			//TODO: update bucket_ptrs outside this lambda
			bucket_ptrs[it] += bucket_ptrs[it-1]
			defer group.Done()
		}(it)
	}
	group.Wait()
	group.Add(int(num_keys))
	for it := int64(0); it < num_keys; it++ {
		go func(it int64) {
			k = key_array[it]
			//TODO: get bucket_ptrs from outside this lambda
			key_buff2[bucket_ptrs[k>>shift]] = k
			bucket_ptrs[k>>shift]++
			defer group.Done()
		}(it)
	}
	group.Wait()
	for myid := 0; myid < num_procs; myid++ {
		if myid < num_procs-1 {
			group.Add(int(num_buckets))
			for i := int64(0); i < num_buckets; i++ {
				go func(myid int, it int64) {
					for k := myid + 1; k < num_procs; k++ {
						//TODO: update bucket_ptrs outside this lambda
						bucket_ptrs[it] += bucket_size[k][it]
					}
					defer group.Done()
				}(myid, i)
			}
		}
	}
	group.Wait()
	group.Add(int(num_buckets))
	for it := int64(0); it < num_buckets; it++ {
		go func(i int64) {
			var (
				k1, k2 int32
				m      int64
			)
			k1 = i * num_bucket_keys
			k2 = k1 + num_bucket_keys
			for k := k1; k < k2; k++ {
				key_buff_ptr[k] = 0
			}
			if i > 0 {
				//TODO: get bucket_ptrs from outside this lambda
				m = bucket_ptrs[i-1]
			} else {
				m = 0
			}
			for k = m; k < bucket_ptrs[i]; k++ {
				key_buff_ptr[key_buff_ptr2[k]]++
			}
			key_buff_ptr[k1] += m
			for k := k1 + 1; k < k2; k++ {
				key_buff_ptr[k] += key_buff_ptr[k-1]
			}
			defer group.Done()
		}(i)
	}
	group.Wait()

	//TODO: partial verify
}
