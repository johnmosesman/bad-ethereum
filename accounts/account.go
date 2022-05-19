package accounts

type Account struct {
	Address [20]byte
	Nonce   uint
	Balance uint

	// TODO: what type should these be?
	// Code    string
	// Storage string
}

func New(addr [20]byte) Account {
	return Account{
		Address: addr,
		Nonce:   0,
		Balance: 0,
	}
}
