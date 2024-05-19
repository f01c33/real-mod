package main

import (
	"fmt"
	"math"
)

func Umul(x, w float64) float64 {
	return math.Floor(x/w) * w
}

func Omul(x, w float64) float64 {
	return math.Ceil(x/w) * w
}

func PrevOp(opDown, opUp func(float64, float64) float64) func(float64, float64) float64 {
	return func(x, w float64) float64 {
		return opUp(math.Floor(opDown(x, w)), w)
	}
}

func NextOp(opDown, opUp func(float64, float64) float64) func(float64, float64) float64 {
	return func(x, w float64) float64 {
		return opUp(math.Ceil(opDown(x, w)), w)
	}
}

func Mod(x, w float64) float64 {
	return x + Omul(x, w) - Umul(x+Omul(x, w), w)
}

func InlineMod(x, w float64) float64 {
	return x + math.Ceil(x/w)*w - math.Floor((x+math.Ceil(x/w)*w)/w)*w
}

func main() {
	x, w := float64(7), float64(3)
	fmt.Println(x, w, Umul(x, w), Omul(x, w))
	Umul_ := PrevOp(func(f1, f2 float64) float64 {
		return f1 / f2
	}, func(f1, f2 float64) float64 {
		return f1 * f2
	})
	Omul_ := NextOp(func(f1, f2 float64) float64 {
		return f1 / f2
	}, func(f1, f2 float64) float64 {
		return f1 * f2
	})
	fmt.Println(x, w, Umul_(x, w), Omul_(x, w))
	for i := range 20 {
		fmt.Println(Mod(float64(i-10), w))
		// fmt.Println(InlineMod(float64(i-10), w))
	}
}
