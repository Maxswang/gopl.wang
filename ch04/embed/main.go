package main

import "fmt"

type Point struct{ X, Y int }

type Circle struct {
	Point  // 匿名字段
	Radius int
}

type Wheel struct {
	Circle
	Spokers int
}

func main() {
	var w Wheel
	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokers: 20,
	}

	fmt.Printf("%v\n", w)

	w.X = 42

	fmt.Printf("%#v\n", w)

}
