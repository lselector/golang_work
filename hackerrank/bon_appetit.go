package main

// https://www.hackerrank.com/challenges/bon-appetit/problem

import (
    "fmt"
    "bufio"
    "io"
    "os"
    "strings"
    "strconv"
)

// ---------------------------------------------------
func check_bill(n, k int64, bill []int64, b int64) {
    // calculate Anna's balance by summing all
    // prices except k-th, then dividing by 2
    var mysum int64

    for ii, bb := range bill {
        if int64(ii) == k {
            continue
        }
        mysum += bb
    }
    mysum = mysum / 2
    if mysum == b {
        fmt.Println("Bon Appetit")
    } else {
        fmt.Println(b-mysum)
    }        
}

// ---------------------------------------------------
func main() {

    fh, err1 := os.Open("input1.txt")
    if err1 != nil {
        fmt.Printf("error opening file: %v\n",err1)
        os.Exit(1)
    }
    reader := bufio.NewReaderSize(fh, 1024 * 1024)
    //reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
    //  input line1: n, k  // n in [2,10e5], k in [0,n)
    //  input line2: n space-separated integers bill[i], i in [0,n)
    //  input line 3: b (int) - money charged to Anna
    //  output: print "Bon Appetit" if all correct, 
    //          print int money to refund to Anna

    line1 := readLine(reader)
    line1_slice := strings.Split(line1, " ")
    var tmp_int int64
    var err error
    // ------------
    tmp_int, err = strconv.ParseInt(line1_slice[0], 10, 32)
    checkError(err)
    n := int64(tmp_int)
    // ------------
    tmp_int, err = strconv.ParseInt(line1_slice[1], 10, 32)
    checkError(err)
    k := int64(tmp_int)
    // ------------
    if (n<2) || (n>100000) || (k<0) || (k > n) {
        panic("ERROR in input line 1: " + line1_slice[0])
    }

    // read 2nd line
    line2 := readLine(reader)
    line2_slice := strings.Split(line2, " ")
    var bill []int64
    var mysum int64
    for ii := 0; ii < int(n); ii++ {
        tmp_int, err := strconv.ParseInt(line2_slice[ii], 10, 32)
        checkError(err)
        elem := int64(tmp_int)
        if elem <0 || elem >= 1e4 {
            panic("ERROR in input line 2: " + line2_slice[ii])
        }
        bill = append(bill, elem)
        mysum += elem
    }

    line3 := readLine(reader)
    tmp_int, err = strconv.ParseInt(line3, 10, 32)
    checkError(err)
    b := int64(tmp_int)
    if b < 0 || b > int64(mysum) {
        panic("ERROR in input line 3: " + line3)
    }
    check_bill(n, k, bill, b)

}

// ---------------------------------------------------
func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

// ---------------------------------------------------
func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
