package hdwalletutil

import (
    "crypto/sha256"
    "code.google.com/p/go.crypto/ripemd160"
    "encoding/binary"
    "github.com/mndrix/btcutil"
    "math/big"
    )

//exponentiation by squaring
//http://simple.wikipedia.org/wiki/Exponentiation_by_squaring
func pow(x, n int) int {
    if n == 1 {
        return x
    } else if n % 2 == 0 {
        return pow(x*x,n/2)
    } else {
        return x * pow(x*x,(n-1)/2)
    }
}

func byte_append(a ...[]byte) []byte {
    var b []byte
    for _, i := range a {
        b = append(b,i...)
    }
    return b
}

func hash160(data []byte) []byte {
    sha := sha256.New()
    ripe := ripemd160.New()
    sha.Write(data)
    ripe.Write(sha.Sum(nil))
    return ripe.Sum(nil)
}

func dbl_sha256(data []byte) []byte {
    sha1 := sha256.New()
    sha2 := sha256.New()
    sha1.Write(data)
    sha2.Write(sha1.Sum(nil))
    return sha2.Sum(nil)
}

func privtopub(key []byte) []byte {
    curve := btcutil.Secp256k1()
    return compress(curve.ScalarBaseMult(key))
}

func compress(x, y *big.Int) []byte {
    two := big.NewInt(2)
    rem := two.Mod(y,two).Uint64()
    rem += 2
    b := make([]byte,2)
    binary.BigEndian.PutUint16(b,uint16(rem))
    rest := x.Bytes()
    return append(b[1:],rest...)
}

func add_privkeys(k1, k2 []byte) []byte {
    x1 := big.NewInt(0).SetBytes(k1[1:])
    y1 := big.NewInt(0).SetBytes(k1[:1])
    x2 := big.NewInt(0).SetBytes(k2[1:])
    y2 := big.NewInt(0).SetBytes(k2[:1])
    curve := btcutil.Secp256k1()
    return compress(curve.Add(x1,y1,x2,y2))
}

func add_pubkeys(k1, k2 []byte) []byte {
    x1 := big.NewInt(0).SetBytes(k1[1:])
    y1 := big.NewInt(0).SetBytes(k1[:1])
    x2 := big.NewInt(0).SetBytes(k2[1:])
    y2 := big.NewInt(0).SetBytes(k2[:1])
    curve := btcutil.Secp256k1()
    return compress(curve.Add(x1,y1,x2,y2))
}

func uint32ToByte(i uint32) []byte {
    a := make([]byte, 4)
    binary.BigEndian.PutUint32(a,i)
    return a
}

func uint16ToByte(i uint16) []byte {
    a := make([]byte, 2)
    binary.BigEndian.PutUint16(a,i)
    return a[1:]
}

func byteToUint16(b []byte) uint16 {
    if len(b) == 1 {
        zero := make([]byte,1)
        b = append(zero,b...)
    }
    return binary.BigEndian.Uint16(b)
}
