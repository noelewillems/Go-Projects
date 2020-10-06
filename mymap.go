// Noel Willems

package main

type mapOp func(a myType) myType

func mymap(in []myType, myOp mapOp) []myType {
	// Output slice
	var result = make([]myType, len(in))
	for i := 0; i < len(in); i++ {
		// Set result slice to myOp elements
	 	result[i] = myOp(in[i])
	}
	return result
}