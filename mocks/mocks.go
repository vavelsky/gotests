package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const startCount = 3

// Countdown counts from the number to zero and prints Go!
func Countdown(out io.Writer) {
	for i := startCount; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
