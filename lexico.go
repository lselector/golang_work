package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// -------------------------------------------------------------
func myprocess(ss1,ss2 string) string {
    if len(ss1) == 0 {
        return ss2
    } else if len(ss2) == 0 {
        return ss1
    }
    b1 := []rune(ss1)
    b2 := []rune(ss2)
    l1 := len(b1)
    l2 := len(b2)
//    fmt.Println(b1, l1)
//    fmt.Println(b2, l2)
    b3 := make([]rune, l1+l2)
    var i1, i2, i3 int
    for {
//        fmt.Println(i1,i2,i3)
        if (i1 >= l1) && (i2 >= l2) {
            break
        } else if (i1 >= l1) && (i2 < l2) {
            b3[i3] = b2[i2]
            i3 += 1
            i2 += 1
        } else if (i1 < l1) && (i2 >= l2) {
            b3[i3] = b1[i1]
            i3 += 1
            i1 += 1
        } else if (i1 < l1) && (i2 < l2) {
            if b1[i1] < b2[i2] {
                b3[i3] = b1[i1]
                i3 += 1
                i1 += 1
            } else if b1[i1] > b2[i2]{
                b3[i3] = b2[i2]
                i3 += 1
                i2 += 1
            } else {
            // ------------------------
                if string(b1[i1:]) <= string(b2[i2:]) {
                    b3[i3] = b1[i1]
                    i3 += 1
                    i1 += 1
                } else {
                    b3[i3] = b2[i2]
                    i3 += 1
                    i2 += 1
                }
            // ------------------------
            }
        }
    }

    return string(b3)

}

// -------------------------------------------------------------
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n_cases := int(nTemp)
    var arr []string

    for ii:= 0; ii<n_cases*2; ii++ {
        arr = append(arr, readLine(reader))
    }

//    n_cases := 2
//    arr := []string {"JACK", "DANIEL","ABACABA","ABACABA"}

    for ii:=0; ii<n_cases; ii++ {
        ss := myprocess(arr[ii*2],arr[ii*2+1])
        fmt.Println(ss)
    }
}

    // -------------------------------------------------------------
    func readLine(reader *bufio.Reader) string {
        str, _, err := reader.ReadLine()
        if err == io.EOF {
            return ""
        }

        return strings.TrimRight(string(str), "\r\n")
    }

    // -------------------------------------------------------------
    func checkError(err error) {
        if err != nil {
            panic(err)
        }
    }

