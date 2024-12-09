package main

import (
	"fmt"
	"mymath/basicmath"
)

func main() {
	fmt.Println("Hello, World!")

	f1 := basicmath.NewFraction(1, 2)

	f2 := basicmath.NewFraction(1, 3)

	result := basicmath.AddTwo(f1, f2)
	fmt.Printf("AddTwo(%v + %v) = %v\n", f1, f2, result) // Output: 5/6
	result = basicmath.DivideTwo(f1, f2)
	fmt.Printf("DivideTwo(%v / %v) = %v\n", f1, f2, result) // Output: 3/2
	result = basicmath.MultiplyTwo(f1, f2)
	fmt.Printf("MultiplyTwo(%v * %v) = %v\n", f1, f2, result) // Output: 1/6
	result = basicmath.SubtractTwo(f1, f2)
	fmt.Printf("SubtractTwo(%v - %v) = %v\n", f1, f2, result) // Output: 1/6

	// f := f1.Add(f2)
	// fmt.Printf("%v + %v = %v\n", f1, f2, f)

	// result := basicmath.NewFraction(
	// 	basicmath.WithNumerator(63),
	// 	basicmath.WithDenominator(49),
	// )
	// fmt.Printf("%v = %v\n", f, result.Factor())
	// result.Simplify()
	// fmt.Printf("%v + %v = %s\n", f1, f2, result)

	// f = f1.Subtract(f2)
	// fmt.Printf("%v - %v = %v\n", f1, f2, f)

	// f = f1.Multiply(f2)
	// fmt.Printf("%v * %v = %v\n", f1, f2, f)

	// f = f1.Divide(f2)
	// fmt.Printf("%v / %v = %v\n", f1, f2, f)

	// fmt.Println()

	// var o1 interfaces.Operable = basicmath.NewFraction(
	// 	basicmath.WithNumerator(3),
	// 	basicmath.WithDenominator(7),
	// )
	// var o2 interfaces.Operable = basicmath.NewFraction(
	// 	basicmath.WithNumerator(5),
	// 	basicmath.WithDenominator(7),
	// )
	// t := o1.Add(o2)
	// fmt.Printf("%v + %v = %v\n", o1, o2, t)

	// result = basicmath.NewFraction(
	// 	basicmath.WithNumerator(56),
	// 	basicmath.WithDenominator(49),
	// )
	// fmt.Printf("%v = %v\n", t, result.Factor())
	// result.Simplify()
	// fmt.Printf("%v + %v = %s\n", o1, o2, result)

	// var L interfaces.LaTeXer = basicmath.NewFraction(
	// 	basicmath.WithNumerator(16),
	// 	basicmath.WithDenominator(25),)

	// fmt.Printf("L: %s", L.LaTeX())
}
