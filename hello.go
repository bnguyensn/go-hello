package main

import (
	"fmt"
)

// VarAndPointer contains an integer variable and a pointer that points to it.
type VarAndPointer struct {
	Var     int
	Pointer *int
}

// CreateVarAndPointer generates a VarAndPointer struct from an integer.
func CreateVarAndPointer(n int) VarAndPointer {
	i := n
	p := &i

	return VarAndPointer{i, p}
}

// PrintVarAndPointer logs the values of a VarAndPointer struct.
func PrintVarAndPointer(varAndPointer VarAndPointer) {
	fmt.Printf(
		"Underlying value: %v, pointer's value: %v, pointer's memory address: %v\n",
		varAndPointer.Var,
		*varAndPointer.Pointer,
		varAndPointer.Pointer,
	)
}

// PrintSlice logs the values of a slice.
func PrintSlice(slice []string) {
	fmt.Printf("%s len=%d cap=%d\n\n", slice, len(slice), cap(slice))
}

func main() {
}
