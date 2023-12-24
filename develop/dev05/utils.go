package main

import "math"

func clamp(val, lb, rb int) int {
	return int(math.Max(float64(lb), math.Min(float64(val), float64(rb))))
}