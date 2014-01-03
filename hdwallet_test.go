package hdwalletutil

import (
    "fmt"
    "testing"
    "encoding/hex"
    )

// implements https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#test-vectors

var (
    masterhex1 string = "000102030405060708090a0b0c0d0e0f"
    m_pub1 string = "xpub661MyMwAqRbcFtXgS5sYJABqqG9YLmC4Q1Rdap9gSE8NqtwybGhePY2gZ29ESFjqJoCu1Rupje8YtGqsefD265TMg7usUDFdp6W1EGMcet8"
    m_prv1 string = "xprv9s21ZrQH143K3QTDL4LXw2F7HEK3wJUD2nW2nRk4stbPy6cq3jPPqjiChkVvvNKmPGJxWUtg6LnF5kejMRNNU3TGtRBeJgk33yuGBxrMPHi"
    m_0p_pub1 string = "xpub68Gmy5EdvgibQVfPdqkBBCHxA5htiqg55crXYuXoQRKfDBFA1WEjWgP6LHhwBZeNK1VTsfTFUHCdrfp1bgwQ9xv5ski8PX9rL2dZXvgGDnw"
    m_0p_prv1 string = "xprv9uHRZZhk6KAJC1avXpDAp4MDc3sQKNxDiPvvkX8Br5ngLNv1TxvUxt4cV1rGL5hj6KCesnDYUhd7oWgT11eZG7XnxHrnYeSvkzY7d2bhkJ7"
    m_0p_1_pub1 string = "xpub6ASuArnXKPbfEwhqN6e3mwBcDTgzisQN1wXN9BJcM47sSikHjJf3UFHKkNAWbWMiGj7Wf5uMash7SyYq527Hqck2AxYysAA7xmALppuCkwQ"
    m_0p_1_prv1 string = "xprv9wTYmMFdV23N2TdNG573QoEsfRrWKQgWeibmLntzniatZvR9BmLnvSxqu53Kw1UmYPxLgboyZQaXwTCg8MSY3H2EU4pWcQDnRnrVA1xe8fs"
    m_0p_1_2p_pub1 string = "xpub6D4BDPcP2GT577Vvch3R8wDkScZWzQzMMUm3PWbmWvVJrZwQY4VUNgqFJPMM3No2dFDFGTsxxpG5uJh7n7epu4trkrX7x7DogT5Uv6fcLW5"
    m_0p_1_2p_prv1 string = "xprv9z4pot5VBttmtdRTWfWQmoH1taj2axGVzFqSb8C9xaxKymcFzXBDptWmT7FwuEzG3ryjH4ktypQSAewRiNMjANTtpgP4mLTj34bhnZX7UiM"
    m_0p_1_2p_2_pub1 string = "xpub6FHa3pjLCk84BayeJxFW2SP4XRrFd1JYnxeLeU8EqN3vDfZmbqBqaGJAyiLjTAwm6ZLRQUMv1ZACTj37sR62cfN7fe5JnJ7dh8zL4fiyLHV"
    m_0p_1_2p_2_prv1 string = "xprvA2JDeKCSNNZky6uBCviVfJSKyQ1mDYahRjijr5idH2WwLsEd4Hsb2Tyh8RfQMuPh7f7RtyzTtdrbdqqsunu5Mm3wDvUAKRHSC34sJ7in334"
    m_0p_1_2p_2_1000000000_pub1 string = "xpub6H1LXWLaKsWFhvm6RVpEL9P4KfRZSW7abD2ttkWP3SSQvnyA8FSVqNTEcYFgJS2UaFcxupHiYkro49S8yGasTvXEYBVPamhGW6cFJodrTHy"
    m_0p_1_2p_2_1000000000_prv1 string = "xprvA41z7zogVVwxVSgdKUHDy1SKmdb533PjDz7J6N6mV6uS3ze1ai8FHa8kmHScGpWmj4WggLyQjgPie1rFSruoUihUZREPSL39UNdE3BBDu76"
    masterhex2 string = "fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542"
    m_pub2 string = "xpub661MyMwAqRbcFW31YEwpkMuc5THy2PSt5bDMsktWQcFF8syAmRUapSCGu8ED9W6oDMSgv6Zz8idoc4a6mr8BDzTJY47LJhkJ8UB7WEGuduB"
    m_prv2 string = "xprv9s21ZrQH143K31xYSDQpPDxsXRTUcvj2iNHm5NUtrGiGG5e2DtALGdso3pGz6ssrdK4PFmM8NSpSBHNqPqm55Qn3LqFtT2emdEXVYsCzC2U"
    m_0_pub2 string = "xpub69H7F5d8KSRgmmdJg2KhpAK8SR3DjMwAdkxj3ZuxV27CprR9LgpeyGmXUbC6wb7ERfvrnKZjXoUmmDznezpbZb7ap6r1D3tgFxHmwMkQTPH"
    m_0_prv2 string = "xprv9vHkqa6EV4sPZHYqZznhT2NPtPCjKuDKGY38FBWLvgaDx45zo9WQRUT3dKYnjwih2yJD9mkrocEZXo1ex8G81dwSM1fwqWpWkeS3v86pgKt"
    m_0_2147483647p_pub2 string = "xpub6ASAVgeehLbnwdqV6UKMHVzgqAG8Gr6riv3Fxxpj8ksbH9ebxaEyBLZ85ySDhKiLDBrQSARLq1uNRts8RuJiHjaDMBU4Zn9h8LZNnBC5y4a"
    m_0_2147483647p_prv2 string = "xprv9wSp6B7kry3Vj9m1zSnLvN3xH8RdsPP1Mh7fAaR7aRLcQMKTR2vidYEeEg2mUCTAwCd6vnxVrcjfy2kRgVsFawNzmjuHc2YmYRmagcEPdU9"
    m_0_2147483647p_1_pub2 string = "xpub6DF8uhdarytz3FWdA8TvFSvvAh8dP3283MY7p2V4SeE2wyWmG5mg5EwVvmdMVCQcoNJxGoWaU9DCWh89LojfZ537wTfunKau47EL2dhHKon"
    m_0_2147483647p_1_prv2 string = "xprv9zFnWC6h2cLgpmSA46vutJzBcfJ8yaJGg8cX1e5StJh45BBciYTRXSd25UEPVuesF9yog62tGAQtHjXajPPdbRCHuWS6T8XA2ECKADdw4Ef"
    m_0_2147483647p_1_2147483646p_pub2 string = "xpub6ERApfZwUNrhLCkDtcHTcxd75RbzS1ed54G1LkBUHQVHQKqhMkhgbmJbZRkrgZw4koxb5JaHWkY4ALHY2grBGRjaDMzQLcgJvLJuZZvRcEL"
    m_0_2147483647p_1_2147483646p_prv2 string = "xprvA1RpRA33e1JQ7ifknakTFpgNXPmW2YvmhqLQYMmrj4xJXXWYpDPS3xz7iAxn8L39njGVyuoseXzU6rcxFLJ8HFsTjSyQbLYnMpCqE2VbFWc"
    m_0_2147483647p_1_2147483646p_2_pub2 string = "xpub6FnCn6nSzZAw5Tw7cgR9bi15UV96gLZhjDstkXXxvCLsUXBGXPdSnLFbdpq8p9HmGsApME5hQTZ3emM2rnY5agb9rXpVGyy3bdW6EEgAtqt"
    m_0_2147483647p_1_2147483646p_2_prv2 string = "xprvA2nrNbFZABcdryreWet9Ea4LvTJcGsqrMzxHx98MMrotbir7yrKCEXw7nadnHM8Dq38EGfSh6dqA9QWTyefMLEcBYJUuekgW4BYPJcr9E7j"
    )

