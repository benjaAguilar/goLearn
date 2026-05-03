package main

import "testing"

func assertCorrectMessage(t testing.TB, got, expect string) {
	t.Helper()
	if got != expect {
		t.Errorf(
			"- got: %v \n - expected: %v",
			got, expect,
		)
	}
}

func TestSayHello(t *testing.T) {
	t.Run("It returns Hello World! if no param is given", func(t *testing.T) {
		got := SayHello("")
		expect := "Hello World!"
		assertCorrectMessage(t, got, expect)
	})
	t.Run("It returns Hello <string> if a param is given", func(t *testing.T) {
		got := SayHello("Rick!")
		expect := "Hello Rick!"

		assertCorrectMessage(t, got, expect)
	})
}
