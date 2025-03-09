package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// return byte slice, which represents the sha256 hash of the input
func sha2506f(input string) []byte {
	//initializes a new SHA-256 hashing algorithm instance
	algorithm := sha256.New()
	//write method takes []byte as input. So we convert input into []byte
	algorithm.Write([]byte(input))
	//Sum finalizes  the hash computation and returns the hash as a byte slice
	return algorithm.Sum(nil)
}

// we are using base58 encoding instead of the popular base64
// base58 uses 58 characvters instead of 64.
// Used in bitcoin addresses and other cryptocurrencies

// Base58 reduces confusion in character output
// the characters o,0, 1,l are highly confusing in certain fonts and hard to differentiate.
// base58 also removes poncutation characters to prevent confusion for line breakers
// double clicking selects teh whole number as one word
func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

// 2 main building blocks
// hashing initialURL + userID with sha256.
// userID is added to prevent providing similar shortend urls to separaate users in case they request the same link
// derive a big integer nnumber from the hash bytes generated
// finally apply base58 on the derived big integer value and pick the first 8 characters.
func GenerateShortLink(initialLink string, userId string) string {
	urlHashBytes := sha2506f(initialLink + userId)
	//the code below is to derive a big integer from the hash bytes to ensure we have a numerical representation of the SHA-256 hash output
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	//we encode here to make the final string more human friendly and also keep it short only 8 char
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
