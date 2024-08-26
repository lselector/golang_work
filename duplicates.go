
package main
  
import (
"fmt"
"sort"
)
  
func unique(myslice []int) []int {
    hash := make(map[int]bool)
    list := []int{} 
    for _, vv := range myslice {
        if _, ok := hash[vv]; !ok {
            hash[vv] = true
            list = append(list, vv)
        }
    }    
    return list
}
  
func main() {
    orig_slice := []int{1,5,3,6,9,9,4,2,3,1,5}
    fmt.Println(orig_slice) 
    deduped := unique(orig_slice)
    sort.Ints(deduped) // in-place sorting
    fmt.Println(deduped)
}