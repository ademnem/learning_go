package main // tells program this is package main
// package main is run first when program is started
// go is made up of packages

// packages can be grouped in a "factored" import statement
import (
	"fmt"
	"math/rand"
)

/* or they can be individual
improt "fmt"
import "math/rand"
*/

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
