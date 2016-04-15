package hamming

import (
	"errors"
	"strings"
)

const testVersion = 4

// Distance measures the Hamming Distance between two strings.
// Each mismatch character accounts for 1 point.
// Measuring two strings with different length will return errors.
func Distance(a, b string) (int, error) {
	a_slice := strings.Split(a, "")
	b_slice := strings.Split(b, "")
	diff := 0
	if len(a_slice) != len(b_slice) {
		return -1, errors.New("Error")
	}
	for i, char := range a_slice {
		if char != b_slice[i] {
			diff += 1
		}
	}
	return diff, nil
}
