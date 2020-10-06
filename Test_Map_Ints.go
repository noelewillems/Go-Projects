// This is the starting point for the Go map lab
package main

import "fmt"
type myType int

func main() {
	var inSlice = []myType{10, 15, 20, 25, 30, 35, 40}
	fmt.Println(inSlice)
	var resultSlice []myType
	
	var add2 mapOp = func(x myType) myType {
	  return x + 2
	}
	resultSlice = mymap(inSlice, add2)
	fmt.Println(resultSlice)

	var doubleIt mapOp = func(x myType) myType {
		return x * 2
	}
	resultSlice = mymap(inSlice, doubleIt)
	fmt.Println(resultSlice)
}