package models

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/gsom95/go-web-dev/rand"
)

const (
	// The minimum number of bytes to be used for each session token.
	MinBytesPerToken = 32
)

type TokenManager struct {
	// BytesPerToken is used to determine how many bytes to use when generating
	// each session token. If this value is not set or is less than the
	// MinBytesPerToken const it will be ignored and MinBytesPerToken will be
	// used.
	BytesPerToken int
}

// NewToken handles creation of a token and its hash.
func (tm TokenManager) NewToken() (token, tokenHash string, err error) {
	bytesPerToken := tm.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err = rand.String(bytesPerToken)
	if err != nil {
		return "", "", fmt.Errorf("TokenManager.NewToken: %w", err)
	}
	tokenHash = tm.Hash(token)

	return token, tokenHash, nil
}

// Hash handles taking in a session token as a string and
// returning the hash of that session token.
func (tm TokenManager) Hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	// base64 encode the data into a string
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}
