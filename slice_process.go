// example processing of slice and arr in a function
// slice is passed by reference
// array - copied, unless you pass a pointer

package main

import (
    "fmt"
)

func process_clice(mylice []byte, i1,i2 int) (int,int) {
    mylice[3] = 65
    i1 += 1
    i2 += 1
    return i1,i2
}

func process_arr(myarr *[6]byte, i1,i2 int) (int,int) {
    myarr[3] = 65
    i1 += 1
    i2 += 2
    return i1,i2
}

func main() {
    b1 := []byte("ABCDEF")
    b2 := [6]byte {66,67,68,69,70,71}
    i1 := 0
    i2 := 0

    i1, i2 = process_clice (b1, i1, i2)
    fmt.Println(string(b1))
    fmt.Println(i1,i2)

    fmt.Println("-------------")

    i1, i2 = process_arr (&b2, i1, i2)
    fmt.Println(string(b2[:]))
    fmt.Println(i1,i2)
}
