package main

import "fmt"
import "rsc.io/quote"


func main() {
	fmt.Printf("test string %s\n", "testando")
	fmt.Printf("test string %s\n", "testando")
	fmt.Printf("test string %s\n", quote.Go())
}