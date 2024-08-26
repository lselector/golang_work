// script recursively walks through directory tree
// looks for .wav files 
// invokes go-routines to convert them into .ogg files
// using a shell executable "oggenc"


package main

import (
    "fmt"
    "log"
    "path/filepath"
    "os"
    "os/exec"
    "strings"
    "sync"
)

var waitGroup1 = &sync.WaitGroup{}

// -------------------------------------------------------------
func convert_to_ogg(path string) {
    oggname := fmt.Sprintf("%v.ogg", strings.TrimSuffix(path,".wav"))
    cmd := exec.Command("/usr/local/bin/oggenc", "-q", "10", "-o", oggname, path)
    log.Printf("Running command and waiting for it to finish...")
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
}

// -------------------------------------------------------------
func visit(path string, info os.FileInfo, err error) error {
    //fmt.Printf("visited %v\n", path)
    if err != nil {
        log.Fatal(err)
    }

    if strings.HasSuffix(path, ".wav") {
        fmt.Printf("visited %v\n", path)
        waitGroup1.Add(1)
        go func (filePath string) {
            defer waitGroup1.Done()
            convert_to_ogg(filePath)
            os.Remove(filePath)
        } (path)
    }

    return nil
}

// -------------------------------------------------------------
func main() {
    dir := ""

    filepath.Walk(dir, visit)
    waitGroup1.Wait()

}