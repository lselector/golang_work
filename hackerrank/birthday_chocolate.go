package main

//    Chocolate bar is an array of squares with numbers on them:
//       s = [2,2,1,3,2]
//    We need to find a number of segments such that:
//      len(segment) = m (birth month)
//      sum(segment) = d (birth day of the month)

import (
    "fmt"
//    "bufio"
//    "io"
//    "os"
//    "strconv"
//    "strings"
)

// -------------------------------------------------------------
func number_of_ways(n int32, s []int32, d,m int32) int32 {
    var count int32
    if n < m {
        return count
    }
    var mysum int32 = 0 
    for ii:=int32(0);ii<m;ii++ {
        mysum += s[ii]
    }
    if mysum == d {
        count += 1
    }
    var i1 int32 = 0      // index of the 1st elem
    var i2 int32 = m      // index after the last elem
    for i2 < n {
        mysum = mysum - s[i1] + s[i2]
        if mysum == d {
            count += 1
        }
        i1 += 1
        i2 += 1
    }
  return int32(count) 

}

//    // -------------------------------------------------------------
//    func main() {
//        var n int32 = 5
//        var s = []int32 {1,2,1,3,2} 
//        var d, m int32 = 3, 2
//    
//        res := number_of_ways(n,s,d,m)
//        fmt.Println(res)
//    }
//    
//    func main() {
//        reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
//        // 5            // n <= 100 - number of squares
//        // 1 2 1 3 2    // 1<=s[i] <= 5, where 0 <= i <= n
//        // 3 2          // 1 <= d <= 31, 1 <= m <= 12
//    
//        // read 1st line
//        line1 := readLine(reader)
//        line1_int, err := strconv.ParseInt(line1, 10, 32)
//        checkError(err)
//        n := int32(line1_int)
//    
//        // read 2nd line
//        line2 := readLine(reader)
//        line2_slice := strings.Split(line2, " ")
//        var s []int32
//        for ii := 0; ii < int(n); ii++ {
//            tmp_int, err := strconv.ParseInt(line2_slice[ii], 10, 32)
//            checkError(err)
//            elem := int32(tmp_int)
//            s = append(s, elem)
//        }
//    
//        // read 3rd line
//        line3 := readLine(reader)
//        line3_slice := strings.Split(line3, " ")
//    
//        tmp_d, err_d := strconv.ParseInt(line3_slice[0], 10, 32)
//        checkError(err_d)
//        d := int32(tmp_d)
//    
//        tmp_m, err_m := strconv.ParseInt(line3_slice[1], 10, 32)
//        checkError(err_m)
//        m := int32(tmp_m)
//    
//        if (d<1) || (m<1) || (d>31) || (m>12) {
//            panic("Exiting on input error - invalid day or month number")
//        }
//    
//        res := number_of_ways(n,s,d,m)
//    
//        fmt.Println(res)
//    
//    }
//    
