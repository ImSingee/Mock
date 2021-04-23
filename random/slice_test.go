package random

import (
	"fmt"
	"github.com/ImSingee/tt"
	"testing"
)

func TestUnique(t *testing.T) {
	type args struct {
		min   int
		max   int
		count int
	}
	tests := []struct {
		name      string
		args      args
		returnNil bool
	}{
		{"invalid", args{0, 0, 1}, true},
		{"invalid", args{1, 0, 1}, true},
		{"invalid", args{0, 3, 0}, true},
		{"normal", args{0, 10, 5}, false},
		{"normal", args{0, 10, 20}, false},
		{"normal", args{0, 10, 20}, false},
	}
	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			for k := 0; k < 100; k++ {
				got := Unique(c.args.min, c.args.max, c.args.count)
				if c.returnNil {
					tt.AssertIsNil(t, got)
					return
				}

				tt.AssertIsNotNil(t, got)

				fmt.Println(c.args.min, c.args.max, c.args.count)
				fmt.Println(got)

				if c.args.count > c.args.max-c.args.min {
					c.args.count = c.args.max - c.args.min
				}

				tt.AssertEqual(t, c.args.count, len(got))

				last := c.args.min - 1
				for _, v := range got {
					tt.AssertTrue(t, v >= c.args.min)
					tt.AssertTrue(t, v < c.args.max)
					tt.AssertTrue(t, v > last)
					last = v
				}
			}
		})
	}

}
