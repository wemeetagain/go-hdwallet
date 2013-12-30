package main

import (
    "crypto/hmac"
    "crypto/sha512"
    "crypto/rand"
    "errors"
    )

func GenSeed(length int) ([]byte, error) {
    b := make([]byte, length)
    if length < 128 {
        return b, errors.New("length must be at least 128 bits")
    }
    _, err := rand.Read(b)
    return b, err
}

func Bip32_master_key(seed []byte) ([]byte,[]byte) {
    key := []byte("Bitcoin seed")
    mac := hmac.New(sha512.New, key)
	mac.Write(seed)
	I := mac.Sum(nil)
	secret := I[:(len(I)/2)]
	chain_code := I[(len(I)/2):]
	return secret, chain_code
}

