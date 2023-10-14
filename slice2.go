// slice2.go
// test script shows that when you resize a slice beyond its capacity,
// it is being moved to a new location in memory, but pointers
// to the slice or its elements still point to old location.
package main

import (
    "fmt"
)

func main() {
    aa := []int64 {10, 11, 12, 13, 14}
    bb00 := aa 
    bb10 := &(aa[0])
    bb11 := &(aa[1])
    bb12 := &(aa[2])
    bb13 := &(aa[3])

    fmt.Printf("%p %p\n", aa, bb00)
    fmt.Printf("%p, %p, %p, %p\n", &(aa[0]) , &(aa[1]), &(aa[2]), &(aa[3]) )
    fmt.Printf("%p, %p, %p, %p\n", bb10,      bb11,     bb12,     bb13)
    fmt.Printf("%d, %d, %d, %d\n", *bb10, *bb11, *bb12, *bb13)
    fmt.Println("----- doing append moves the slice")
    aa = append(aa ,5,6,7,8,9)
    aa[0] = 77
    fmt.Println("----- Note that aa moved, but bb* didn't change, it still point to old array")
    fmt.Printf("%p %p\n", aa, bb00)
    fmt.Printf("%p, %p, %p, %p\n", &(aa[0]) , &(aa[1]), &(aa[2]), &(aa[3]) )
    fmt.Printf("%d, %d, %d, %d\n",  (aa[0]) ,  (aa[1]),  (aa[2]),  (aa[3]) )
    fmt.Printf("%p, %p, %p, %p\n", bb10,      bb11,     bb12,     bb13)
    fmt.Printf("%d, %d, %d, %d\n", *bb10, *bb11, *bb12, *bb13)

}
