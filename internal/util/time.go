package util

import "time"

// TimeInterval defines a time interval, determined by start and end hour
type TimeInterval struct {
	StartHour int
	EndHour   int
}

// Contains checks if the given time is within the given interval. The date part is ignored.
func (interval *TimeInterval) Contains(t time.Time) bool {
	h := t.Hour()
	if h >= interval.StartHour && h <= interval.EndHour {
		return true
	}
	return false
}
