package main

import "fmt"

const Pi = 3.14159

func main() {
	// const pi = 3.14159

	// 在一条const语句 定义多个 const常量
	const (
		e  = 2.71828
		pi = 3.14159
	)

	// fmt.Println("常量pi值", pi)

	fmt.Println("hello")

	println("hello world") // 内置 println/print函数 仅常用于调试使用

}
