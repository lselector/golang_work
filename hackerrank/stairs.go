package main

import (
    "fmt"
)

func staircase(n int32) {
    nn := int(n)
    if nn <= 0 {
        return
    }
    ss := make([]int, nn)
    for ii:=0;ii<nn;ii++ {
        n_pounds := ii+1 
        n_spaces := nn-n_pounds
        for jj:=0;jj<n_spaces;jj++ {
            ss[jj] = 32
        }         
        for jj:=n_spaces;jj<nn;jj++ {
            ss[jj] = 35
        }
        for _,ch := range ss {
            fmt.Print(string(ch))
        }
        fmt.Println()
    }
}

func main() {
    n:= int32(6)
    staircase(n)
}

