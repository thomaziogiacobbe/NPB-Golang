package commons

import (
	"sync"
)

func ParallelFor(
	n int64,
	numCPU int64,
	f func(id int64, i int64),
) {
	var group sync.WaitGroup
	group.Add(int(numCPU))
	for myid := int64(0); myid < numCPU; myid++ {
		go func(id int64) {
			for it := id; it < n; it += numCPU {
				f(id, it)
			}
			defer group.Done()
		}(myid)
	}
	group.Wait()
}
