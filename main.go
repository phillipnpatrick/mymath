package main

import (
	"mymath/interfaces"
	"mymath/basicmath"
	"fmt"
)

func main() {
    fmt.Println("Hello, World!")

	f1 := basicmath.NewFraction(
		basicmath.WithNumerator(3),
		basicmath.WithDenominator(7),
	)

	f2 := basicmath.NewFraction(
		basicmath.WithNumerator(6),
		basicmath.WithDenominator(7),
	)

	f := f1.Add(f2)
	fmt.Printf("%v + %v = %v\n", f1, f2, f)

	result := basicmath.NewFraction(
		basicmath.WithNumerator(63),
		basicmath.WithDenominator(49),
	)
	fmt.Printf("%v = %v\n", f, result.Factor())
	result.Simplify()
	fmt.Printf("%v + %v = %s\n", f1, f2, result)

	f = f1.Subtract(f2)
	fmt.Printf("%v - %v = %v\n", f1, f2, f)

	f = f1.Multiply(f2)
	fmt.Printf("%v * %v = %v\n", f1, f2, f)

	f = f1.Divide(f2)
	fmt.Printf("%v / %v = %v\n", f1, f2, f)

	fmt.Println()

	var o1 interfaces.Operable = basicmath.NewFraction(
		basicmath.WithNumerator(3),
		basicmath.WithDenominator(7),
	)
	var o2 interfaces.Operable = basicmath.NewFraction(
		basicmath.WithNumerator(5),
		basicmath.WithDenominator(7),
	)
	t := o1.Add(o2)
	fmt.Printf("%v + %v = %v\n", o1, o2, t)
	
	result = basicmath.NewFraction(
		basicmath.WithNumerator(56),
		basicmath.WithDenominator(49),
	)
	fmt.Printf("%v = %v\n", t, result.Factor())
	result.Simplify()
	fmt.Printf("%v + %v = %s\n", o1, o2, result)
}