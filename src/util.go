package main

import (
    "crypto/sha256"
    "code.google.com/p/go.crypto/ripemd160"
    )

//exponentiation by squaring
//http://simple.wikipedia.org/wiki/Exponentiation_by_squaring
func pow(x, n int) int {
    if n ==1 {
        return x
    } else if n % 2 == 0 {
        return pow(x*x,n/2)
    } else {
        return x * pow(x*x,(n-1)/2)
    }
}

func hash160(data []byte) []byte {
}

func privtopub(key []byte) []byte {
}

func compress() {
}

func add_pubkeys() {
}
