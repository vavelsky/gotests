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

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

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
