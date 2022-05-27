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

func GenECDSA() (*ecdsa.PrivateKey, error) {
	pk := new(ecdsa.PrivateKey)

	bigI := new(big.Int)

	bigIBytes, _ := hex.DecodeString("f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315")
	bigI.SetBytes(bigIBytes)
	pk.D = bigI // genRandomKey()

	fmt.Println("private key is", pk.D)
	fmt.Println("private key is", hex.EncodeToString(pk.D.Bytes()))

	curve := secp256k1.S256()

	// Setting the curve and using that for the ellipic functions doesn't work, as it thinks it's
	// of type elliptic. Referencing the curve directly works.
	// pk.PublicKey.Curve = curve

	pk.PublicKey.X, pk.PublicKey.Y = curve.ScalarBaseMult(pk.D.Bytes())
	combinedPubKey := curve.Marshal(pk.PublicKey.X, pk.PublicKey.Y)

	fmt.Println("combinedPubKey", combinedPubKey)
	fmt.Println("combinedPubKey hex", hex.EncodeToString(combinedPubKey))

	// Remove the leading 04 signifying uncompressed point
	stripPK := combinedPubKey[1:]
	fmt.Println("stripPK", stripPK)
	fmt.Println("stripPK hex", hex.EncodeToString(stripPK))

	// var hash []byte = stripPK
	kec := sha3.NewLegacyKeccak256()
	kec.Write(stripPK)
	hash := kec.Sum(nil)

	fmt.Println("hash", hash)
	fmt.Println("hash hex", hex.EncodeToString(hash))

	pubAddr := hash[len(hash)-20:]

	fmt.Println("pubAddr", pubAddr)
	fmt.Println("pubAddr hex", hex.EncodeToString(pubAddr))

	return pk, nil

	// fmt.Println("max", max)
	// fmt.Println("pk is", pk.D)
	// fmt.Println("pub is", pk.PublicKey.X, pk.PublicKey.Y)

	// fmt.Println("mar", elliptic.Marshal(elliptic.P256(), pk.PublicKey.X, pk.PublicKey.Y))

	// fmt.Println(hex.EncodeToString(elliptic.Marshal(elliptic.P256(), pk.PublicKey.X, pk.PublicKey.Y)))
	// thing := elliptic.Marshal(elliptic.P256(), pk.PublicKey.X, pk.PublicKey.Y)
	// fmt.Println("mar", hex.EncodeToString(elliptic.Marshal(elliptic.P256(), pk.PublicKey.X, pk.PublicKey.Y)))

	// var x []byte

	// hash := sha3.NewLegacyKeccak256()

}

func genRandomKey() *big.Int {
	// 2^256 - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(256), nil).Sub(max, big.NewInt(1))

	fmt.Println(max)

	v, err := rand.Int(rand.Reader, max)

	if err != nil {
		// stuff here?
	}

	return v
}
