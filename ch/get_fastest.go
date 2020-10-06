package ch

import (
	"time"
)

func getFastest(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	done := make(chan interface{})
	go func() {
		defer close(done)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-getFastest(append(channels[3:], done)...):
			}
		}
	}()
	return done
}

// Sleep for a given duration and return a closed channel
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func getFastestDuration(duration [5]time.Duration) time.Duration {
	start := time.Now()
	<-getFastest(
		sig(duration[0]),
		sig(duration[1]),
		sig(duration[2]),
		sig(duration[3]),
		sig(duration[4]),
	)
	end := time.Since(start)

	switch {
	case duration[0] < end:
		return duration[0]
	case duration[1] < end && end < duration[0]:
		return duration[1]
	case duration[2] < end && end < duration[1]:
		return duration[2]
	case duration[3] < end && end < duration[2]:
		return duration[3]
	case duration[4] < end && end < duration[3]:
		return duration[4]
	default:
		return time.Duration(0) // if args are less than 5
	}
}
