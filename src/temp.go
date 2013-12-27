package main

import (
    "fmt"
    )
    
func main() {
    seed, err := GenSeed(128)
    if err != nil {
        panic(err)
    }
    s, c := GenMasterKey(seed)
    fmt.Println(len(s))
    fmt.Println(len(c))
}
