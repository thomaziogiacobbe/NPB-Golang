package commons

import (
	"sync"
)

func ParallelFor(
	n int64,
	numCPU int64,
	f func(id int64, i int64),
	_schedule ...string,
) {
	var (
		group    sync.WaitGroup
		schedule = "static"
	)
	if len(_schedule) > 0 {
		schedule = _schedule[0]
	}
	group.Add(int(numCPU))
	switch schedule {
	case "static":
		for myid := int64(0); myid < numCPU; myid++ {
			go func(id int64) {
				for it := id; it < n; it += numCPU {
					f(id, it)
				}
				defer group.Done()
			}(myid)
		}
		break
	case "dynamic":
		var (
			arrayLock sync.Mutex
		)
		indexes := make([]int64, n)
		for it := int64(0); it < n; it++ {
			indexes[it] = it
		}
		for myid := int64(0); myid < numCPU; myid++ {
			go func(id int64) {
				var it int64
				for {
					arrayLock.Lock()
					if len(indexes) == 0 {
						arrayLock.Unlock()
						break
					}
					it = indexes[0]
					indexes = indexes[1:]
					arrayLock.Unlock()
					f(id, it)
				}
				defer group.Done()
			}(myid)
		}
		break
	}
	group.Wait()
}
