
// demo of sorting for int and strings

package main
import "fmt"
import "sort"

// -------------------------------------------------------------
func main() {
    ss := []string{"cc", "aa", "bb"}
    fmt.Println("ss orig :", ss)

    sort.Strings(ss)
    fmt.Println("ss asc  :", ss)

    sort.Sort(sort.Reverse(sort.StringSlice(ss)))
    fmt.Println("ss desc :", ss)

    fmt.Println("---------------")


    // ascending sort
    ints := []int{7, 2, 4}
    fmt.Println("ints orig : ", ints)
    sort.Ints(ints)
    fmt.Println("ints asc  : ", ints)
    sort.Sort(sort.Reverse(sort.IntSlice(ints)))
    fmt.Println("ints desc : ", ints)

}
