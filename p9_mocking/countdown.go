package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const FINAL_WORD = "GO!"
const COUNTDOWN_START = 3

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := COUNTDOWN_START; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, FINAL_WORD)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}