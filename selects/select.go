package selects

import (
	"net/http"
	"time"
)

// Racer returns url which responds faster
func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration > bDuration {
		return b
	}
	return a
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)

	return duration
}
