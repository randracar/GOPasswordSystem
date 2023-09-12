package main

import (
	"golang.org/x/crypto/sha3"
	"encoding/hex"
	"log"
)

func hash_string(s string) string {
	hash := sha3.New512()
	_, err := hash.Write([]byte(s))

	if err != nil {
		log.Fatal(err)
	}

	sha3 := hash.Sum(nil)

	es := hex.EncodeToString(sha3)
	return es
}
