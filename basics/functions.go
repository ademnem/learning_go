package main

import "fmt"

func add(x int, y int) int { // type comes after the variable, this is true for the functor as well
	return x + y // notice how there are no semi-colons
}

func sub(x, y int) int { // when connsecutive named function parameters share a type you can omit the type from everything but the last
    return x - y
}

func swap(x, y string) (string, string) { // a function can return any number of results
    return y, x
}

func split(sum int) (x, y int) { // named return values
    x = sum * 4 / 9
    y = sum - x
    return // "naked" return (returns named return values)
    // should only be used in short functions where the variables are easy to find
    // readability gets worse with "naked" returns
}

func main() {
	fmt.Println(add(42, 13)) // fmt implements the I/O functionalities

    fmt.Println(sub(50, 25)) 

    a, b := swap("hello", "world") // is assignment :=
    fmt.Println(a, b)

    fmt.Println(split(17))
}

