package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Etherium(10))
		assertBalance(t, wallet, Etherium(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Etherium(20)}
		err := wallet.Withdraw(Etherium(15))
		assertNoError(t, err)
		assertBalance(t, wallet, Etherium(5))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Etherium(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Etherium(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Etherium) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Received an error but did not want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Wanted an error but did not receive one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
