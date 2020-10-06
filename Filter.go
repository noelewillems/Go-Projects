package main

type filterOp func(a myType) myType

func myfilter(in []myType, myOp filterOp) []myType {
	var result = make([]myType, 0)
	for i := 0; i < len(in); i++ {
		if(filterOp(in[i])) {
			result = append(result, in[i])
		}
	}
	return result
}