package common

import (
	"strconv"
)

const DateLayout = "2006-01-02"

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func ToFloat64(val string) (float64, error) {
	s, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return -1, err
	}
	return s, nil
}
