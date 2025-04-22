package clock

import (
	"time"
	"fmt"
)

// Define the Clock type here.
type Clock struct {
	hr int
	min int
}

func New(h, m int) Clock {
	t := time.Date(2006, time.June, 12, h, m, 0, 0, time.UTC)
	return Clock{hr: t.Hour(), min: t.Minute()}
}

func (c Clock) Add(m int) Clock {
	t := time.Date(2006, time.June, 12, c.hr, c.min, 0, 0, time.UTC)
	addT := t.Add(time.Minute * time.Duration(m))
	return New(addT.Hour(), addT.Minute())
}

func (c Clock) Subtract(m int) Clock {
	toAdd := -(time.Minute * time.Duration(m))
	t := time.Date(2006, time.June, 12, c.hr, c.min, 0, 0, time.UTC)
	subT := t.Add(toAdd)
	return New(subT.Hour(), subT.Minute())
}

func (c Clock) String() string {
	if c.hr < 10 {
		if c.min < 10 {
			return fmt.Sprintf("0%d:0%d", c.hr, c.min)
		} else {
			return fmt.Sprintf("0%d:%d", c.hr, c.min)
		}
	} else {
		if c.min < 10 {
			return fmt.Sprintf("%d:0%d", c.hr, c.min)
		} else {
			return fmt.Sprintf("%d:%d", c.hr, c.min)
		}
	}
}
