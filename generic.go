// The generic dilemma: what do you want? 
//    slow programmers ?
//    slow compilers and bloated binaries ?
//    slow execution times ?
//
// Go has no generics (as well as no inheritance, polymorphism, etc.)
//
// arrays
// slices
// maps
// map[string]struct{} can be used as a Set type
//
// len() works with strings, arrays, slices, and maps.
//
// See examples here:
//    https://appliedgo.net/generics/

package main

import (
    "fmt"
)

func main() {
    m := map[string]interface{}{"a":"apple", "b":2}
    for kk := range m {
        fmt.Println(kk)
    }
}


