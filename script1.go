package main

// https://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/
//


import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    var (
        cmdOut []byte
        err    error
    )
    cmdName := "ls"
    cmdArgs := []string{"-a", "-l", "-F"}
    if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
        fmt.Fprintln(os.Stderr, "There was an error running command: ", err)
        os.Exit(1)
    }
    sha := string(cmdOut)
    fmt.Println(sha)
}
