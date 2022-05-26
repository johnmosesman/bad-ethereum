package keys

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenPrivateKey() []byte {
	// 2^256 - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(256), nil).Sub(max, big.NewInt(1))

	fmt.Println(max)

	v, err := rand.Int(rand.Reader, max)

	if err != nil {
		// stuff here?
	}

	fmt.Println(v)
	fmt.Println(v.Bytes())

	return v.Bytes()
}
