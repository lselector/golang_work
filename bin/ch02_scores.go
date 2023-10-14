package main
import "fmt"


func win(a int) int {
    if a > 0 {
        return 1
    }
    return 0
}

func main() {
    var a1, a2, a3 int
    var b1, b2, b3 int
    var pa, pb int
    fmt.Scanf("%v %v %v", &a1,&a2,&a3)
    fmt.Scanf("%v %v %v", &b1,&b2,&b3)
    pa = win(a1-b1)+win(a2-b2)+win(a3-b3)
    pb = win(b1-a1)+win(b2-a2)+win(b3-a3)
    fmt.Println(pa,pb)

}

