package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	number, ok := fnb.(FancyNumber)

	if ok {
		result, _ := strconv.Atoi(number.n)
		return result
	}

	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	number := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(number))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	integer, isInteger := i.(int)

	if isInteger {
		return DescribeNumber(float64(integer))
	}

	floating64, isFloating64 := i.(float64)

	if isFloating64 {
		return DescribeNumber(floating64)
	}

	numberBox, isNumberBox := i.(NumberBox)

	if isNumberBox {
		return DescribeNumberBox(numberBox)
	}

	fancyNumberBox, isFancyNumberBox := i.(FancyNumberBox)

	if isFancyNumberBox {
		return DescribeFancyNumberBox(fancyNumberBox)
	}

	return "Return to sender"
}
