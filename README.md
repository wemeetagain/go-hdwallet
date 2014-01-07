Go HD Wallet tools
------------------

 - BIP32 - https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki

###Get this library

        $ go get github.com/WeMeetAgain/gohdwalletutil

###Example

        package main
        
        import (
            "fmt"
            "github.com/WeMeetAgain/gohdwalletutil"
            )
        
        func main() {
            // Generate a random 256 bit seed
            seed,_ := hdwalletutil.Gen_seed(256)
            
            // Create a master private key
            masterprv := hdwalletutil.Bip32_master_key(seed)
            
            // Convert a private key to public key
            masterpub := hdwalletutil.Bip32_privtopub(masterprv)
            
            // Generate new child key based on private or public key
            childprv := hdwalletutil.Bip32_ckd(masterprv,0)
            childpub := hdwalletutil.Bip32_ckd(masterpub,0)
            
            // Create bitcoin address from public key
            address := hdwalletutil.Bip32_pubtoaddress(childpub)
            
            _ = childprv
            fmt.Println(address)
        }

###Dependencies

        go get code.google.com/p/go.crypto/ripemd160
        go get github.com/conformal/btcutil
        go get github.com/mndrix/btcutil
