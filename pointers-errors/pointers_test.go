package errors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		// fmt.Printf("address of balance in test is %v\n", &wallet.balance)

		got := wallet.Balance()
		want := Bitcoin(10)

		assertBalance(t, got, want)
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		assertBalance(t, got, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(5)}
		err := wallet.Withdraw(Bitcoin(10))

		assertError(t, err, InsufficientFundsError.Error())

	})

}

func assertBalance(t *testing.T, got Bitcoin, want Bitcoin) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, err error, want string) {
	t.Helper()

	if err == nil {
		t.Fatalf("Expected an error but didn't receive one")
	}

	if err.Error() != want {
		t.Errorf("got %v want %v", err, want)
	}
}

func assertNoError(t *testing.T, err error)  {
	t.Helper()
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
}