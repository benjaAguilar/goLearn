package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

const write = "write"
const sleep = "sleep"

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("Prints expected value", func(t *testing.T) {

		buff := bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(&buff, spySleeper)

		got := buff.String()
		expect := `3
2
1
Go!`

		if got != expect {
			t.Errorf("got %q want %q", got, expect)
		}
		if spySleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
		}
	})

	t.Run("Print and waits in expected order", func(t *testing.T) {
		spyOp := &SpyCountdownOperations{}
		Countdown(spyOp, spyOp)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spyOp.Calls) {
			t.Errorf("wanted calls %v got %v", want, spyOp.Calls)
		}
	})
}
