// you can use "range" to go over array/slice, map, or string.
// For all of those three cases it returns 
// tuples of two values: index and value
// for array  : index and value
// for map    :   key and value
// for string : index rune (integer code of char)
//   for ii, vv := range array { ... }
//   for ii, vv := range map   { ... }
//   for ii, vv := range str   { ... }

package main

import "fmt"

func main() {
    for ii:=0; ii<5; ii++ {
        fmt.Println(ii)
    }

    for _ , val := range []int{1,2,3,4} {
        fmt.Println(val)
    }

    myarr := []int{1,2,3,4}
    for ii , val := range myarr {
        fmt.Println(ii, val)
    }

    m := map[string]string{"k1":"v1", "k2":"v2", "k3":"v3"}
    for kk, vv := range m {
        fmt.Printf("%s : %s\n", kk, vv)
    }
    // keys only
    for kk := range m {
        fmt.Printf("%s\n", kk)
    }

    ss := "Русский English"
    for ii , ch := range ss {
        fmt.Println(ii, ch, string(ch))
    }


}
