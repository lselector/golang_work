// package main line 1
// package main line 2
package slice1

import (
	"fmt"
)

func main() {
	var myarr [5]int
	for i := 0; i < 5; i++ {
		myarr[i] = i
	}
	myslice := myarr[0:3]

	fmt.Printf("myarr: %v\n", myarr)
	fmt.Printf("myslice: %v\n", myslice)

	//myslice[8] = 1500
	//fmt.Printf("myarr: %v\n", myarr)
	//fmt.Printf("myslice: %v\n", myslice)

	fmt.Println("\n\nDOING APPEND\n")

	myslice2 := append(myslice, 3000)
	fmt.Printf("myarr: %v\n", myarr)
	fmt.Printf("myslice: %v\n", myslice)
	fmt.Printf("myslice2: %v\n", myslice2)

	fmt.Println("\n\nDOING APPENDs\n")
	myslice3 := append(myslice2, 4000, 5000, 6000)
	fmt.Printf("myarr: %v\n", myarr)
	fmt.Printf("myslice: %v\n", myslice)
	fmt.Printf("myslice2: %v\n", myslice2)
	fmt.Printf("myslice3: %v\n", myslice3)

	fmt.Println("\n\nALTERING SLICE\n")
	myslice3[0] = 50000
	fmt.Printf("myarr: %v\n", myarr)
	fmt.Printf("myslice: %v\n", myslice)
	fmt.Printf("myslice2: %v\n", myslice2)
	fmt.Printf("myslice3: %v\n", myslice3)

}
