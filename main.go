package main

import (
	"fmt"
	"johnmosesman/bad-ethereum/accounts"
	"johnmosesman/bad-ethereum/keys"
)

func main() {
	key := [20]byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	account := accounts.New(key)
	fmt.Println("hey")
	fmt.Println("account nonce is", account.Nonce)
	fmt.Println("account address is", account.Address)

	keys.PublicAddress("f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315")
}
