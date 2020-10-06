// Noel Willems

package main

type filterOp func(a myType) bool

func myfilter(in []myType, myOp filterOp) []myType {
	// Return result
	var result = make([]myType, 0)
	// Iterate through input
	for i := 0; i < len(in); i++ {
		// If the operation returns true, append to empty result slice
		if(myOp(in[i])) {
			result = append(result, in[i])
		}
	}
	return result
}