// sort_by_function.go
// Example: sort strings by their length

package main

import "sort"
import "fmt"

// Create a `byLength` type 
// that is just an alias for `[]string` built-in type
type byLength []string

// Implement sort.Interface: Len(), Less(), Swap()
// so we can use the sort.Sort() generic function.

func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

// Now we can do sort.Sort
func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)
}
