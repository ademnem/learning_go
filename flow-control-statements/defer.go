package main

import "fmt"


func stacking_defers() {
    fmt.Println("counting")

    for i := 0; i < 10; i++ { // defers get added to a stack
        defer fmt.Println(i)
    } 

    fmt.Println("done")
}

func main() {
	defer fmt.Println("world") // evaluated immediately
    
    stacking_defers()

	fmt.Println("hello")
} // execution deferred until function returns or terminates
