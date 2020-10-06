package main

type mapOp func(a myType) myType

func mymap(in []myType, myOp mapOp) []myType {
	var result = make([]myType, len(in))
	for i := 0; i < len(in); i++ {
	  result[i] = myOp(in[i])
	}
	return result
}