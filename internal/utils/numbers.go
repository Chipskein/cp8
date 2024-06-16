package utils

import (
	"fmt"
	"math"
)

/*
ExtractDecimalHouses function is used to extract the hundreds, tens and ones from a given integer.
only to be used with 3 digit numbers.
Used to implement Fx33 opcode.
*/
func ExtractDecimalHouses(value int32) (int32, int32, int32, error) {
	value = int32(math.Abs(float64(value)))
	if value > 255 {
		return 0, 0, 0, fmt.Errorf("value %d is greater than 255", value)
	}
	hundreds := value / 100
	tens := (value / 10) % 10
	ones := value % 10
	return hundreds, tens, ones, nil
}
