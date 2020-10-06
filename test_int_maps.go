// Noel Willems
package main
import "fmt"

type myType int

func main() {
	// Input slice of type myType
	var inSlice = []myType{10, 15, 20, 25, 30, 35, 40}
	fmt.Println(inSlice)
	// "Output" slice of type myType
	var resultSlice []myType
	
	// Map function - adds 2
	var add2 mapOp = func(x myType) myType {
	  return x + 2
	}
	resultSlice = mymap(inSlice, add2)
	fmt.Println(resultSlice)

	// Map function - doubles
	var doubleIt mapOp = func(x myType) myType {
		return x * 2
	}
	resultSlice = mymap(inSlice, doubleIt)
	fmt.Println(resultSlice)
}