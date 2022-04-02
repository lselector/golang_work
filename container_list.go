package main

// https://www.slideshare.net/jgrahamc/go-containers

import (
    "fmt"
    "container/list"
)

func main() {
    mylist := list.New()
    e0 := mylist.PushBack(42)
    e1 := mylist.PushFront(13)
    e2 := mylist.PushBack(7)
    mylist.InsertBefore(3,e0)
    mylist.InsertAfter(196,e1)
    mylist.InsertAfter(1729,e2)

    for e := mylist.Front(); e != nil; e = e.Next() {
        fmt.Printf("%d ", e.Value.(int))
    }
    fmt.Printf("\n")    
    
}
