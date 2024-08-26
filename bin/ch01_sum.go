package main
import "fmt"

func solveMeFirst(a uint32,b uint32) uint32{
  // Hint: Type return (a+b) below
    return (a+b)
}

func main() {
    var a, b, res uint32
    fmt.Print("Enter 1st number:")
    fmt.Scanf("%v", &a)
    fmt.Print("Enter 2nd number:")
    fmt.Scanf("%v", &b)
    res = solveMeFirst(a,b)
    fmt.Println(res)
}

