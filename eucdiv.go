package eucdiv

import (
	"errors"
	"math"
)

type eucdivresult struct {
	Quotient  int
	Remainder int
}

func Eucdiv(dividend, divisor float64) (result eucdivresult, err error) {
	if divisor == 0 {
		return eucdivresult{0, 0}, errors.New("eucdiv: Division by zero is undefined")
	}

	if dividend == 0 {
		return eucdivresult{0, 0}, nil
	}

	var quotient int
	var remainder int
	var proceed bool
	var negative bool

	if dividend < 0 || divisor < 0 {
		negative = true
	}

	if dividend < 0 && divisor < 0 {
		negative = false
	}

	for proceed == true {
		if math.Abs(dividend) < math.Abs(divisor) {
			remainder = int(math.Abs(dividend))
			proceed = false
		} else {
			dividend = math.Abs(dividend) - math.Abs(divisor)
			quotient += 1
		}
	}

	if negative {
		quotient = quotient - 2*quotient
	}

	return eucdivresult{quotient, remainder}, nil
}
