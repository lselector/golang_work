
// https://www.hackerrank.com/challenges/between-two-sets/problem
// given two int arrays "a" and "b"
// find number of int values "x" such as:
//    1. x is divisible by all elements in "a"
//    2. all elements in "b" are divisible by "x"

// ---------------------------------------------------
// See discussion here:
//   https://www.hackerrank.com/challenges/between-two-sets/forum
//
// GCD = greatest common divisor:
//   the largest positive integer that 
//   divides each of the integers. 
//   For example, the gcd of 8 and 12 is 4."
// LCM = least common multiple:
//   the smallest positive integer that is 
//   divisible by each of the integers.

// Possible solution logic (Why it works ??):
// Step1. find LCM of all integers in array "a"
// Step2. find GCD of all integers in array "b"
// Step3: find number of multiples of LCM that
//        evenly divide the GCD
// ---------------------------------------------------

package main

import (
    "fmt"
)

// ---------------------------------------------------
// accepts slice of positive integers
// returns min and max values of this slice
// ---------------------------------------------------
func slice_min_max(arr []int32) (int32, int32) {
    var amin int32 = 0
    var amax int32 = 0
    if len(arr) <= 0 {
        return amin,amax 
    }
    amin = arr[0]
    amax = arr[0]
    for _, elem := range arr {
        if elem > amax {
            amax = elem
        }
        if elem < amin {
            amin = elem
        }
    }
    return amin,amax
}

// ---------------------------------------------------
// Each "x" should have the following properties
// x >= max(a)
// x <= min(b)
// for each a: x % a == 0
// for each b: b % x == 0
// ---------------------------------------------------
func getTotalX(a []int32, b []int32) int32 {
    var counter int32 = 0
    if (len(a) <= 0) || (len(b) <= 0) {
        return counter
    }
    _     , max_a := slice_min_max(a)
    min_b , _     := slice_min_max(b)
    // go through a list of all multiples of max_a
    // between max_a and min_b
    n1 := int32(1+min_b/max_a)
    for ii:=int32(0);ii<n1;ii++ {
        val := max_a*(ii+1)
        flag_a := true
        flag_b := true
        // test that it is divisible by all "a"
        for _, aa := range a {
            if val % aa != 0 {
                flag_a = false
                break
            }
        }
        if flag_a == false {
            continue
        }
        // test that it divides all "b"
        for _, bb := range b {
            if bb % val != 0 {
                flag_b = false
                break
            }
        }
        if flag_b == false {
            continue
        }
        // if we are here, both flags are good
        counter += 1
    }
    return counter
}

func main() {
    a := []int32 {2,4}
    b := []int32 {16,32,96}
    total := getTotalX(a, b)
    fmt.Println(total)
}

