package main
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

func check_jumps(x1,v1,x2,v2 int64) string {
    fmt.Println(x1,v1,x2,v2)
    dx := x2-x1
    dv := v2-v1
    if ((dv == 0) && (dx != 0)) || ((dv != 0) && (dx == 0)) {
        return "NO"
    } else if (dv == 0) && (dx == 0) {
        return "YES"
    } else if (dv*dx > 0) {
        return "NO"
    }
    // if we are here, we move in same direction and can divide.
    steps := float64(dx)/float64(dv)
    r_steps := math.Round(steps)
    if math.Abs(steps-r_steps) < 1e-6 {
        return "YES"
    } else {
        return "NO"
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
    tmp_arr := strings.Split(readLine(reader), " ")
    var arr []int64
    for ii := 0; ii < 4; ii++ {
        item, _ := strconv.ParseInt(tmp_arr[ii], 10, 64)
        arr = append(arr, item)
    }
    result := check_jumps(arr[0], arr[1], arr[2], arr[3])
    fmt.Println(result)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

