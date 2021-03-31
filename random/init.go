package random

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"time"
)

var r *rand.Rand

func init() {
	seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err == nil {
		r = rand.New(rand.NewSource(seed.Int64()))
	} else {
		// fallback 到时间戳作为种子
		_, _ = os.Stderr.WriteString(fmt.Sprintf("WARNING: Use timestamp as random seed (%s)\n", err))

		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
}
