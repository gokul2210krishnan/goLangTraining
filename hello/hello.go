package main

import "fmt"

func main() {
	n, _ := fmt.Println("Hello world") // underscore character to throw away returns
	n, e := fmt.Println("Hello world")
	fmt.Println(n)
	fmt.Println(e) // incase of no error <nil> will be the output
}

// using variadic parameter in a func
func sample(a ...bool) {

}
