package IS

var (
	total_keys_log_2  int
	max_key_log_2     int
	num_buckets_log_2 int
	total_keys        int
	max_key           int
	num_buckets       int
	num_keys          int
	size_of_buffers   int
)

const (
	MAX_ITERATIONS  = 10
	TEST_ARRAY_SIZE = 5
)

func ExecIS() {
	var (
		i, iteration int
		timecounter  float64
	)

	//TODO: Initialize the verification arrays if a valid class
	//TODO: IS print

	//TODO: create_seq
	/* Generate random number sequence and subsequent keys on all procs */
	/* ... */
	/* Wait, does the variables keep existing after the parallel block's end? */
	//create_seq(314159265.00 /* Random number gen seed */,
	//	1220703125.00 /* Random number gen mult */);
}
