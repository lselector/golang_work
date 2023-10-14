package main


import (
    "fmt"
//    "bufio"
//    "io"
//    "os"
//    "strconv"
//    "strings"
)

// -------------------------------------------------------------
func number_of_pairs(n,k int32, a []int32) int32 {
    var count int32
    for i,ai := range a[0:n-1] {
        for _, aj := range a[i+1:n] {
            mysum :=  ai + aj
            if mysum % k == 0 {
                count += 1
            }
        }
    }
    return count
}

// -------------------------------------------------------------
//    func main() {
//        var n,k int32 = 6,3
//        var a = []int32 {1,3,2,6,1,2} 
//    
//        res := number_of_pairs(n,k,a)
//        fmt.Println(res)
//    }

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
    //  input line1: n, k
    //  input line2: n space-separated integers a[i]
    //  output: one integer number of pairs a[i]+a[j] %k == 0

    // read 1st line
    line1 := readLine(reader)
    line1_slice := strings.Split(line1, " ")
    var tmp_int int
    var err error

    tmp_int, err = strconv.ParseInt(line1_slice[0], 10, 32)
    checkError(err)
    n := int32(tmp_int)

    tmp_int, err = strconv.ParseInt(line1_slice[1], 10, 32)
    checkError(err)
    k := int32(tmp_int)

    if (n<1) || (n>100) || (k<=0) || (k > 100) {
        panic("Exiting on input error - invalid n or k value")
    }

    // read 2nd line
    line2 := readLine(reader)
    line2_slice := strings.Split(line2, " ")
    var a []int32
    for ii := 0; ii < int(n); ii++ {
        tmp_int, err := strconv.ParseInt(line2_slice[ii], 10, 32)
        checkError(err)
        elem := int32(tmp_int)
        if elem <=0 || elem > 100 {
            panic("Exiting on input error - invalid array value")
        }
        a = append(a, elem)
    }


    res := number_of_pairs(n,k, a)

    fmt.Println(res)

}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
