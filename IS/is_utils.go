package IS

import (
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
