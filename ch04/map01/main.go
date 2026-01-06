package main

import "fmt"

func main() {

	// map 创建方式一： 使用 make
	ages := make(map[string]int)
	ages["wang"] = 10
	ages["zhang"] = 15

	// map 创建方式二：
	ages1 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	//   上面 ages1 与 下面 ages1 定义 两者等价
	// ages1 := make(map[string]int)
	// ages1["alice"] = 31
	// ages1["charlie"] = 34

	// 对map 已有的key值 赋值会覆盖原有的value值
	// 通过 常规下标 来访问 map的元素
	ages1["alice"] = 28
	fmt.Println(ages1["alice"])

	// 先创建一个空的 map 再新增元素
	emp := map[string]int{}
	emp["mike"] = 15
	emp["jerry"] = 20

	// map创建方式三
	var a map[string]int
	a = make(map[string]int)
	a["wangs"] = 8

	// 不能对 map元素 进行取地址操作
	// _ := &a

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// map的零(zero)值为 nil
	var ages2 map[string]int
	fmt.Println(ages2 == nil)
	fmt.Println(len(ages2) == 0)
	delete(ages2, "hello")

	eq := equal(map[string]int{"A": 0}, map[string]int{"B": 42})
	fmt.Println(eq)

}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
		// if xv != y[k] { 
			return false
		}
	}
	return true
}
