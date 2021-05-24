package main

import "fmt"

func main() {
	component := GetInstance()
	component = GetInstance()
	// Call Count : 2, Init Count : 1
	fmt.Printf("Call Count : %d, Init Count : %d\n", component.Count, component.InitCount)
}
