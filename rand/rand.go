package rand

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
)

// ErrNotEnoughRead occurs when not enough bytes were read from crypto/rand.Read().
var ErrNotEnoughRead = errors.New("didn't read enough random bytes")

// Bytes creates a slice of random bytes.
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	nRead, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("bytes: %w", err)
	}
	if nRead < n {
		return nil, fmt.Errorf("bytes: %w", ErrNotEnoughRead)
	}
	return b, nil
}

// String returns a random base64 encoded string using crypto/rand.
// n is the number of bytes being used to generate the random string.
func String(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", fmt.Errorf("string: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
