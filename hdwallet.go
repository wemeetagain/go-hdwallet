//based off https://github.com/vbuterin/pybitcointools/blob/master/pybitcointools/deterministic.py

package hdwalletutil

import (
    "bytes"
    "crypto/hmac"
    "crypto/sha512"
    "crypto/rand"
    "encoding/hex"
    "errors"
    "github.com/conformal/btcutil"
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
    //testpubstr = "043587CF"
    //TESTPUBLIC []byte
    //4358394 private
    //testprvstr = "04358394"
    //TESTPRIVATE []byte
    )

func init() {
    PUBLIC,_ = hex.DecodeString(pubstr)
    PRIVATE,_ = hex.DecodeString(prvstr)
    //TESTPUBLIC,_ := hex.DecodeString(testpubstr)
    //TESTPRIVATE,_ := hex.DecodeString(testprvstr)
}

type HDWallet struct {
    vbytes []byte //4 bytes
    depth uint16 //1 byte
    fingerprint []byte //4 bytes
    i []byte //4 bytes
    chaincode []byte //32 bytes
    key []byte //33 bytes
}

func raw_bip32_ckd(w HDWallet, i uint32) HDWallet {
    var priv, pub []byte
    if bytes.Compare(w.vbytes, PRIVATE) == 0 {
        priv = w.key
        pub = privtopub(priv)
    } else {
        pub = w.key
    }

    mac := hmac.New(sha512.New, w.chaincode)
    if i >= uint32(0x80000000) {
        if bytes.Compare(w.vbytes, PUBLIC) == 0 {
            panic("Can't do private derivation on public key!")
        }
        mac.Write(append(priv,uint32ToByte(i)...))
    } else {
        mac.Write(append(pub,uint32ToByte(i)...))
    }
    I := mac.Sum(nil)

    var newkey, fingerprint []byte
    if bytes.Compare(w.vbytes, PRIVATE) == 0 {
        newkey = add_privkeys(I[:32], priv)
        fingerprint = hash160(privtopub(w.key))[:4]
    }
    if bytes.Compare(w.vbytes, PUBLIC) == 0 {
        newkey = add_pubkeys(privtopub(I[:32]), w.key)
        fingerprint = hash160(w.key)[:4]
    }
    return HDWallet{w.vbytes, w.depth +1, fingerprint, uint32ToByte(i), I[32:], newkey}
}

func bip32_serialize(w HDWallet) string {
    depth := uint16ToByte(uint16(w.depth % 256))
    //bindata = vbytes||depth||fingerprint||i||chaincode||key
    bindata := append(w.vbytes,append(depth,append(w.fingerprint,append(w.i,append(w.chaincode,w.key...)...)...)...)...)
    chksum := dbl_sha256(bindata)[:4]
    return btcutil.Base58Encode(append(bindata,chksum...))
}

func bip32_deserialize(data string) HDWallet {
    dbin := btcutil.Base58Decode(data)
    if bytes.Compare(dbl_sha256(dbin[:(len(dbin)-4)])[:4], dbin[(len(dbin)-4):]) != 0 {
        panic("Invalid checksum")
    }
    vbytes := dbin[0:4]
    depth := byteToUint16(dbin[4:5])
    fingerprint := dbin[5:9]
    i := dbin[9:13]
    chaincode := dbin[13:45]
    key := dbin[45:78]
    return HDWallet{vbytes, depth, fingerprint, i, chaincode, key}
}

func raw_bip32_privtopub(w HDWallet) HDWallet {
    return HDWallet{PUBLIC, w.depth, w.fingerprint, w.i, w.chaincode, privtopub(w.key)}
}

func Bip32_privtopub(data string) string {
    return bip32_serialize(raw_bip32_privtopub(bip32_deserialize(data)))
}

func Bip32_ckd(data string ,i uint32) string {
    return bip32_serialize(raw_bip32_ckd(bip32_deserialize(data),i))
}

func Bip32_extract_key(data string) []byte {
    w := bip32_deserialize(data)
    return w.key
}

func Gen_seed(length int) ([]byte, error) {
    b := make([]byte, length)
    if length < 128 {
        return b, errors.New("length must be at least 128 bits")
    }
    _, err := rand.Read(b)
    return b, err
}

func Bip32_master_key(seed []byte) string {
    key := []byte("Bitcoin seed")
    mac := hmac.New(sha512.New, key)
    mac.Write(seed)
    I := mac.Sum(nil)
    secret := I[:len(I)/2]
    chain_code := I[len(I)/2:]
    depth := 0
    i := make([]byte, 4)
    fingerprint := make([]byte, 4)
    zero := make([]byte,1)
    w := HDWallet{PRIVATE,uint16(depth),fingerprint,i,chain_code,append(zero,secret...)}
    return bip32_serialize(w)
}
