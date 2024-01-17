package pointers

import (
	"testing"
)

func TestWallet(t *testing.T){

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin){
		t.Helper()
		got := wallet.Balance()
		if got != want{
			t.Errorf("got %s bitcoins, want %s bitcoins", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T){
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T){
		wallet := Wallet{Amount: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})
	
}