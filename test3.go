package main

import ( 
    "fmt"
)

// -------------------------------------------------------------
func main() {
    pp := []struct{ name string }{{name: "native"}, {name: "cgo"}}
    fmt.Println(pp[0].name)
    fmt.Println(pp[1].name)
}
