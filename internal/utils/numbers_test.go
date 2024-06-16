package utils

import (
	"errors"
	"testing"
)

type CasesExtractDecimalHouses struct {
	value    int32
	hundreds int32
	tens     int32
	ones     int32
	err      error
}

func TestExtractDecimalHouses(t *testing.T) {
	var test []CasesExtractDecimalHouses = []CasesExtractDecimalHouses{
		{198, 1, 9, 8, nil},
		{255, 2, 5, 5, nil},
		{300, 0, 0, 0, errors.New("value 300 is greater than 255")},
		{0, 0, 0, 0, nil},
		{-200, 2, 0, 0, nil},
		{-241, 2, 4, 1, nil},
		{10, 0, 1, 0, nil},
		{5, 0, 0, 5, nil},
	}
	for _, c := range test {
		hundreds, tens, ones, err := ExtractDecimalHouses(c.value)
		if err != c.err && err.Error() != c.err.Error() {
			t.Errorf("Expected error to be %s, got %s", c.err.Error(), err.Error())
		}
		if hundreds != c.hundreds {
			t.Errorf("Expected hundreds to be %d, got %d", c.hundreds, hundreds)
		}
		if tens != c.tens {
			t.Errorf("Expected tens to be %d, got %d", c.tens, tens)
		}
		if ones != c.ones {
			t.Errorf("Expected ones to be %d, got %d", c.ones, ones)
		}
	}
}
