package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(hypot(3, 4))
	f(1, 2, 3, "hello ", "world")

	// Printf %T 任何值的类型
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)

}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

// func f(i int,j int,k int, s string,t string) {
func f(i, j, k int, s, t string) { // 与上一行代码等价， 形参列表的参数和返回值列表中如果是相同类型，可以合并只写一次
	fmt.Println(i, j, k, s, t)
}


// 
func add(x int, y int) int {
	return x + y
}

func sub(x, y int) (z int) {
	z = x - y
	return
}

func first(x int, _ int) int {
	return x
}

func zero(int, int) int {
	return 0
}
