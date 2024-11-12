package main

import (
	"fmt"
	"runtime"
    "time"
)

func main() {
	fmt.Print("Go runs on ")

    // switch cases automatically have the break at the end
    // switch cases don't have to be integers
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

    fmt.Println()

    fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	    case today + 0:
	    	fmt.Println("Today.")
	    case today + 1:
	    	fmt.Println("Tomorrow.")
	    case today + 2:
	    	fmt.Println("In two days.")
	    default:
	    	fmt.Println("Too far away.")
	}

    fmt.Println()

    t := time.Now()
    // switch case without condition == switch true
	switch {
	    case t.Hour() < 12:
	    	fmt.Println("Good morning!")
	    case t.Hour() < 17:
	    	fmt.Println("Good afternoon.")
	    default:
	    	fmt.Println("Good evening.")
	} // good way to write long if-else statements
}

