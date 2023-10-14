// create and return anonymous function (closure)

package main

import (
    "fmt"
)

// ---------------------------------------------------
// create and return a function with a given string inside
func f1_greet(name string) func() {
    return func() {
        fmt.Printf("Hello %s\n", name)
    }
}

// ---------------------------------------------------
// create and return a function with accepts an argument
func f2_greet(name string) func(kk int) {
    return func(kk int) {
        fmt.Printf("Hello %s %d\n", name, kk)
    }
}

// ---------------------------------------------------
// create and return a function that returns counter
// which increments after each call to it.
func f2_int_seq() func() int {
    i := 0
    fmt.Println("reset counter")
    return func() int {
        i++
        return i
    }
}

// ---------------------------------------------------
func main() {
    f1 := f1_greet("Crocodile")
    f1()
    f2  := f2_greet("Crocodile")
    f2(22)
    
    f3 := f2_int_seq()
    for ii := 0; ii < 3; ii++ {
        fmt.Println(f3())
    }
    f3 = f2_int_seq()  // reset counter
    for ii := 0; ii < 3; ii++ {
        fmt.Println(f3())
    }
}
