package main

import "fmt"


type Vertex struct { // struct is a collection of values
    X int
    Y int
}


func main() {
    v := Vertex{1, 2}
    fmt.Println(v)
    v.X = 6
    fmt.Println(v.X)
    fmt.Println(v.Y)
}
