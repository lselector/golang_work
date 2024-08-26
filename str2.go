
// -------------------------------------------------------------
// testing methods to join strings together
//            test_Join : 11.240 microseconds
//         test_Sprintf : 69.625 microseconds
//             test_Add : 12.644 microseconds
//     test_Add_strings : 12.536 microseconds
//       test_Add_bytes : 17.237 microseconds
// -------------------------------------------------------------

// golang doesn't use "\" to continue code on next line.
// A line of code will automatically continue on next line 3
// unless ";" is inserted manually or automatically by language itself.
// golang uses semicolons ";" as separator/terminator.
// Here are the two rules to know:
//
//  1. When the input is broken into tokens, a semicolon 
//     is automatically inserted into the token stream 
//     immediately after a line's final token if that token is
//        - an identifier
//        - an integer, floating-point, imaginary, rune, or string literal
//        - one of the keywords: break, continue, fallthrough, or return
//        - one of the operators and delimiters ++, --, ), ], or }
//
//  2. To allow complex statements to occupy a single line, 
//     a semicolon may be omitted before a closing ")" or "}".
//
// So if the line ends in a simple operator like "+" or opening "(", etc.,
// it will continue onto the next line.
// -------------------------------------------------------------

package main

import (
    "fmt"
    "strings"
    "time"
)

// -------------------------------------------------------------
func print_elapsed_time(start time.Time, name string) {
    elapsed_micro_seconds := time.Since(start).Seconds() / 1.0e-6
    fmt.Printf(" %20s : %.3f microseconds\n", 
               name, elapsed_micro_seconds)
}

// -------------------------------------------------------------
func test_Join(N int64, td []string) {
    defer print_elapsed_time(time.Now(), "test_Join")
    for i := 0; i < int(N); i++ {
        s := strings.Join(td, ":")
        _ = s
    }
}

// -------------------------------------------------------------
func test_Sprintf(N int64, td []string) {
    defer print_elapsed_time(time.Now(), "test_Sprintf")
    for i := 0; i < int(N); i++ {
        s := fmt.Sprintf("%s:%s:%s:%s:%s", 
                         td[0], td[1], td[2], td[3], td[4])
        _ = s
    }
}

// -------------------------------------------------------------
func test_Add(N int64, td []string) {
    defer print_elapsed_time(time.Now(), "test_Add")
    for i := 0; i < int(N); i++ {
        s := td[0] + ":" + td[1] + ":" + td[2] + 
                     ":" + td[3] + ":" + td[4]
        _ = s
    }
}

// -------------------------------------------------------------
func test_Add_strings(N int64) {
    ss := "abcde"
    defer print_elapsed_time(time.Now(), "test_Add_strings")
    for i := 0; i < int(N); i++ {
        s := ss[0:1] + ":" + ss[1:2] + ":" + ss[2:3] + 
                       ":" + ss[3:4] + ":" + ss[4:5]
        _ = s
    }
}

// -------------------------------------------------------------
// careful - string[ii] returns not a char, but its code
// so we need to convert it into string
func test_Add_bytes(N int64) {
    ss := "abcde"
    defer print_elapsed_time(time.Now(), "test_Add_bytes")
    for i := 0; i < int(N); i++ {
        // note how long string is continued
        s := string(ss[0]) + ":" + string(ss[1]) + 
                             ":" + string(ss[2]) + 
                             ":" + string(ss[3]) + 
                             ":" + string(ss[4])
        _ = s
    }
}

// -------------------------------------------------------------
func main() {
    testData := []string {"a", "b", "c", "d", "e"}
    N        := int64(100)
    test_Join(N, testData)
    test_Sprintf(N, testData)
    test_Add(N, testData)
    test_Add_strings(N)
    test_Add_bytes(N)
}
