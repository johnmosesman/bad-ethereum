package accounts

import (
	"testing"
)

func TestNew(t *testing.T) {
	key := [20]byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	account := New(key)

	if account.Address != key {
		t.Error("Address is not passed in address")
	}
	if account.Nonce != 0 {
		t.Error("Nonce is not 0")
	}
	if account.Balance != 0 {
		t.Error("Balance is not 0")
	}
}
