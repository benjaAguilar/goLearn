package pointerserrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.GetBalance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, err error, want error) {
		t.Helper()
		if err == nil {
			t.Error("Wanted an error but received nil")
		}

		if err != want {
			t.Errorf("got %q, want %q", err, want)
		}
	}

	t.Run("Balance", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(15.2))
		expect := Bitcoin(15.2)

		assertBalance(t, wallet, expect)
	})

	t.Run("whitdraw", func(t *testing.T) {
		wallet := Wallet{Balance: 20.0}

		err := wallet.Whitdraw(Bitcoin(13.6))

		expect := Bitcoin(6.4)

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, expect)
	})

	t.Run("whitdraw insufficient founds", func(t *testing.T) {
		wallet := Wallet{Balance: Bitcoin(5)}

		err := wallet.Whitdraw(Bitcoin(33))
		expect := Bitcoin(5)

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, expect)
	})
}
