package commons

import (
	"sync"
	"sync/atomic"
)

type Scheduler int

const (
	DYNAMIC Scheduler = iota
	STATIC
)

func ParallelFor(
	n int64,
	numCPU int64,
	f func(id int64, i int64),
	schedulling Scheduler) {
	var (
		group sync.WaitGroup
	)
	group.Add(int(numCPU))
	switch schedulling {
	case STATIC:
		for myid := int64(0); myid < numCPU; myid++ {
			go func(id int64) {
				for it := id; it < n; it += numCPU {
					f(id, it)
				}
				defer group.Done()
			}(myid)
		}
		break
	case DYNAMIC:
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
