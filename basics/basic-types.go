package main

// "factored" import
import (
	"fmt"
	"math/cmplx"
)


// "factored" var
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i) // you can use scientific notation?!
)
/* basic types
bool
string
int (8 16 32 64)
uint (8 16 32 64 ptr)
byte (alias for unit8)
rune (alias for int32, for unicode codepoint)
float32, float64
complex64, complex128
*/


func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)


    var (
        i int
        f float64
        b bool
        s string
    )

    fmt.Printf("%v %v %v %q\n", i, f, b, s) // this is like printf from C
    // %v default format value
    // %q quoted string, suitable for string literal

    fmt.Printf("Empty string using %%v: %v\n", s)
    fmt.Printf("Empty string using %%q: %q\n", s)

    s = "hello"
    fmt.Printf("String using %%v: %v\n", s)
    fmt.Printf("String using %%q: %q\n", s)


    // TYPE CONVERSION 
    in := 42
    fl := float64(in)
    ui := uint(fl)
    st := string(ui) // does not work

    fmt.Println(in, fl, ui, st) // st comes out as a *


    v := 42
    fmt.Printf("v is of type %T\n", v)
    
    a := 3.142
    comp := 0.867 + 0.5i
    fmt.Printf("a is of type %T\n", a)
    fmt.Printf("comp is of type %T\n", comp)


    // CONSTANTS 
    const Pi = 3.14 
    const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth) 
}

