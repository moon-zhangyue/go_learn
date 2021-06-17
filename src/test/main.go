package main

import (
	"fmt"
	"time"
)

//func add(a, b int) int {
//	a *= 2
//	b *= 3
//	return a + b
//}
//func main() {
//	x, y := 1, 2
//	z := add(x, y)
//	fmt.Printf("add(%d, %d) = %d\n", x, y, z)
//}

//func add(a, b *int) int {
//	*a *= 2
//	*b *= 3
//	return *a + *b
//}

//func myfunc(numbers ...int) {
//	for _, number := range numbers {
//		fmt.Println(number)
//	}
//}

//func myPrintf(args ...interface{}) {
//	for _, arg := range args {
//		switch reflect.TypeOf(arg).Kind() {
//		case reflect.Int:
//			fmt.Println(arg, "is an int value.")
//		case reflect.String:
//			fmt.Printf("\"%s\" is a string value.\n", arg)
//		case reflect.Array:
//			fmt.Println(arg, "is an array type.")
//		default:
//			fmt.Println(arg, "is an unknown type.")
//		}
//	}
//}

//func main() {
//	//x, y := 1, 2
//	//z := add(&x, &y)
//	//fmt.Printf("add(%d, %d) = %d\n", x, y, z)
//
//	//myfunc(1, 12, 3, 45)
//
//	//slice := []int{1, 2, 3, 4, 5}
//	//myfunc(slice...)
//	//myfunc(slice[1:3]...)
//
//	myPrintf(1, "1", [1]int{1}, true)
//}

//func add(a, b *int) (int, error) {
//	if *a < 0 || *b < 0 {
//		err := errors.New("只支持非负整数相加")
//		return 0, err
//	}
//
//	*a *= 2
//	*b *= 3
//	return *a + *b, nil
//}
//
//func main() {
//	x, y := 1, 2
//	z, err := add(&x, &y)
//	if err != nil {
//		fmt.Printf(err.Error())
//		return
//	}
//	fmt.Printf("add(%d, %d) = %d\n", x, y, z)
//}

//func main() {
//	// 1、将匿名函数赋值给变量
//	//add := func(a, b int) int {
//	//	return a + b
//	//}
//	//
//	//// 调用匿名函数 add
//	//fmt.Println(add(1, 2))
//	//
//	//func(a, b int) {
//	//	fmt.Println(a + b)
//	//} (1, 2)
//
//	//var j int = 1
//	//f := func() {
//	//	var i int = 1
//	//	fmt.Printf("i, j: %d, %d\n", i, j)
//	//}
//	//f()
//	//j += 2
//	//f()
//
//	add := func(a, b int) int {
//		return a + b
//	}
//	// 将函数类型作为参数
//	func(call func(int, int) int) {
//		fmt.Println(call(1, 2))
//	}(add)
//}
//
//func main() {
//	// 普通的加法操作
//	add1 := func(a, b int) int {
//		return a + b
//	}
//	// 定义多种加法算法
//	base := 10
//	add2 := func(a, b int) int {
//		return a*base + b
//	}
//	handleAdd(1, 2, add1)
//	handleAdd(1, 2, add2)
//}
//
//// 将匿名函数作为参数
//func handleAdd(a, b int, call func(int, int) int) {
//	fmt.Println(call(a, b))
//}

// 将函数作为返回值类型
//func deferAdd(a, b int) func() int {
//	return func() int {
//		return a + b
//	}
//}
//func main() {
//	// 此时返回的是匿名函数
//	addFunc := deferAdd(1, 2)
//	// 这里才会真正执行加法操作
//	fmt.Println(addFunc())
//}

//// 为函数类型设置别名提高代码可读性
//type MultiPlyFunc func(int, int) int
//
//// 乘法运算函数
//func multiply(a, b int) int {
//	return a * b
//}
//
//// 通过高阶函数在不侵入原有函数实现的前提下计算乘法函数执行时间
//func execTime(f MultiPlyFunc) MultiPlyFunc {
//	return func(a, b int) int {
//		start := time.Now()      // 起始时间
//		c := f(a, b)             // 执行乘法运算函数
//		end := time.Since(start) // 函数执行完毕耗时
//		fmt.Printf("--- 执行耗时: %v ---\n", end)
//		return c // 返回计算结果
//	}
//}
//func main() {
//	a := 2
//	b := 8
//	// 通过修饰器调用乘法函数，返回的是一个匿名函数
//	decorator := execTime(multiply)
//	// 执行修饰器返回函数
//	c := decorator(a, b)
//	fmt.Printf("%d x %d = %d\n", a, b, c)
//}

// 为函数类型设置别名提高代码可读性
type MultiPlyFunc func(int, int) int

// 乘法运算函数1（算术运算）
func multiply1(a, b int) int {
	return a * b
}

// 乘法运算函数2（位运算）
func multiply2(a, b int) int {
	return a << b
}

// 通过高阶函数在不侵入原有函数实现的前提下计算乘法函数执行时间
func execTime(f MultiPlyFunc) MultiPlyFunc {
	return func(a, b int) int {
		start := time.Now()      // 起始时间
		c := f(a, b)             // 执行乘法运算函数
		end := time.Since(start) // 函数执行完毕耗时
		fmt.Printf("--- 执行耗时: %v ---\n", end)
		return c // 返回计算结果
	}
}
func main() {
	a := 2
	b := 8
	fmt.Println("算术运算：")
	decorator1 := execTime(multiply1)
	c := decorator1(a, b)
	fmt.Printf("%d x %d = %d\n", a, b, c)
	fmt.Println("位运算：")
	decorator2 := execTime(multiply2)
	a = 1
	b = 4
	c = decorator2(a, b)
	fmt.Printf("%d << %d = %d\n", a, b, c)
}