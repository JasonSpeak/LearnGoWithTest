package pointanderror

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBlance := func(t *testing.T, wallet Wallet, want BitCoin) {
		t.Helper()
		if wallet.blance != want {
			t.Errorf("want %s but get %s", want, wallet.blance)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Errorf("wanted an error but didn't get one")
		}
		if got.Error() != want.Error() {
			t.Errorf("got %s but want %s", got.Error(), want.Error())
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		want := BitCoin(10)
		assertBlance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{BitCoin(20)}
		err := wallet.Withdraw(BitCoin(30))
		want := BitCoin(20)
		assertBlance(t, wallet, want)
		assertError(t, err, InsufficientFundsError)
	})

}
