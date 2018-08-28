package main

import "fmt"

func main() {
	a := make([]int,0)
	fmt.Printf("%v\n",a)
	a = append(a,1)
	fmt.Printf("%v\n",a)
	var b map[int]int
	fmt.Printf("%v\n",b)
	if b == nil {
		fmt.Printf("b is nil.\n" )
	}
	b = make(map[int]int,2)
	if b == nil {
		fmt.Printf("b is nil.\n" )
	}
	fmt.Printf("%v\n",b)
}