func TestCKDPub(t *testing.T) {
    cpub := Bip32_ckd(m_pub2,0)
    fmt.Println(cpub)
    fmt.Println(m_0_pub2)
    w1 := bip32_deserialize(cpub)
    w2 := bip32_deserialize(m_0_pub2)
    fmt.Println(w1.key)
    fmt.Println(w2.key)
}

func TestCKDPrv(t *testing.T) {
    cprv := Bip32_ckd(m_prv2,0)
    if cprv != m_0_prv2 {
        t.Errorf("%s\n%s",cprv,m_0_prv2)
        w1 := bip32_deserialize(cprv)
        w2 := bip32_deserialize(m_0_prv2)
        fmt.Println(w1.key)
        fmt.Println(w2.key)
        fmt.Println(len(cprv))
        fmt.Println(len(m_0_prv2))
    }
}

func TestVector1(t *testing.T) {
    seed, _ := hex.DecodeString(masterhex1)
    masterprv := Bip32_master_key(seed)
    if masterprv != m_prv1 {
        t.Errorf("m private key was %s, should have been %s",masterprv,m_prv1)
    }
    masterpub := Bip32_privtopub(masterprv)
    if masterpub != m_pub1 {
        t.Errorf("m public key was %s, should have been %s",masterpub,m_pub1)
    }
    var i uint32
    i = 0x80000000
    var prv string
    prv = Bip32_ckd(masterprv,i)
    if prv != m_0p_prv1 {
        t.Errorf("m/0' private key was %s, should have been %s",prv,m_0p_prv1)
        //w1 := bip32_deserialize(prv)
        //w2 := bip32_deserialize(m_0p_prv1)
        //fmt.Println(w1.key)
        //fmt.Println(w2.key)
    }
    i = 0x00000001
    prv = Bip32_ckd(m_0p_prv1,1)
    if prv != m_0p_1_prv1 {
        t.Errorf("m/0'/1 private key was %s, should have been %s",prv,m_0p_1_prv1)
    }
}

func TestSerialize(t *testing.T) {
    if m_prv2 != bip32_serialize(bip32_deserialize(m_prv2)) {
        t.Errorf("private key not de/reserializing properly")
    }
    if m_pub2 != bip32_serialize(bip32_deserialize(m_pub2)) {
        t.Errorf("public key not de/reserializing properly")
    }
}
