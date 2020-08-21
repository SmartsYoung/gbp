package main

import "fmt"

func main() {
	i := a1()
	fmt.Println(i)
}

func a1() int {
	var i int
	defer func() {
		i++
	}()
	return i
}
