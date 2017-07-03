package fizz

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		fizzStr string
		buzzStr string
		fizzVal int
		buzzVal int
		limit   int
		want    []string
	}{
		{
			"fizz",
			"buzz",
			2,
			3,
			7,
			[]string{
				"1",
				"fizz",
				"buzz",
				"fizz",
				"5",
				"fizzbuzz",
				"7",
			},
		},
	}
	for _, test := range tests {
		fb := &FizzBuzz{
			test.fizzStr,
			test.buzzStr,
			test.fizzVal,
			test.buzzVal,
			test.limit,
		}
		res := fb.Generate()
		for i, line := range res {
			if line != test.want[i] {
				fmt.Println(test.want[i], line)
				t.Fail()
			}
		}
	}
}
