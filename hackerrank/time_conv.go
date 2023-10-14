package main

import (
    "fmt"
    "strings"
    "strconv"
)

// convert time to 24h format

func timeConversion(s string) string {
    nn := len(s)
    bb := strings.Split(s[:nn-2], ":")
    ii, _ := strconv.Atoi(bb[0])
    ii = ii % 12
    if s[nn-2:] == "PM" {
        ii += 12
    }
    return fmt.Sprintf("%02d:%s:%s",ii,bb[1],bb[2])    
}

func main() {
    times_in := []string {"12:00:00AM", "12:00:00PM", 
                          "12:00:01AM", "12:00:01PM",
                          "11:00:00AM", "11:00:00PM", 
                          "00:00:00AM", "00:00:00PM", 
                          "00:00:01AM", "00:00:01PM", 
                          "01:00:00AM", "01:00:00PM", 
                          }
    for _, t_in := range times_in {
        t_out := timeConversion(t_in)
        fmt.Println(t_in, " => ", t_out)
    }
}


