package hamming

import "errors"

const testVersion = 4

// Distance measures the Hamming Distance between two strings.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Error")
	}
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff, nil
}
