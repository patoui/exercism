package clock

import (
	"fmt"
)

var minutesInAnHour = 60
var hoursInADay = 24
var minutesInADay = 1440

// Define the Clock type here.
type Clock struct {
	Minutes int
}

func New(h, m int) Clock {
	h = h % hoursInADay

	if h < 0 {
		h += hoursInADay
	}

	m += h * minutesInAnHour
	m = m % minutesInADay

	if m < 0 {
		m += minutesInADay
	}

	return Clock{Minutes: m}
}

func (c Clock) Add(m int) Clock {
	m = (c.Minutes + m) % minutesInADay

	if m < 0 {
		m += minutesInADay
	}

	return Clock{Minutes: m}
}

func (c Clock) Subtract(m int) Clock {
	m = (c.Minutes - m) % minutesInADay

	if m < 0 {
		m += minutesInADay
	}

	return Clock{Minutes: m}
}

func (c Clock) String() string {
	return fmt.Sprintf(
		"%02d:%02d",
		c.Minutes/minutesInAnHour,
		c.Minutes%minutesInAnHour,
	)
}
