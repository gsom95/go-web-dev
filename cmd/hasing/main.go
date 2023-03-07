package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	keyForHash := `a-secret-key`

	pass := `My Very Long and secret pass!1`
	h := hmac.New(sha256.New, []byte(keyForHash))

	h.Write([]byte(pass))
	result := h.Sum(nil)
	fmt.Println(hex.EncodeToString(result))
}
