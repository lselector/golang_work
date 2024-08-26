// -------------------------------------------------------------
// Golang solution
// https://play.golang.org/p/T7K1VJEAFA
// -------------------------------------------------------------
package main

import (
    "fmt"
    "time"
//    "os"
//    "os/exec" 
)

const Towers = 3
const Disks = 7

type Hanoi [Towers][]int

func main() {
    var state Hanoi
    for ii := 0; ii<=Disks+2; ii++ {
        fmt.Printf("\n")
    }
    state.init(Disks)
    state.move(Disks, 0, 1, 2)
}

// ---------------------------------------------------
func (h *Hanoi) init(n_disks int) {
    h[0] = make([]int, n_disks)
    for i := range h[0] {
        h[0][i] = n_disks - i
    }
    h.print()
}

// ---------------------------------------------------
func (h *Hanoi) move(n_disks, a, b, c int) {
    if n_disks <= 0 {
        return
    }
    h.move(n_disks-1, a, c, b)  // n-1 disks a->b via c
    
    disk := h[a][len(h[a])-1]   // take disk value from  a
    h[a] = h[a][:len(h[a])-1]   // remove it from a
    h[c] = append(h[c], disk)   // add it to c
    h.print()
    
    h.move(n_disks-1, b, a, c)  // n-1 disks b->c via a
}

// ---------------------------------------------------
func (h *Hanoi) print() {
    // https://en.wikipedia.org/wiki/ANSI_escape_code
    fmt.Printf("\033[%dA",Disks+2)
    for i := Disks; i >= 0; i-- {
        for j := 0; j < Towers; j++ {
            if i == 0 {
                fmt.Print("_/||\\_")
            } else if len(h[j]) >= i {
                fmt.Printf("  %02d  ", h[j][i-1])
            } else {
                fmt.Print("  ||  ")
            }
        }
        fmt.Println()
    }
    fmt.Println()
    time.Sleep(time.Second / 10)
}

