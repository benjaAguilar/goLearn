package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const CountdownStart = 3
const FinalWord = "Go!"

func Countdown(out io.Writer, s Sleeper) {
	for i := CountdownStart; i > 0; i-- {
		fmt.Fprintf(out, "%d\n", i)
		s.Sleep()
	}
	fmt.Fprintf(out, FinalWord)
}

// we cant have a main func if it is not main package xd
func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
