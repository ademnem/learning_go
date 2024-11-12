package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ { // init statemnt, condition expression, post statements
		sum += i
	} 
    // NOTICE: there are no () around for loop
    // {} are always required

    // you can still add the () if you want
    // init and post statemnts are optional
    for(; sum < 1000;) {
        sum += sum
    } 

    // WHILE LOOP
    // if you drop the semi-colons, the for loop becomes a while loop
    for sum < 1000 {
        sum += 2;
    } 

        
    for { // this makes an infinite loop like a while(true)
        fmt.Println("infinite loop")
    }

	fmt.Println(sum)
}

