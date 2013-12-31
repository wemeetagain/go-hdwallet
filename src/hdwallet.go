package main

import (
    //"crypto/elliptic"
    "crypto/hmac"
    "crypto/sha512"
    "crypto/rand"
    "encoding/hex"
    "errors"
    "github.com/mndrix/btcutil"
    )

var (
    //mainnet
    //0488B21E public
    pubstr = "0488B21E"
    PUBLIC []byte
    //0488ADE4 private
    prvstr = "0488ADE4"
    PRIVATE []byte
    //testnet
    //043587CF public
    testpubstr = "043587CF"
    TESTPUBLIC []byte
    //4358394 private
    testprvstr = "04358394"
    TESTPRIVATE []byte
    )

func init() {
    PUBLIC, err := hex.DecodeString(pubstr)
    if err != nil {
        panic(err)
    }
    PRIVATE, err := hex.DecodeString(prvstr)
    if err != nil {
        panic(err)
    }
    TESTPUBLIC, err := hex.DecodeString(testpubstr)
    if err != nil {
        panic(err)
    }
    TESTPRIVATE, err := hex.DecodeString(testprvstr)
    if err != nil {
        panic(err)
    }
}

type HDWallet struct {
    vbytes []byte //4 bytes
    depth []byte //1 byte
    fingerprint []byte //4 bytes
    i []byte //4 bytes
    chaincode []byte //32 bytes
    key []byte //33 bytes
}

func raw_bip32_ckd(w HDWallet, i uint32) HDWallet {
    var priv, pub []byte
    if w.vbytes == PRIVATE {
        priv = w.key
    } else {
        pub = w.key
    }

    mac := hmac.New(sha512.New, w.chaincode)
    if i >= pow(2,31) {
        if w.vbytes == PUBLIC {
            panic("Can't do private derivation on public key!")
        }
        zero,_ := hex.DecodeString("00")
        mac.Write(append(zero,priv[:32],i...))
    } else {
        mac.Write()
    }
    I := mac.Sum(nil)
    
    var newkey, fingerprint []byte
    if w.vbytes == PRIVATE {
        newkey = 
        fingerprint = 
    }
    if w.vbytes == PUBLIC {
        newkey = 
        fingerprint = 
    }

    return HDWallet{w.vbytes, depth +1, fingerprint, i, I[32:], newkey}
}

func bip32_serialize(w HDWallet) []byte {
}

func bip32_deserialize(data []byte) HDWallet {
}

func raw_bip32_privtopub(w HDWallet) HDWallet {
    return HDWallet{PUBLIC, w.depth, w.fingerprint, w.i, w.chaincode, privtopub(w.key)}
}

func Bip32_privtopub(data []byte) []byte {
    return bip32_serialize(raw_bip32_privtopub(bip32_deserialize(data)))
}

func Bip32_ckd(data []byte ,i uint32) []byte {
    return bip32_serialize(raw_bip32_ckd(bip32_deserialize(data),i))
}

func Bip32_extract_key(data []byte) []byte {
}

func Gen_seed(length int) ([]byte, error) {
    b := make([]byte, length)
    if length < 128 {
        return b, errors.New("length must be at least 128 bits")
    }
    _, err := rand.Read(b)
    return b, err
}

func Bip32_master_key(seed []byte) []byte {
    key := []byte("Bitcoin seed")
    mac := hmac.New(sha512.New, key)
    mac.Write(seed)
    I := mac.Sum(nil)
    secret := I[:len(I)/2]
    chain_code := I[len(I)/2:]

}
