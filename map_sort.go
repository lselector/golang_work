// demo how to get sorted list of map keys

package main

import (
    "fmt"
    "sort"
)

func main() {
    // populate a map
    m := make(map[string]string)
    m["k1"] = "v1"
    m["k2"] = "v2"
    m["k3"] = "v3"

    // create a slice of keys - and sort it
    var keys []string
    for k := range m {
        keys = append(keys, k)
    }
    sort.Strings(keys)  // in-place sort

    for _, k := range keys {
        fmt.Println("Key:", k, "Value:", m[k])
    }
}

