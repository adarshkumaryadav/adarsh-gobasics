package pointerscode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalletBalance(t *testing.T) {
	w := &Wallet{}
	t.Run("Balace", func(t *testing.T) {
		want := w.Deposit(10)
		got := w.Balace()
		if got != want {
			t.Errorf("want %d got %d", want, got)
			t.Fail()
		}
	})
	t.Run("0 balance but still trying to withdraw", func(t *testing.T) {
		w := &Wallet{}
		_, err := w.Deduct(5)
		assert.NotNil(t, err)
	})
	t.Run("trying to withdraw", func(t *testing.T) {
		w := &Wallet{}
		w.Deposit(10)
		remaining, err := w.Deduct(5)
		assert.Nil(t, err)
		assert.Equal(t, remaining, 5)
	})
}
