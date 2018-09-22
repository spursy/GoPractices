package main

import "fmt"

func main()  {
	complexArray1 := [3][]string {
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}

	fmt.Printf("The complex array is %v\n", complexArray1)

	complexArray2 := modifyArray(complexArray1)
	fmt.Printf("After updated, the complex array is %v\n", complexArray2)

	fmt.Printf("The origin complex array is %v\n", complexArray1)
}

func modifyArray(a [3][]string)([3][]string){
	a[1][1] = "O"
	return a
} 