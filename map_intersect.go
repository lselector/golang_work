// demo how to get intersection (common keys) of two maps

package main

import (
    "fmt"
)

func main() {
    m1 := map[string]string{"k1":"v1", "k2":"v2", "k3":"v3"}
    m2 := map[string]string{"k1":"v1", "k2":"v2", "k4":"v4"}
    inter := map[string]bool{}

    for kk, _ := range m1 {
        _, ok := m2[kk]
        if ok {
            inter[kk] = true
        }
    }

    for kk, _ := range inter {
        fmt.Println("Key:", kk)
    }

}

