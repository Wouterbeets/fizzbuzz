package fizz

import (
	"strconv"
)

type FizzBuzz struct {
	BuzzName string
	FizzName string
	Fizz     int
	Buzz     int
	Limit    int
}

func (f FizzBuzz) Generate() []string {
	s := []string{}
	for i := 1; i <= f.Limit; i++ {
		if i%f.Fizz == 0 && i%f.Buzz == 0 {
			s = append(s, f.FizzName+f.BuzzName)
		} else if i%f.Fizz == 0 {
			s = append(s, f.FizzName)
		} else if i%f.Buzz == 0 {
			s = append(s, f.BuzzName)
		} else {
			s = append(s, strconv.Itoa(i))
		}
	}
	return s
}
