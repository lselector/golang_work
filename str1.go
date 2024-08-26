
// tests of string operations
// https://blog.golang.org/strings

package main
import (
    "fmt"
    "unicode/utf8"
    "strings"
    "regexp"
)

// -------------------------------------------------------------
func main() {
    test1()
    test2()
    test3()
    split_into_words() 
    trim_split_join()
    split_regex()
}

// -------------------------------------------------------------
func test1() {
    fmt.Println("---- loop through bytes, print formats")
    const ss = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
    fmt.Print("Print:")
    fmt.Println(ss)
    fmt.Print("Byte loop       : ")
    for i := 0; i < len(ss); i++ {
        fmt.Printf("%x ", ss[i])
    }
    fmt.Printf("\n")
    fmt.Printf("Printf with %%x  : %x\n", ss)
    fmt.Printf("Printf with %% x : % x\n", ss)
    fmt.Printf("Printf with %%q  : %q\n", ss)
    fmt.Printf("Printf with %%+q : %+q\n", ss)
}

// -------------------------------------------------------------
func test2() {
    fmt.Println("---- utf8 character")
    const ss = "\xe2\x8c\x98"
    fmt.Printf("plain  string : %s\n", ss)
    fmt.Printf("quoted string : %+q\n", ss)
    fmt.Printf("hex bytes     : ")
    for i := 0; i < len(ss); i++ {
        fmt.Printf("%x ", ss[i])
    }
    fmt.Printf("\n")
}

// -------------------------------------------------------------
func test3() {

    fmt.Println("------------------------------- runeValue 1")
    const ss = "\u65E5\u672C\u8A9E"

    for index, runeValue := range ss {
        fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
    }

    fmt.Println("------------------------------- runeValue 2")

    for i, w := 0, 0; i < len(ss); i += w {
        runeValue, width := utf8.DecodeRuneInString(ss[i:])
        fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
        w = width
    }

}

// -------------------------------------------------------------
func split_into_words(){
    fmt.Println("------------------------------- split_and_join")

    ss := `one    two   three four
 five    six
seven

`
    words := strings.Fields(ss)
    for ii,vv := range words {
        fmt.Printf("%d:%s ",ii,vv)
    }
    fmt.Println()
}


// -------------------------------------------------------------
func trim_split_join() {
    fmt.Println("------------------------------- trim_white_space")
    aa := "\t Hello, \n World \n\n\n "
    fmt.Println(len(aa)," : ", aa)

    fmt.Println("using strings.TrimSpace()")
    bb := strings.TrimSpace(aa)
    fmt.Println(len(bb)," : ", bb)

    fmt.Println("using strings.Replace() for each whitespace char")
    cc := strings.Replace(aa, "\n", " ", -1)
    cc = strings.Replace(cc, " ", "", -1)
    fmt.Println(len(cc)," : ", cc)

    fmt.Println("using strings.Fields() and strings.Join()")
    // simpliest thing - split into words - and join back
    dd := strings.Join(strings.Fields(aa), " ")
    fmt.Println(len(dd)," : ", dd)

}

// -------------------------------------------------------------
func split_regex() {
    fmt.Println("------------------------------- trim_white_space")
    aa := "\t Hello, \n World \n\n\n "
    fmt.Println(len(aa)," : ", aa)

    fmt.Println("split using regexp.MustCompile() , r.FindAllString(), strings.Join()")
    r2 := regexp.MustCompile("[^\\s]+")
    cc := strings.Join(r2.FindAllString(aa, -1), " ")
    fmt.Println(len(cc)," : ", cc)

}
