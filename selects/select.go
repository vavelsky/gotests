package selects

import (
	"net/http"
	"time"
)

// Racer returns url which responds faster
func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuration > bDuration {
		return b
	}
	return a
}
