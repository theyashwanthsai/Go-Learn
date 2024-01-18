package pointers

import (
	"testing"
)

func TestWallet(t *testing.T){
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

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T){
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, InsufficientFundsErr)
		assertBalance(t, wallet, startingBalance)
	})
}


func assertBalance(t testing.TB, wallet Wallet, want Bitcoin){
	t.Helper()
	got := wallet.Balance()
	if got != want{
		t.Errorf("got %s bitcoins, want %s bitcoins", got, want)
	}
}

func assertError(t testing.TB, got error, want error){
	t.Helper()

	if got == nil{
		t.Fatal("Wanted an error but didnt get one")
	}

	if got != want{
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}