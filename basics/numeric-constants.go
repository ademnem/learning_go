package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)
// an untyped constant takes the type needed by its constext

func needInt(x int) int { return x*10 + 1 } // 32 or 64 bits depending on the system
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small)) // 21
	fmt.Println(needFloat(Small)) // 0.2
	fmt.Println(needFloat(Big)) // 1.2676506002282295e+29
    fmt.Println(needInt(Big)) // ./prog.go:22:22: cannot use Big (untyped int constant 1267650600228229401496703205376) as int value in argument to needInt (overflows)
    fmt.Println(Big) // does not work because a type can not be inferred
}
