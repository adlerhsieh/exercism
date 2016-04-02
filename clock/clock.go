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
	cm := minute
	ch := hour
	modifier := 0
	for cm > 59 {
		modifier += 1
		cm -= 60
	}
	for cm < 0 {
		modifier -= 1
		cm += 60
	}
	ch = ch + modifier
	for ch > 23 {
		ch -= 24
	}
	for ch < 0 {
		ch += 24
	}
	return Clock{ch, cm}
}

func (c Clock) String() string {
	h := fmt.Sprintf("%d", c.Hour)
	if len(h) == 1 {
		h = "0" + h
	}
	m := fmt.Sprintf("%d", c.Minute)
	if len(m) == 1 {
		m = "0" + m
	}
	return h + ":" + m
}

func (c Clock) Add(minutes int) Clock {
	ch := c.Hour
	cm := c.Minute + minutes
	return New(ch, cm)
}
