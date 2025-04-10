package main

import (
	"fmt"
	"mymath/geometry"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println()
	fmt.Println("Hello, World!")	
	fmt.Println()

	line := geometry.NewLine(geometry.Point{X: 2, Y: 1}, geometry.Point{X: 8, Y: 6})

	fmt.Printf("%s\n", line.LaTeX())

	// floatToFraction(2383158)

	// floatToFraction(123.456789)

	// a := geometry.Point{X: 3, Y: 4}
	// b := geometry.Point{X: 1, Y: 2}

	// fmt.Printf("Distance between a and b: %.2f\n", a.Distance(b))
	// fmt.Println("a + b =", a.Add(b))
	// fmt.Println("a - b =", a.Subtract(b))
	// fmt.Println("a scaled by 2 =", a.Scale(2))
	// fmt.Printf("Dot product of a and b: %.2f\n", a.Dot(b))
	// fmt.Printf("Magnitude of a: %.2f\n", a.Magnitude())

	// a.Move(1, -1)
	// fmt.Println("a after Move(1, -1):", a)

	// f1 := basicmath.NewFraction(1, 2)

	// f2 := basicmath.NewFraction(1, 3)

	// result := basicmath.AddTwo(f1, f2)
	// fmt.Printf("AddTwo(%v + %v) = %v\n", f1, f2, result) // Output: 5/6
	// result = basicmath.DivideTwo(f1, f2)
	// fmt.Printf("DivideTwo(%v / %v) = %v\n", f1, f2, result) // Output: 3/2
	// result = basicmath.MultiplyTwo(f1, f2)
	// fmt.Printf("MultiplyTwo(%v * %v) = %v\n", f1, f2, result) // Output: 1/6
	// result = basicmath.SubtractTwo(f1, f2)
	// fmt.Printf("SubtractTwo(%v - %v) = %v\n", f1, f2, result) // Output: 1/6

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
	fmt.Println()
}

func floatToFraction(value float64) {
	s := strconv.FormatFloat(value, 'f', -1, 64)
	parts := strings.Split(s, ".")
	leftPart := int(math.Floor(value))
	rightPart := value - float64(leftPart)

	fmt.Printf("value: %f\n", value)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("parts: %v\n", parts)
	fmt.Printf("leftPart: %d\n", leftPart)
	fmt.Printf("rightPart: %f\n", rightPart)

	if len(parts) == 2 {
		length := len(parts[1])
		fmt.Printf("length: %v\n", length)
		
		nTemp, _ := strconv.Atoi(parts[0])
		dTemp, _ := strconv.Atoi(parts[1])
		t := float64(dTemp)
		f := math.Pow10(length)

		fmt.Println("After decimal:", parts[1])
		fmt.Println("t: ", t)
		fmt.Println("f: ", f)

		t /= f
		fmt.Println("t: ", t)

		d := int(math.Pow10(length))
		n := (nTemp * d) + dTemp

		fmt.Println("Numerator: ", n)
		fmt.Println("Denominator: ", d)
	} else {
		fmt.Printf("Numerator: %s\n", parts[0])
	}

	fmt.Println()
}
