package IS

import (
	npb "NPB-Golang/commons"
	"fmt"
	"os"
)

func getNPBClass(class string) {
	switch class {
	case "S":
		total_keys_log_2 = 16
		max_key_log_2 = 11
		num_buckets_log_2 = 9
	case "W":
		total_keys_log_2 = 20
		max_key_log_2 = 16
		num_buckets_log_2 = 10
	case "A":
		total_keys_log_2 = 23
		max_key_log_2 = 19
		num_buckets_log_2 = 10
	case "B":
		total_keys_log_2 = 25
		max_key_log_2 = 21
		num_buckets_log_2 = 10
	case "C":
		total_keys_log_2 = 27
		max_key_log_2 = 23
		num_buckets_log_2 = 10
	case "D":
		total_keys_log_2 = 31
		max_key_log_2 = 27
		num_buckets_log_2 = 10
	default:
		fmt.Println("Incorrect class argument")
		os.Exit(1)
	}

	total_keys = 1 << total_keys_log_2
	max_key = 1 << max_key_log_2
	num_buckets = 1 << num_buckets_log_2
	num_keys = total_keys
	size_of_buffers = num_keys
}

func initializeVerificationArrays(class string) {
	for i := 0; i < TEST_ARRAY_SIZE; i++ {
		switch class {
		case "S":
			test_index_array[i] = S_test_index_array[i]
			test_rank_array[i] = S_test_rank_array[i]
		case "A":
			test_index_array[i] = A_test_index_array[i]
			test_rank_array[i] = A_test_rank_array[i]
		case "W":
			test_index_array[i] = W_test_index_array[i]
			test_rank_array[i] = W_test_rank_array[i]
		case "B":
			test_index_array[i] = B_test_index_array[i]
			test_rank_array[i] = B_test_rank_array[i]
		case "C":
			test_index_array[i] = C_test_index_array[i]
			test_rank_array[i] = C_test_rank_array[i]
		case "D":
			test_index_array[i] = D_test_index_array[i]
			test_rank_array[i] = D_test_rank_array[i]
		}

	}
}

func FindMySeed(
	kn int,
	np int,
	nn int32,
	s float64,
	a float64,
) float64 {
	var (
		t1, t2         float64
		mq, nq, kk, ik int32
	)
	if kn == 0 {
		return s
	}
	mq = (nn/4 + int32(np) - 1) / int32(np)
	nq = mq * 4 * int32(kn)

	t1 = s
	t2 = a
	kk = nq

	for kk > 1 {
		ik = kk / 2
		if 2*ik == kk {
			npb.Randlc(&t2, t2)
			kk = ik
		} else {
			npb.Randlc(&t1, t2)
			kk--
		}
	}
	npb.Randlc(&t1, t2)

	return t1
}
