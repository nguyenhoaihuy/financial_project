package utils

import (
	"strconv"
)

// Convert values to thousands
func convert(value string) int64 {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0 // Default to 0 if conversion fails
	}
	return v / 1000
}