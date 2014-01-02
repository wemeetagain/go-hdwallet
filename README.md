Go HD Wallet tools (WIP -- DOES NOT WORK RIGHT NOW!!)
------------------

BIP 32 wallet tools
 - BIP32 - https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki

 - sort of based off https://github.com/vbuterin/pybitcointools/blob/master/pybitcointools/deterministic.py

###Sample Use

        seed,_ := hdwalletutil.Gen_seed(256)
        masterprv := hdwalletutil.Bip32_master_key(seed)
        masterpub := hdwalletutil.Bip32_privtopub(masterprv)
        childprv := hdwalletutil.Bip32_ckd(masterprv,0)
        childpub := hdwalletutil.Bip32_ckd(masterpub,0)

###Dependencies

        go get code.google.com/p/go.crypto/ripemd160
        go get github.com/conformal/btcutil
        go get github.com/mndrix/btcutil
