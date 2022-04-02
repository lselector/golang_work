package main

import "fmt"

// -------------------------------------------------------------
// how to do recover()
//
// https://blog.golang.org/defer-panic-and-recover
// see real-time example of recover in JSON parsing:
//
// https://golang.org/src/encoding/json/decode.go
// using recover in side a defered anonymous function
// using named return values to be able to handle them
// while recovering

// -------------------------------------------------------------
func myfunc(a, b int64) (res int64, err error) {

    defer func() {
        if rr := recover(); rr != nil {
            fmt.Println("in anon:", res, rr)
            err = rr.(error)
        }
    }()

    res = a / b
    return res, err
}

// -------------------------------------------------------------
func main() {
    res, err := myfunc(1, 0)
    fmt.Println("in main:", res, err)
}