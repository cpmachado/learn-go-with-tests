package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

// Sleeper allows you to put delays.
type Sleeper interface {
	Sleep()
}

// Countdown prints a countdown from a static countdownStart to w with a
// delay between count provided by Sleeper.
func Countdown(w io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}

	fmt.Fprint(w, finalWord)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep will pause execution for the defined Duration.
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
