// fizz implements the fizzbuz logic
package fizz

import (
	"strconv"
)

// FizzBuzz is a type in which you can set all the fizzbuzz parameters
type FizzBuzz struct {
	FizzName string
	BuzzName string
	Fizz     int
	Buzz     int
	Limit    int
}

// Generate returns a list strings with a number or a fizzbuzz string
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
