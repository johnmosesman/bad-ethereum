package keys

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"golang.org/x/crypto/sha3"
)

func PublicAddress(secret string) string {
	pk := new(ecdsa.PrivateKey)
	pk.D = hexToBigInt((secret))

	curve := secp256k1.S256()

	// Setting the curve and using that for the ellipic functions doesn't work, as it thinks it's
	// of type elliptic. Referencing the curve directly works.
	// pk.PublicKey.Curve = curve

	pk.PublicKey.X, pk.PublicKey.Y = curve.ScalarBaseMult(pk.D.Bytes())

	combinedPubKey := curve.Marshal(pk.PublicKey.X, pk.PublicKey.Y)

	// Remove the leading 04 signifying uncompressed point
	strippedPubKey := combinedPubKey[1:]

	keccak := sha3.NewLegacyKeccak256()
	keccak.Write(strippedPubKey)
	hash := keccak.Sum(nil)

	publicAddr := hash[len(hash)-20:]

	// Add 0x to signify Ethereum address
	fmtPublicAddr := fmt.Sprintf("0x%s", hex.EncodeToString(publicAddr))

	fmt.Println("fmtPubAddr", fmtPublicAddr)

	return fmtPublicAddr
}

func PrivateKey() string {
	// 2^256 - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(256), nil).Sub(max, big.NewInt(1))

	fmt.Println(max)

	v, err := rand.Int(rand.Reader, max)

	if err != nil {
		// stuff here?
	}

	pk := hex.EncodeToString(v.Bytes())

	fmt.Println("Private key:", pk)

	return pk
}

func hexToBigInt(str string) *big.Int {
	i := new(big.Int)
	i.SetString(str, 16)

	return i
}
