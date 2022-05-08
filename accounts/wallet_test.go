package accounts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	assert := assert.New(t)
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance().String()

		assert.Equal(want.String(), got)
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdrawl", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(50)}

		err := wallet.Withdrawl(Bitcoin(25))
		expectedBalance := Bitcoin(25)

		assert.Nil(err)
		assertBalance(t, wallet, expectedBalance)
	})

	t.Run("Withdrawal with insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdrawl(Bitcoin(20))

		assertBalance(t, wallet, startingBalance)
		assert.Error(err)
		assert.ErrorIs(err, nsfError)
	})
}
