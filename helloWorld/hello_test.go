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
		got := SayHello("", "")
		expect := "Hello World!"
		assertCorrectMessage(t, got, expect)
	})
	t.Run("It returns Hello <string> if a param is given", func(t *testing.T) {
		got := SayHello("Rick!", "")
		expect := "Hello Rick!"

		assertCorrectMessage(t, got, expect)
	})
	t.Run("It can handle spanish languaje", func(t *testing.T) {
		got := SayHello("", "spanish")
		expect := "Hola Mundo!"
		assertCorrectMessage(t, got, expect)
	})
	t.Run("It can handle french languaje", func(t *testing.T) {
		got := SayHello("", "french")
		expect := "Bonjour Monde!"
		assertCorrectMessage(t, got, expect)
	})
	t.Run("It can handle portuguese languaje", func(t *testing.T) {
		got := SayHello("", "portuguese")
		expect := "Olá Mundo!"
		assertCorrectMessage(t, got, expect)
	})
	t.Run("It defaults to english when lang does not match", func(t *testing.T) {
		got := SayHello("Patrick", "")
		expect := "Hello Patrick"
		assertCorrectMessage(t, got, expect)
	})
	t.Run("It defaults to english when lang does not exist", func(t *testing.T) {
		got := SayHello("", "chinese")
		expect := "Hello World!"
		assertCorrectMessage(t, got, expect)
	})
	t.Run("Can handle case sensitive", func(t *testing.T) {
		got := SayHello("", "SpaNISH")
		expect := "Hola Mundo!"
		assertCorrectMessage(t, got, expect)
	})
}
