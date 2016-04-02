package clock

import (
	"fmt"
)

const testVersion = 3

type Clock struct {
	Hour   int
	Minute int
}

func New(hour, minute int) Clock {
	for minute > 59 {
		hour += 1
		minute -= 60
	}
	for minute < 0 {
		hour -= 1
		minute += 60
	}
	for hour > 23 {
		hour -= 24
	}
	for hour < 0 {
		hour += 24
	}
	return Clock{hour, minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%0.2d:%0.2d", c.Hour, c.Minute)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.Hour, c.Minute+minutes)
}
