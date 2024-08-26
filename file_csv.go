// file_csv.go

package main

import (
    "fmt"
    "os"
    "encoding/csv"
    "strconv"
    "io"
)

type myrow_type struct {
    myrow   int
    myname  string
    mynum   float64
}

// -------------------------------------------------------------
func write_csv_file() {
    fmt.Println("---- write_csv_file ")
    fh, err := os.Create("./junk.csv")
    if err != nil {
        panic(err)
    }
    defer fh.Close()
    mytext := "1,mama,1.1\n2,papa,2.2\n"
    _, err = fh.WriteString(mytext)
    fh.Sync()
}

// -------------------------------------------------------------
func parse_row(record []string) myrow_type {
    result := myrow_type{0,"",0.0}
    ii, _ := strconv.Atoi(record[0])
    result.myrow = ii
    result.myname = record[1]
    rr, _ := strconv.ParseFloat(record[2], 64)
    result.mynum  = rr
    return result
}

// -------------------------------------------------------------
func read_csv_file() {
    fname := "./junk.csv"
    fmt.Println("---- read_csv_file ", fname)
    fh, err := os.Open(fname)
    if err != nil {
        panic(err)
    }
    defer fh.Close()
    r := csv.NewReader(fh)
    myrows := []myrow_type{}
    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            panic(err)
        }
        myrows = append(myrows, parse_row(record))
        for ii := range record {
            fmt.Printf(">>> %d %s %T\n",ii,record[ii], record[ii])
        }
        fmt.Println("-----")
    }
    fmt.Println(myrows)
    
}


func main() {
    write_csv_file()
    read_csv_file()
}
