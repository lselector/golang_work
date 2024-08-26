//
//  writing/reading/removing/checking files,
//
//  Note interesting topic of using defer fh.Close()
//  The problem is that the Close() can return an error itself.
//      https://groups.google.com/forum/#!topic/golang-nuts/-eo7navkp10
//  Official go blogpost that misses this subtlety
//      https://blog.golang.org/defer-panic-and-recover
//  Interesting blogpost on the topic:
//      https://pocketgophers.com/handling-errors-in-defer/
//

package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "io/ioutil"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// -------------------------------------------------------------
func write_file(fname string) {
    fmt.Print("\n---- write_file ", fname)
    fmt.Println(" using ioutil.WriteFile")
    d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile(fname, d1, 0644)
    check(err)

    fmt.Println("or open file for writing:")
    fh, err := os.Create(fname)
    check(err)
    defer fh.Close()

    // Write byte slices
    d2 := []byte("line1\n")
    // d2 := []byte{115, 111, 109, 101, 10}
    n2, err := fh.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    // write line of text
    n3, err := fh.WriteString("line2\n")
    fmt.Printf("wrote %d bytes\n", n3)

    // Sync flushes writes to stable storage
    fh.Sync()

    // bufio provides buffered writers and readers
    w := bufio.NewWriter(fh)
    n4, err := w.WriteString("line3\n")
    fmt.Printf("wrote %d bytes\n", n4)

    // Use Flush to push everything to disk
    w.Flush()
}

// -------------------------------------------------------------
func read_file(fname string) {
    fmt.Print("\n---- read_file ", fname)
    fmt.Println(" using os.Open(fname) and bufio.NewScanner(fh)")
    fh, err := os.Open(fname)
    if err != nil {
        log.Fatal(err)
    }
    // if we are here - file was open for reading
    // so we specify to close it at the end
    defer fh.Close()  // note - here we ignore possible error on closing

    scanner := bufio.NewScanner(fh)

    for scanner.Scan() {             // internally, it advances token based on sperator
        fmt.Println(scanner.Text())  // token in unicode-char
      //fmt.Println(scanner.Bytes()) // token in bytes
    }
}


// -------------------------------------------------------------
func slurp_file(fname string) {
    fmt.Print("---- slurp_file ", fname)
    fmt.Println(" using ioutil.ReadFile(fname)")

    b, err := ioutil.ReadFile(fname) // read bytes from file
    if err != nil {
        fmt.Print(err)
    }
    str := string(b) // convert content from bytes to a string
    str = strings.TrimSpace(str)
    fmt.Println(str)
}

// -------------------------------------------------------------
func delete_file(fname string){
    var err = os.Remove(fname)
    check(err)
    fmt.Println("\n---- done deleting file", fname)
}

// -------------------------------------------------------------
func detect_file(fname string) {
    fmt.Println("\n---- detect_file", fname)
    var _, err = os.Stat(fname)
    if os.IsNotExist(err) {
        fmt.Printf("file %s doesn't exist\n", fname)
    } else {
        fmt.Printf("file %s exists\n", fname)
    }
}

// -------------------------------------------------------------
func main() {
    fname := "./junk2.txt"
    detect_file(fname)
    write_file(fname)
    read_file(fname)
    slurp_file(fname)
    delete_file(fname)
    fmt.Println()
}
