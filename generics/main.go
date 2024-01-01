package main

import "fmt"

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntOrFloats[K comparable, V int64 | float64] (m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number] (m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	ints := map[string]int64{
		"first" : 12,
		"s" : 12,
	}

	floats := map[string]float64{
		"first" : 12.0,
		"s" : 12.2,
	}

	fmt.Printf("%v %v\n", SumInts(ints), SumFloats(floats))

	fmt.Printf("%v %v\n", SumIntOrFloats(ints), SumIntOrFloats(floats))

	fmt.Printf("%v %v\n", SumNumbers(ints), SumNumbers(floats))
}
