// Merge two sorted arrays A and B, 
// where A has enough extra space to accommodate 
// all elements of B, into one sorted array

package main
import "fmt"

// -------------------------------------------------------------
func myrange(start, count, step int) []int {
    a := make([]int, count)
    for i := range a {
        a[i] = start + i*step
    }
    return a
}

// -------------------------------------------------------------
func main() {
    N  := 100
    NA := 60
    NB := 40

    A := myrange(0,N ,1)   // only first 60 values matter)
    B := myrange(0,NB,1)   // all 40 values matter
    C := myrange(0,NA,1)

    // step 1: shift first 60 elements of A to the right end
    // 0123456____ => ____0123456
    for ii := range C {
        A[N-1-ii] = A[NA-1-ii]
    }

    // step 2:
    ia := 0
    ib := 0
    ii := 0
    for (ia < NA) && (ib < NB) {
        if A[ia+NB] < B[ib] { // take A
            A[ii] = A[ia+NB]
            ia += 1
        } else {
            A[ii] = B[ib]
            ib += 1
        }
        ii += 1
    }
    
    // add remaining elements
    for ia < NA {
        A[ii] = A[ia+NB]
        ii += 1
        ia += 1
    }

    for ib < NB {
        A[ii] = B[ib]
        ii += 1
        ib += 1
    }

    fmt.Println(A)
}
