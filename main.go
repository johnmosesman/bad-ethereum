package main

import (
	"fmt"
	"johnmosesman/bad-ethereum/accounts"
)

func main() {
	key := [20]byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	account := accounts.New(key)
	fmt.Println("hey")
	fmt.Println("account nonce is", account.Nonce)
	fmt.Println("account address is", account.Address)
}
