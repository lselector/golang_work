
// maximum sub-array problem
//  - https://en.wikipedia.org/wiki/Maximum_subarray_problem
// 
// A group of share traders are sitting in a line. 
// Each one has some profit or loss from the day.
// You need to find maximum profit for a subset of traders sitting next to each other.
//    input = [-10,2,3,-8,4,5,6,-2]
//    output = 15
// Note - it should run with complexity better than O(N2).

// Here is the python code with complexity O(N) (by Aleksey Orekhov)
//    input = [2,3,-10,4,5,6,-2]
//    N = len(input)
//    ss = input[0]
//    output = ss
//    
//    for ii in range(1,N):
//        sn = max(ss+input[ii], input[ii])
//        output = max(ss,sn)
//        ss = sn
//    
//    output = max(output, 0)
//    print(output)
// ==================
//    Here is an explanation of this algorithm:
//    
//    We have a series (list) of numbers [a0, a1, a2, a3, ... aN].
//    For each index "i" we can look at all possible series 
//    (starting with 0,1,... i-1  and ending with "i"),
//    and select the biggest sum() among them - call it "si".
//    We will calculate those numbers for all "i"-s, and select the maximum among them.
//    This will be the answer.
//    So how can we calculate those largest sums?
//    It turned out to be a simple process.
//    
//    i=0: There is only one series - [a0], so s0 = a0
//    i=1: There are two series: [a0,a1],[a1].
//           So s1 = max(a0+a1, a1)=max(s0 + a1, a1)
//    i=2: There are 3 series (sums). But we can simplify two 2.
//           We need to find s2 = max ((a0+a1+a2), (a1+a2), (a2))
//           Look at max of first two elements:
//              max((a0+a1+a2), (a1+a2)) = max((a0+a1),a1) + a2 = s1+a2
//           so we simplify to max of 2 elements: s2 = max(s1+a2, a2)
//    
//    We can use same logic going forward.
//    Thus:
//    
//    s0=a0
//    s1=max(s0+a1, a1)
//    s2=max(s1+a2, a2)
//    s3=max(s2+a3, a3)
//    ...
//    sN=max(s(N-1)+aN, aN)
//    
//    So all we need to do is to go once through the series calculating values s1, s2, ...
//    and remembering the maximum value among them.


package main

import "fmt"

// -------------------------------------------------------------
func main() {
    input := []int{2,3,-10,4,5,6,-2}
    N := len(input)
    ss := input[0]
    output := ss

    sn := 0
    for ii:=1; ii<N; ii++ {
        sn = max(ss+input[ii], input[ii])
        output = max(ss,sn)
        ss = sn
    }
    output = max(output, 0)
    fmt.Println(output)

}

// -------------------------------------------------------------
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
