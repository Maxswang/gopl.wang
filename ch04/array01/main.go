package main

import "fmt"

func main() {

	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	// 打印出数组的类型
	fmt.Printf("%T\n", a)

	// 打印出数组元素的内存地址
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &a[0])
	fmt.Printf("%p\n", &a[1])
	fmt.Printf("%p\n", &a[2])

	for i, v := range a {
		fmt.Printf("%d\t%d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// 使用 字面量 来初始化数组
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q)
	fmt.Println(r[2])

	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3} // 省略号... 表示 让Go自动推断数组的长度
	arr3 := [...]int{99: -1}  // 初始化时指定 索引来初始化， 其它未指定索引的元素使用元素类型的零值初始化

	fmt.Printf("%T\n",arr1)
	fmt.Printf("%T\n",arr2)
	fmt.Printf("%T\n",arr3)

}
