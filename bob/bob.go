package bob

import (
	"fmt"
	"regexp"
	"strings"
)

const testVersion = 2

// Hey returns the response from Bob.
func Hey(input string) string {
	str := Compact(input)
	switch {
	case isNothing(str):
		return "Fine. Be that way!"
	case isShouting(str):
		return "Whoa, chill out!"
	case isQuestion(str):
		return "Sure."
	default:
		return "Whatever."
	}
}

// Compact clears blanks and commas for returns the input statement.
func Compact(input string) string {
	str := strings.Replace(input, ",", "", -1)
	str = strings.Replace(str, " ", "", -1)
	return str
}

// isShouting checks if the statement is a shouting.
func isShouting(str string) bool {
	reg := regexp.MustCompile("[0-9:)?]+")
	str = reg.ReplaceAllString(str, "")
	fmt.Println(str)
	if len(str) > 0 {
		return str == strings.ToUpper(str)
	}
	return false
}

// isQuestion checks if the statement is a question.
func isQuestion(str string) bool {
	endChar := str[len(str)-1 : len(str)]
	return endChar == "?"
}

// isNothing checks if the statement is nothing.
func isNothing(str string) bool {
	reg := regexp.MustCompile("[0-9a-zA-Z:]+")
	return !reg.MatchString(str)
}
