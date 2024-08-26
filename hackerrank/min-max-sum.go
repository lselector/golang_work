package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the miniMaxSum function below.
 */
func miniMaxSum(arr []int32) {
    n := len(arr)
    v_min := int64(arr[0])
    v_max := v_min
    v_sum := v_min
    val := int64(0)
    for ii:=1;ii<n;ii++ {
        val = int64(arr[ii])
        if val > v_max { v_max = val }
        if val < v_min { v_min = val }
        v_sum += val
    }
    fmt.Println(v_sum-v_max, v_sum-v_min)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for arrItr := 0; arrItr < 5; arrItr++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[arrItr], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
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

