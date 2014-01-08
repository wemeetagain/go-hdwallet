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

func rawBip32Ckd(w HDWallet, i uint32) HDWallet {
    var priv, pub []byte
    if bytes.Compare(w.vbytes, PRIVATE) == 0 {
        priv = w.key
        pub = privToPub(priv)
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
        newkey = addPrivKeys(I[:32], priv)
        fingerprint = hash160(privToPub(w.key))[:4]
    }
    if bytes.Compare(w.vbytes, PUBLIC) == 0 {
        newkey = addPubKeys(privToPub(I[:32]), w.key)
        fingerprint = hash160(w.key)[:4]
    }
    return HDWallet{w.vbytes, w.depth +1, fingerprint, uint32ToByte(i), I[32:], newkey}
}

func bip32Serialize(w HDWallet) string {
    depth := uint16ToByte(uint16(w.depth % 256))
    //bindata = vbytes||depth||fingerprint||i||chaincode||key
    bindata := append(w.vbytes,append(depth,append(w.fingerprint,append(w.i,append(w.chaincode,w.key...)...)...)...)...)
    chksum := dblSha256(bindata)[:4]
    return btcutil.Base58Encode(append(bindata,chksum...))
}

func bip32Deserialize(data string) HDWallet {
    dbin := btcutil.Base58Decode(data)
    if bytes.Compare(dblSha256(dbin[:(len(dbin)-4)])[:4], dbin[(len(dbin)-4):]) != 0 {
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

func rawBip32PrivToPub(w HDWallet) HDWallet {
    return HDWallet{PUBLIC, w.depth, w.fingerprint, w.i, w.chaincode, privToPub(w.key)}
}

func Bip32PrivToPub(data string) string {
    return bip32Serialize(rawBip32PrivToPub(bip32Deserialize(data)))
}

func Bip32Ckd(data string ,i uint32) string {
    return bip32Serialize(rawBip32Ckd(bip32Deserialize(data),i))
}

func Bip32ExtractKey(data string) []byte {
    w := bip32Deserialize(data)
    return w.key
}

func Bip32PubToAddress(data string) string {
    x, y := expand(Bip32ExtractKey(data))
    four,_ := hex.DecodeString("04")
    padded_key := append(four,append(x.Bytes(),y.Bytes()...)...)
    zero,_ := hex.DecodeString("00")
    addr_1 := append(zero,hash160(padded_key)...)
    chksum := dblSha256(addr_1)
    return btcutil.Base58Encode(append(addr_1,chksum[:4]...))
}

func GenSeed(length int) ([]byte, error) {
    b := make([]byte, length)
    if length < 128 {
        return b, errors.New("length must be at least 128 bits")
    }
    _, err := rand.Read(b)
    return b, err
}

func Bip32MasterKey(seed []byte) string {
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
    return bip32Serialize(w)
}

func Bip32IsValidKey(key string) bool {
    dbin := btcutil.Base58Decode(key)
    if len(dbin) < 78 || len(dbin) > 82 {
        return false
    }
    // check for correct public or private vbytes
    public,_ := hex.DecodeString("0488B21E")
    private,_ :=  hex.DecodeString("0488ADE4")
    if bytes.Compare(dbin[:4],public) != 0 && bytes.Compare(dbin[:4],private) != 0 {
        return false
    }
    // if public, check x coord is on curve
    x, y := expand(dbin[45:78])
    if bytes.Compare(dbin[:4],public) == 0 && !onCurve(x,y) {
        return false
    }
    return true
}
