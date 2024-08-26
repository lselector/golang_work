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
func main() {

    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
    fmt.Print("Enter space-separated numbers:")
    ss := readLine(reader)
    ss_slice := strings.Fields(ss)

    // parse as int or float
    var arr []int32
    // var arr []float32
    nn := len(ss_slice)

    for ii := 0; ii < nn; ii++ {
        tmp, err := strconv.ParseInt(ss_slice[ii], 10, 64)
        // tmp, err := strconv.ParseFloat(ss_slice[ii], 64)
        panic_on_error(err)
        arr = append(arr, int32(tmp))
    }

    fmt.Println(arr)

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
func panic_on_error(err error) {
    if err != nil {
        panic(err)
    }
}

