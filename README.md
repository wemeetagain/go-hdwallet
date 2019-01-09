Go HD Wallet tools [![Build Status](https://travis-ci.org/wemeetagain/go-hdwallet.svg?branch=master)](https://travis-ci.org/wemeetagain/go-hdwallet)
------------------

 - BIP32 - https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki
 - Documentation - http://godoc.org/github.com/wemeetagain/go-hdwallet

### Get this library

        go get github.com/wemeetagain/go-hdwallet

### Example

        // Generate a random 256 bit seed
        seed, err := hdwallet.GenSeed(256)
        
        // Create a master private key
        masterprv := hdwallet.MasterKey(seed)
        
        // Convert a private key to public key
        masterpub := masterprv.Pub()
        
        // Generate new child key based on private or public key
        childprv, err := masterprv.Child(0)
        childpub, err := masterpub.Child(0)
        
        // Create bitcoin address from public key
        address := childpub.Address()

        // Convenience string -> string Child and ToAddress functions
        walletstring := childpub.String()
        childstring, err := hdwallet.StringChild(walletstring, 0)
        childaddress, err := hdwallet.StringAddress(childstring)

### Dependencies

        go get golang.org/x/crypto/ripemd160
        go get github.com/btcsuite/btcutil/base58
        go get github.com/btcsuite/btcd/btcec

### License

Unlicense
