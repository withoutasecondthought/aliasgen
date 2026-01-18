package aliasgen

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

const (
	defaultLength = 5
	alphabet      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_.~"
)

var ErrGenerateRandomString = errors.New("failed to generate alias from string")

// GenerateRandomString generates a random string of the specified length.
//
// If no length is provided, it defaults to 5.
//
// The string is composed of characters from the defined alphabet.
// alphabet includes lowercase and uppercase letters, digits, and some special characters.
//
// Example usage:
//
//	str := GenerateRandomString() > yNwb.          // generates a random string of length 5
//	str10 := GenerateRandomString(10) > ZFd~XLbzQM // generates a random string of length 10
func GenerateRandomString(l ...int) (string, error) {
	length := defaultLength
	if len(l) > 0 && l[0] > 0 {
		length = l[0]
	}

	result := make([]byte, length)
	maxIndex := big.NewInt(int64(len(alphabet)))

	for i := range result {
		index, err := rand.Int(rand.Reader, maxIndex)
		if err != nil {
			return "", fmt.Errorf("%w: %w", ErrGenerateRandomString, err)
		}

		result[i] = alphabet[index.Int64()]

	}

	return string(result), nil
}

// MustGenerateRandomString wraps GenerateRandomString and panics if an error occurs.
func MustGenerateRandomString(l ...int) string {
	str, err := GenerateRandomString(l...)
	if err != nil {
		panic(err)
	}

	return str
}
