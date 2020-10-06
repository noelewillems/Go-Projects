// Noel Willems

package main
import "fmt"

type myType string

func main() {
	// Input slice of type myType
	var inSlice = []myType{"one", "fish", "two", "fish", "red", "fish", "blue", "fish"}
	fmt.Println(inSlice)
	// "Output" slice of type myType
	var resultSlice []myType

	// Map function - appends "X" to end of strings
	var endWithX mapOp = func (x myType) myType {
		ret:= x + "X"
	  	return ret
	}
	resultSlice = mymap(inSlice, endWithX)
	fmt.Println(resultSlice)
}