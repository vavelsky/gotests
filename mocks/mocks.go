package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const startCount = 3
const write = "write"
const sleep = "sleep"

// Sleeper allows to put delays
type Sleeper interface {
	Sleep()
}

// SpySleeper stores number of Sleep Calls
type SpySleeper struct {
	Calls int
}

// DefaultSleeper is a default struct for Sleep
type DefaultSleeper struct{}

// CountdownOperationsSpy allows to count the calls for sleep
type CountdownOperationsSpy struct {
	Calls []string
}

// Sleep counts the calls for Sleeper
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// Sleep allows to sleep for 1 second
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Sleep counts calls
func (c *CountdownOperationsSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

// Countdown counts from the number to zero and prints Go!
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := startCount; i > 0; i-- {
		sleeper.Sleep()
	}

	for i := startCount; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
