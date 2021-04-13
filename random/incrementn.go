package random

import (
	"sync"
)

var incmap = map[string]int64{}
var inclock = sync.RWMutex{}

func IncrementNWithStep(ns string, step int64) int64 {
	inclock.Lock()
	defer inclock.Unlock()

	v := incmap[ns] + step
	incmap[ns] = v

	return v
}

func IncrementN(ns string) int64 {
	return IncrementNWithStep(ns, 1)
}

func IncrementNWithDelta(ns string, delta int64) int64 {
	return IncrementN(ns) + delta
}

func IncrementNWithStepAndDelta(ns string, step, delta int64) int64 {
	return IncrementNWithStep(ns, step) + delta
}

func IncrementNReset(ns string) {
	inclock.Lock()
	defer inclock.Unlock()

	delete(incmap, ns)
}
