
package main
  
import (
"fmt"
"strconv"
)

// -------------------------------------------------------------
func copy_arr(){
    fmt.Println("------- copy array by assignment")
    var a, b [4]int

    a[2] = 42
    b = a             // copy into existing array
    bb := a           // create a new duplicate array
    fmt.Println(a, b, bb)
//    [0 0 42 0] [0 0 42 0] [0 0 42 0]

    // 2D array
    var c, d [3][5]int
    c[1][2] = 314
    d = c
    fmt.Println(c)
    fmt.Println(d)

//    [[0 0 0 0 0] [0 0 314 0 0] [0 0 0 0 0]]
//    [[0 0 0 0 0] [0 0 314 0 0] [0 0 0 0 0]]

}

// -------------------------------------------------------------
func copy_slice(){
    fmt.Println("------- copy slice by copy(dst,src)")
    a := []int{1, 2, 3}  // slice
    b := [3]int{0,0,0}   // array
    fmt.Println("a:", a, ", b:", b, " before copying")  

    // copy slice into array
    copy(b[:], a)   // copy(dst, src) works for min length of dst and src
    fmt.Println("a:", a, ", b:", b, " 'b' is now same as 'a'")

    b[1] = 9
    fmt.Println("a:", a, ", b:", b, " 'b' changed")

    // make a new destination slice
    bb := make([]int, len(a))
    copy(bb, a) // copy a into bb
    bb[1] = 99
    fmt.Println("a:", a, ", bb:", bb, " 'bb' changed")

//    a: [1 2 3] , b: [0 0 0]  before copying
//    a: [1 2 3] , b: [1 2 3]  'b' is now same as 'a'
//    a: [1 2 3] , b: [1 9 3]  'b' changed
//    a: [1 2 3] , bb: [1 99 3]  'bb' changed
}
  
// -------------------------------------------------------------
func copy_map() {
    fmt.Println("------- copy map using for loop")

    map1 := map[string]int {
        "k1":11,
        "k2":22,
    }

    map2 := map[string]int {}

    for kk,vv := range map1 {
        map2[kk] = vv
    }

    for kk,_ := range map1 {
        delete(map1, kk)
    }

    for kk,vv := range map1 {
        fmt.Println("map1:",kk,vv)
    }

    for kk,vv := range map2 {
        fmt.Println("map2:",kk,vv)
    }

//    map2: k2 22
//    map2: k1 11
}
  

// -------------------------------------------------------------
func copy_2dim() {
    fmt.Println("------- copy 2dim slice")

    a := make([][]string, 4)   // destination
    
    b := make([][]string, 4)   // source with 4 rows
    for i := range b {
        b[i] = make([]string, 4)  // row with 4 columns
        for j := range b[i] {
            b[i][j] = strconv.Itoa(i + j)
        }
    }
    
    // manual deep copy
    for i := range b {
        a[i] = make([]string, len(b[i]))
        copy(a[i], b[i])
    }
    
    b[0][0] = "apple"

    fmt.Println("a: ", a)
    fmt.Println("b: ", b)
}

// -------------------------------------------------------------
func main() {
    copy_arr()
    copy_slice()
    copy_map()
    copy_2dim()

}
