package commons

import "sync"

func ParallelFor(n int64, step int64, group *sync.WaitGroup, f func(id int64, i int64, group *sync.WaitGroup)) {
	(*group).Add(int(n))
	for it := int64(0); it < n; it += step {
		limit := step
		if it+step > n {
			limit = n - it
		}
		for myid := int64(0); myid < limit; myid++ {
			go f(myid, it+myid, group)
		}
	}
	(*group).Wait()
}
