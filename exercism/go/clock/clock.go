package clock

import (
	"fmt"
	// "strings"
)

type Clock struct {
	h int
	m int
}

// Define the Clock type here.

func New(h, m int) Clock {
	var c Clock
	leap_hour := 0

	for m < 0 {
		m += 60
		leap_hour++
	}

	if m >= 60 {
		h += m / 60
		m %= 60
	}

	h -= leap_hour
	for h < 0 {
		h += 24
	}

	if h >= 24 {
		h %= 24
	}

	c.h = h
	c.m = m
	return c

	panic("Please implement the New function")
}

func (c Clock) Add(m int) Clock {
	c.m += m
	for c.m > 60 {
		c.m -= 60
		c.h++
	}

	for c.h > 24 {
		c.h -= 24
	}

	if c.h == 24 {
		c.h = 0
	}
	return c
	panic("Please implement the Add function")
}

func (c Clock) Subtract(m int) Clock {

	total_minutes := c.h*60 + c.m

	total_minutes -= m

	c.h = total_minutes / 60
	c.m = total_minutes % 60

	if c.h <= 0 {
		c.h += 23
	}

	if c.h == 24 || c.h == -24 {
		c.h = 0
	}

	if c.m < 0 {
		c.m += 60
	}

	return c

	panic("Please implement the Subtract function")
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
	panic("Please implement the String function")
}
