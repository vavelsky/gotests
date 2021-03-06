package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const startCount = 3

// Sleeper allows to put delays
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper is configurabale version of putting delays
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep for configurable sleeper
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Countdown counts from the number to zero and prints Go!
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := startCount; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	// for i := startCount; i > 0; i-- {
	// 	fmt.Fprintln(out, i)
	// }
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
