
// https://blog.golang.org/strings
// https://www.calhoun.io/6-tips-for-using-strings-in-go/

// -------------------------------------------------------------
// appending a character to a string in golang
// 
// string is a sequence of bytes
// Note - It is not required to hold Unicode text, UTF-8 text, or any other predefined format.
// But usually string consists of valid utf-8 chars.
//     ss := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
// Exampling the string: 
//    for i := 0; i < len(ss); i++ { fmt.Printf("%x ", ss[i]) }
//                             // bd b2 3d bc 20 e2 8c 98
// Or you can do same without for loop: 
//    fmt.Printf("% x\n", ss)  // bd b2 3d bc 20 e2 8c 98
//    fmt.Printf("%q\n" , ss)  // "\xbd\xb2=\xbc ⌘"
//    fmt.Printf("%+q\n", ss)  // "\xbd\xb2=\xbc \u2318"
// 
// individual chars can be represented by 1..4 bytes
// code-point is a value representing the char, 
//   for example unicode value U+2318 represents symbol ⌘.
// rune - same as code-point, but implemented as alias of type int32
// so the type and value of expression '⌘' is rune with integer value 0x2318.
//
// Summary:
//    Go source code is always UTF-8.
//    A string holds arbitrary bytes.
//    A string literal always holds valid UTF-8 sequences (plus maybe some byte-level escapes).
//    Those sequences represent Unicode code points, called runes.
//    No guarantee is made in Go that characters in strings are normalized.
//
// Note: when you loop through the string, 
//   the for range loop decodes one UTF-8-encoded rune on each iteration. 
//   Each time around the loop, the index of the loop is 
//   the starting position of the current rune (in bytes), 
//   and the code point is its value. 
//
//    const nihongo = "日本語"
//    for index, runeValue := range nihongo {
//        fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
//    }
// output:
//    U+65E5 '日' starts at byte position 0
//    U+672C '本' starts at byte position 3
//    U+8A9E '語' starts at byte position 6
//
// The for range loop and DecodeRuneInString are defined 
// to produce exactly the same iteration sequence.
// Here we use regular for loop shifting by "w" bytes
// and printing utf8.DecodeRuneInString(nihongo[i:]):
//
//    for i, w := 0, 0; i < len(nihongo); i += w {
//        runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
//        fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
//        w = width
//    }
// output same:
//    U+65E5 '日' starts at byte position 0
//    U+672C '本' starts at byte position 3
//    U+8A9E '語' starts at byte position 6



    
// utf-8 encoded sequence of runes
// rune is int32
//
//    In Go strings are not sequences of runes, 
//    they are utf-8 encoded sequences of runes. 
//    When you range over a string you get runes, 
//    but you cannot simply append a rune to a string. 
//    For example: euro sign '€' is an integer 0x20AC 
//    (this is called code point). 
//    But when you encode euro sign in utf-8 you get 3 bytes: 
//    0xE2 0x82 0xAC 
//    http://www.fileformat.info/info/unicode/char/20aC/index.htm
//
// So appending a char actually works like this:
//     s = append(s, encodeToUtf8(c))
// Note that encodings are done at compile time.
//
// Utf-8 can encode a character with 1, 2, 3, or 4 bytes. 
// Utf-16 can encode a character with 2 or with 4 bytes.
// So Go usually appends 1 byte (for ascii) or 2, 3, 4 bytes for Chinese, 
// and Java usually appends 2 bytes (for ascii) or 4 bytes for Chinese.
// Since most characters that we (west) use can be encoded with 2 bytes 
// Java gives the false belief that strings are sequences of 2byte char-s,
// which is true until you need to encode some Chinese
//

//    go strings are a sequence of bytes, not runes. 
//    Ranging over a string implicitly returns runes, 
//    along with the byte indexes in the string.
//    
//    Actually they are not because not every byte sequence is a valid string. 
//    They can be type casted into sequences of bytes.
//    
//    No, I assure you they are (there is no type casting, only conversion). 
//    The string data structure consists solely of an array of bytes, 
//    and the length that of that array. 
//    The range operator scans for runes, but the string is still bytes.
//    The string and []byte types have the same underlying data, 
//    and can be converted in both directions. 
//    
//    I just wanted to point out that even thou "under the hood" some things are same, 
//    does not mean type is the same. 
//    For example int32 and unit32 are same under the hood 
//    but behave differently in sign operations or when printing. 
//    Contrast this with int32 and rune which are same both in value and in type (aliases)
//    


// -------------------------------------------------------------

package main

import (
    "fmt"
    "unicode/utf8"
)

// -------------------------------------------------------------
func main(){
    test1()
    test2()
    test3()
}

// -------------------------------------------------------------
func test1() {
    s := "hello";
    c := 'x';
    fmt.Println(s + string(c));
}

// -------------------------------------------------------------
func test2() {
    s := "美国必须死"
    b := []byte(s)
    fmt.Println(s)

    fmt.Println("len(s):", len(s))
    fmt.Println("len(b):", len(b))

    for i, r := range s {
        fmt.Printf("Index %d, rune %s, from bytes: %s, from string: %s\n",
            i, string(r), b[i:i+utf8.RuneLen(r)], s[i:i+utf8.RuneLen(r)])
    }

}

// -------------------------------------------------------------
func test3() {
    b := []byte("美国必须死")

    fmt.Println("range equiv:")
    for i := 0; i < len(b); {
        r, width := utf8.DecodeRune(b[i:])
        fmt.Println("  ", i, string(r), string(b[i:i+width]))
        i += width
    }

}

// -------------------------------------------------------------
func test4() {
    s := "美国必须死"

    fmt.Println("range:")
    for i, r := range s {
        fmt.Println("  ", i, string(r), s[i:i+utf8.RuneLen(r)])
    }

    fmt.Println("range equiv:")
    for i := 0; i < len(s); {
        r, width := utf8.DecodeRuneInString(s[i:])
        fmt.Println("  ", i, string(r), s[i:i+width])
        i += width
    }

}
