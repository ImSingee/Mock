package random

import "sync/atomic"

var inc int64 = 0

func Increment() int64 {
	return atomic.AddInt64(&inc, 1)
}

func IncrementWithStep(step int64) int64 {
	return atomic.AddInt64(&inc, step)
}

func IncrementWithDelta(delta int64) int64 {
	return Increment() + delta
}

func IncrementWithStepAndDelta(step, delta int64) int64 {
	return IncrementWithStep(step) + delta
}

func IncrementReset() {
	atomic.StoreInt64(&inc, 0)
}
