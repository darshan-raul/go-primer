package point

import (
	//"fmt"
	"testing"
)


func TestWallet(t *testing.T) {

	assertError := func(t testing.TB, got, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
	
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

// ==== iter 1 - simple balance
// func TestWallet(t *testing.T) {
// 	wallet := Wallet{}

// 	wallet.Deposit(10)

// 	got := wallet.Balance()
// 	fmt.Printf("address of balance in test is %v \n", &wallet.balance)
// 	want := 10

// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}

// }
