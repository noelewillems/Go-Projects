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

	// Filter function - returns true if even
	var isEven filterOp = func(x myType) bool {
		return x % 2 == 0
	}
	resultSlice = myfilter(inSlice, isEven)
	fmt.Println(resultSlice)
}