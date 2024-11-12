package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
    // you don't need () but you can still have it
	if x < 0 {
		return sqrt(-x) + "i"
	} // {} are necessary
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // if statement can have a short statement ot execute before the condition
		return v
	} // v is only defined within the if statement's scope
    else { // if statement declaration is visible to else block as well
        return v
    }

    return lim
}


func main() {
	fmt.Println(sqrt(2), sqrt(-4))

    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
}

