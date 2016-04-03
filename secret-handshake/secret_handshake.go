package secret

import (
	"fmt"
	"strings"
)

func Handshake(num int) []string {
	if num <= 0 {
		return nil
	}
	binary := fmt.Sprintf("%b", num)
	slice := strings.Split(binary, "")
	slice = reverseSlice(slice)
	result := []string{}
	for i := 0; i < len(slice); i++ {
		if slice[i] == "0" {
			continue
		}
		switch i {
		case 0:
			result = append(result, "wink")
		case 1:
			result = append(result, "double blink")
		case 2:
			result = append(result, "close your eyes")
		case 3:
			result = append(result, "jump")
		case 4:
			result = reverseSlice(result)
		}
	}
	return result
}

func reverseSlice(slice []string) []string {
	reverse_slice := []string{}
	for c := len(slice) - 1; c >= 0; c-- {
		reverse_slice = append(reverse_slice, slice[c])
	}
	return reverse_slice
}
