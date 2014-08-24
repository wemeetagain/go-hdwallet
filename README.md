Go HD Wallet tools
------------------

 - BIP32 - https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki
 - Documentation - http://godoc.org/github.com/WeMeetAgain/go-hdwallet

###Get this library

        go get github.com/WeMeetAgain/go-hdwallet

###Example

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
        childstring, err := hdwallet.StringChild(walletstring,0)
        childaddress, err := hdwallet.StringAddress(childstring)

###Dependencies

        go get code.google.com/p/go.crypto/ripemd160
        go get github.com/conformal/btcutil
        go get github.com/conformal/btcec

###Donate
If you found this useful, consider donating to 15bi481QnYeMXEcS3nUUBXWFq2XqJBFCRQ
