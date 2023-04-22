package commons

import (
	"sync"
	"sync/atomic"
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
		var index atomic.Int64
		index.Store(-1)
		for myid := int64(0); myid < numCPU; myid++ {
			go func(id int64) {
				for it := index.Add(1); it < n; it = index.Add(1) {
					f(id, it)
				}
				defer group.Done()
			}(myid)
		}
		break
	}
	group.Wait()
}
