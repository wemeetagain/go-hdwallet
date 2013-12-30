package main

import (
    //"crypto/elliptic"
    "github.com/mndrix/btcutil"
    )

var (
    //mainnet
    //0488B21E public
    //0488ADE4 private
    //testnet
    //043587CF public
    //4358394 private
    )
    
    
type HDWallet struct {
    vbytes []byte //4 bytes
    depth []byte //1 byte
    fingerprint []byte //4 bytes
    i []byte //4 bytes
    chaincode []byte //32 bytes
    key []byte //33 bytes
}


func raw_bip32_ckd(w HDWallet, i uint32) ([]byte, []byte) {
    curve := btcutil.Secp256k1()
    temp := []byte("")
    return temp, temp
}

func bip32_serialize(w HDWallet) []byte {
}

func bip32_deserialize(data []byte) HDWallet {
}

func raw_bip32_privtopub(w HDWallet) HDWallet {
}

func Bip32_privtopub(data []byte) []byte {
    return bip32_serialize(raw_bip32_privtopub(bip32_deserialize(data)))
}

func Bip32_ckd(data []byte ,i uint32) []byte {
    return bip32_serialize(raw_bip32_ckd(bip32_deserialize(data),i))
}


