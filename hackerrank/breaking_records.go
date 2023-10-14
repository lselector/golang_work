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
 * Complete the breakingRecords function below.
 */
func breakingRecords(score []int32) []int32 {
    var cnt_min, cnt_max, mymin, mymax int32
    if len(score) <= 0 {
        return []int32 {0,0}
    }
    mymin = score[0]
    mymax = score[0]
    for _, sc := range score {
        if sc > mymax {
            cnt_max += 1
            mymax = sc
        } else if sc < mymin {
            cnt_min += 1
            mymin = sc
        }
    }
    return []int32 {cnt_max,cnt_min}
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer outputFile.Close()

    writer := bufio.NewWriterSize(outputFile, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    scoreTemp := strings.Split(readLine(reader), " ")

    var score []int32

    for scoreItr := 0; scoreItr < int(n); scoreItr++ {
        scoreItemTemp, err := strconv.ParseInt(scoreTemp[scoreItr], 10, 64)
        checkError(err)
        scoreItem := int32(scoreItemTemp)
        score = append(score, scoreItem)
    }

    result := breakingRecords(score)

    for resultItr, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if resultItr != len(result) - 1 {
            fmt.Fprintf(writer, " ")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
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

