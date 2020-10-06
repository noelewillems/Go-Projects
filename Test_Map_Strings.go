// This is the starting point for the Go map lab
package main

import "fmt"
type myType string

func main() {
	var inSlice = []myType{"one", "fish", "two", "fish", "red", "fish", "blue", "fish"}
	fmt.Println(inSlice)
	var resultSlice []myType

	var endWithX mapOp = func (x myType) myType {
		ret:= x + "X"
	  	return ret
	}
	resultSlice = mymap(inSlice, endWithX)
	fmt.Println(resultSlice)
}