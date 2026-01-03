package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	School string
}

type Point struct{ X, Y int }

func main() {

	p := Point{1, 2}

	fmt.Println(p.X, p.Y)
	fmt.Println(p)
}
