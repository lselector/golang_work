package main

import ( "fmt"
)



// -------------------------------------------------------------
func main() {
    a := [][]int64{
        {99, 1, 2, 3},
        {4, 5, 6, 7},
    }
    fmt.Println(a)
    
    var b int64
    b = a[0][2]
    fmt.Println(b)

    a[0][2] = 77
    fmt.Println(a)
    b = a[0][2]
    fmt.Println(b)
}
