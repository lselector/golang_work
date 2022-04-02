//
//  setting env variable from inside Go
//

package main

import (
    "fmt"
    "os"
)

func main() {

    home := os.Getenv("HOME")
    fmt.Println("PWD:", os.Getenv("PWD"))

    newdir := home+"/gowork"
    err := os.Chdir(newdir)
    if err != nil {
        panic(err)
    }
    fmt.Println("PWD:", os.Getenv("PWD"))

    os.Setenv("PWD", newdir)
    fmt.Println("PWD:", os.Getenv("PWD"))

}


