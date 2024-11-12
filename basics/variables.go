package main

import "fmt"

var c, python, java bool // var DECLARATION
// the type is always last
// can be at package or function level

var i, j = 1, 2 // var INITIALIZER
// type can be omitted since the variable will take the type of the initializer

var str, num = "hello", 1 // initializer list with var means multiple variables are initialized
// declaring the type is faster right? or does it not matter?
// declaring the type does not matter in terms of speed, the Go compiler is highly efficient
// declaring the type is good for readabiltiy though, especially is someone does not know Go

func main() {
	var a int // default val of int is 0
	fmt.Println(a, c, python, java) // default val of bool is false
    // passing multiple args to Println yields each val with space inbetween
    fmt.Println(i, j)

    fmt.Println(str, num) 

    // short assignment := can be used in place of var declaration with implicit type
    c, python, java := true, false, "no!"
    // outside a function, every statement begins with a keyword (var, func, etc.) so the := is not available
    fmt.Println(c, python, java)
}

