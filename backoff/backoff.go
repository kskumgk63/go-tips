package backoff

import (
	"errors"
	"time"
)

// backoff polls for f() to have been completed, with exponential backoff.
func backoff(initialDelay, factor, maxFactor, timeout time.Duration, f func() (bool, error)) error {
	maxDelay := initialDelay * maxFactor
	finish := time.Now().Add(timeout)
	for delay := initialDelay; time.Now().Before(finish); delay = min(delay*factor, maxDelay) {
		ok, err := f()
		if ok || err != nil {
			return err
		}
		// If we don't have time to try again, stop
		if time.Now().Add(delay).After(finish) {
			break
		}
		time.Sleep(delay)
	}
	return errors.New("timeout")
}

func min(t1, t2 time.Duration) time.Duration {
	if t1 < t2 {
		return t1
	}
	return t2
}
