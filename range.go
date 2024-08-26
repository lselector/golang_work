// create a range (like in Python)

package main

import "fmt"

// -------------------------------------------------------------
func slice_in_range(start, end, step int) []int {
    if step <= 0 || end < start {
        return []int{}
    }
    s := make([]int, 0, 1+(end-start)/step)
    for start <= end {
        s = append(s, start)
        start += step
    }
    return s
}

// -------------------------------------------------------------
func slice_start_count(start, count, step int) []int {
    s := make([]int, count)
    for i := range s {
        s[i] = start
        start += step
    }
    return s
}

func main() {
    s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Println(s) // [0 1 2 3 4 5 6 7 8 9]

    fmt.Println(slice_in_range(10, 19, 1))  // [10 11 12 13 14 15 16 17 18 19]
    fmt.Println(slice_in_range(10, 28, 2))  // [10 12 14 16 18 20 22 24 26 28]
    fmt.Println(slice_in_range(-10, -1, 1)) // [-10 -9 -8 -7 -6 -5 -4 -3 -2 -1]

    fmt.Println(slice_start_count(10, 10, 1))  // [10 11 12 13 14 15 16 17 18 19]
    fmt.Println(slice_start_count(10, 10, 2))  // [10 12 14 16 18 20 22 24 26 28]
    fmt.Println(slice_start_count(-1, 10, -1)) // [-1 -2 -3 -4 -5 -6 -7 -8 -9 -10]
    fmt.Println(slice_start_count(20, 10, 0))  // [20 20 20 20 20 20 20 20 20 20]
}

