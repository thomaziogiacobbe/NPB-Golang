package IS

import (
	npb "NPB-Golang/commons"
	"fmt"
	"sync"
	"time"
)

/************************************************/
/* 	This code does not define INT_TYPE 			*/
/* 	instead it assumes all INT_TYPE as int64 	*/
/************************************************/

const (
	MAX_ITERATIONS  = 10
	TEST_ARRAY_SIZE = 5
)

/* used by full_verify to get */
/* copies of rank info */
var (
	key_buff_ptr_global []int64
	passed_verification int
)

/* #define USE_BUCKETS */
/* USE_BUCKETS is always defined */
/* because of that, it is never declared */
var (
	/* bucket_ptrs is originally threadprivate */
	/* and should be passed as parameter to any go func */
	bucket_ptrs [][]int64
	/* originally a pointer to pointer */
	bucket_size [][]int64
)

/************************************/
/* These are the three main arrays. */
/* See SIZE_OF_BUFFERS def below    */
/************************************/
var (
	key_array           []int64 //size = size_of_buffers
	key_buff1           []int64 //size = max_key
	key_buff2           []int64 //size = size_of_buffers
	partial_verify_vals []int64
)

var (
	total_keys_log_2  int64
	max_key_log_2     int64
	num_buckets_log_2 int64
)

var (
	max_key         int64
	num_buckets     int64
	num_keys        int64
	size_of_buffers int64
)

var total_keys int64

/**********************/
/* Partial verify info */
/**********************/
var (
	test_index_array [TEST_ARRAY_SIZE]int
	test_rank_array  [TEST_ARRAY_SIZE]int

	S_test_index_array = [TEST_ARRAY_SIZE]int{48427, 17148, 23627, 62548, 4431}
	S_test_rank_array  = [TEST_ARRAY_SIZE]int{0, 18, 346, 64917, 65463}

	W_test_index_array = [TEST_ARRAY_SIZE]int{357773, 934767, 875723, 898999, 404505}
	W_test_rank_array  = [TEST_ARRAY_SIZE]int{1249, 11698, 1039987, 1043896, 1048018}

	A_test_index_array = [TEST_ARRAY_SIZE]int{2112377, 662041, 5336171, 3642833, 4250760}
	A_test_rank_array  = [TEST_ARRAY_SIZE]int{104, 17523, 123928, 8288932, 8388264}

	B_test_index_array = [TEST_ARRAY_SIZE]int{41869, 812306, 5102857, 18232239, 26860214}
	B_test_rank_array  = [TEST_ARRAY_SIZE]int{33422937, 10244, 59149, 33135281, 99}

	C_test_index_array = [TEST_ARRAY_SIZE]int{44172927, 72999161, 74326391, 129606274, 21736814}
	C_test_rank_array  = [TEST_ARRAY_SIZE]int{61147, 882988, 266290, 133997595, 133525895}

	D_test_index_array = [TEST_ARRAY_SIZE]int{1317351170, 995930646, 1157283250, 1503301535, 1453734525}
	D_test_rank_array  = [TEST_ARRAY_SIZE]int{1, 36538729, 1978098519, 2145192618, 2147425337}
)

func ExecIS() {
	var (
		n_threads      int
		mops           float64
		groupCreateSec sync.WaitGroup
		tt             time.Duration
	)

	getNPBClass(npb.Class)

	//TODO: verify array allocations
	key_array = make([]int64, size_of_buffers)
	key_buff1 = make([]int64, max_key)
	key_buff2 = make([]int64, size_of_buffers)
	partial_verify_vals = make([]int64, TEST_ARRAY_SIZE, TEST_ARRAY_SIZE)

	initializeVerificationArrays(npb.Class)

	fmt.Println("\n\n NAS Parallel Benchmarks 4.1 Parallel Golang version - IS Benchmark")
	fmt.Println(" Size: ", total_keys, " (class ", npb.Class, ")")
	fmt.Println(" Iterations: ", MAX_ITERATIONS, "\n")

	//TODO: create_seq has 1 counter
	groupCreateSec.Add(n_threads)
	for i := 0; i < n_threads; i++ {
		go CreateSeq(314159265.00, 1220703125.00, i, &groupCreateSec)
	}
	groupCreateSec.Wait()

	bucket_ptrs = make([][]int64, 0, n_threads)
	bucket_size = make([][]int64, 0, n_threads)

	for iter := 0; iter < n_threads; iter++ {
		temp := make([]int64, num_buckets)
		bucket_ptrs = append(bucket_size, temp)
		temp = make([]int64, num_buckets)
		bucket_size = append(bucket_size, temp)
	}

	//parallel for is not needed for array initialization
	//when it's declared, it already initializes with value 0

	//TODO: rank has 1 counter
	Rank(1)
	passed_verification = 0

	if npb.Class != "S" {
		fmt.Println("\n   Iteration")
	}

	start := time.Now()
	for iteration := 1; iteration <= MAX_ITERATIONS; iteration++ {
		if npb.Class != "S" {
			fmt.Println("\t\t", iteration)
		}
		Rank(iteration)
	}
	tt = time.Since(start)

	//TODO: full_verify has 1 counter
	FullVerify()

	if passed_verification != 5*MAX_ITERATIONS+1 {
		passed_verification = 0
	}

	//TODO: print results (values and time)
	npb.Print_results(0,
		MAX_ITERATIONS,
		&tt,
		mops,
		"Keys Ranked",
		passed_verification != 0,
	)
}
