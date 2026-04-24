package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// basicTypeDemo()
	// arrayDemo()
	// sliceDemo()
	mapDemo()
}

func basicTypeDemo() {
	fmt.Printf(" basic type start \n")

	//布尔类型
	var isTrue bool = true
	var isFalse bool = false
	fmt.Printf(" booltype-- isTrue: %t  isFalse: %t\n", isTrue, isFalse)

	//有符号整数
	var num1 int = 234       //系统32位,就是int32;64位就是int64
	var num2 int8 = 127      //-128-127
	var num3 int16 = 2222    //-32768-32767
	var num4 int32 = 2343434 //-2^32-2^32-1
	var num5 int64 = -232323 //-2^63-2^63-1
	fmt.Printf(" inttype-- int:%d  int8:%d  int16:%d  int32:%d  int64:%d\n", num1, num2, num3, num4, num5)

	//无符号整数
	var unum1 uint = 123           //系统32位,就是uint32;64位就是uint64
	var unum2 uint8 = 255          //0-255
	var unum3 uint16 = 65535       //0-65535
	var unum4 uint32 = 234324      //0-2^32-1
	var unum5 uint64 = 34534534543 //0-^64-1
	fmt.Printf(" uinttype-- uint:%d  uint8:%d  uint16:%d  uint32:%d  uint64:%d\n", unum1, unum2, unum3, unum4, unum5)

	//类型别名
	var b byte = 65  //byte是uint8别名
	var r rune = '哈' //rune是int32别名
	fmt.Printf(" byte: %d (%c)  rune:%d （%c）", b, b, r, r)

	//显示类型占用内存的大小
	fmt.Printf("\n类型占用内存大小\n")
	fmt.Printf("bool size:%d bytes\n", unsafe.Sizeof(isTrue))
	fmt.Printf("int size:%d bytes\n", unsafe.Sizeof(num1))
	fmt.Printf("int8 size:%d bytes\n", unsafe.Sizeof(num2))
	fmt.Printf("int16 size:%d bytes\n", unsafe.Sizeof(num3))
	fmt.Printf("int32 size:%d bytes\n", unsafe.Sizeof(num4))
	fmt.Printf("int64 size:%d bytes\n", unsafe.Sizeof(num5))
	fmt.Printf("uint size:%d bytes\n", unsafe.Sizeof(unum1))
	fmt.Printf("uint8 size:%d bytes\n", unsafe.Sizeof(unum2))
	fmt.Printf("uint16 size:%d bytes\n", unsafe.Sizeof(unum3))
	fmt.Printf("uint32 size:%d bytes\n", unsafe.Sizeof(unum4))
	fmt.Printf("uint64 size:%d bytes\n", unsafe.Sizeof(unum5))

	//浮点型
	var price float32 = 88.88
	var pie float64 = 3.141592679
	fmt.Printf("floattype-- float32:%.2f,float64:%.10f\n", price, pie)
	fmt.Printf("float32 size:%d,float64 size %d\n", unsafe.Sizeof(price), unsafe.Sizeof(pie))

	//字符串类型
	var name string = "hello friend"
	fmt.Printf("stringtype-- string:%s\n", name)
	fmt.Printf("字符串长度：%d(bytes)\n", len(name))
	var firstbyte byte = name[0]
	fmt.Printf("第一个字节：%d(%c)\n", firstbyte, firstbyte)

	//原始字符串字面量
	var raw string = `这是
	一个多行
	字符串`
	fmt.Printf("原始字符串:" + raw)

	var char rune = '哈'
	fmt.Printf("\n字符rune:%c\n", char)

	// 复数类型
	var c1 complex64 = 1 + 2i
	var c2 complex128 = 3.14 + 6.28i
	c3 := complex(5.0, 10.0) // 使用complex函数创建

	fmt.Println("\n复数类型:")
	fmt.Printf("complex64: %v, 实部: %.1f, 虚部: %.1f\n", c1, real(c1), imag(c1))
	fmt.Printf("complex128: %.2f, 实部: %.2f, 虚部: %.2f\n", c2, real(c2), imag(c2))
	fmt.Printf("使用complex函数: %v\n", c3)
	fmt.Printf("complex64 size: %d bytes, complex128 size: %d bytes\n", unsafe.Sizeof(c1), unsafe.Sizeof(c2))
}

func arrayDemo() {
	fmt.Printf("数组示例\n")
	var arr [5]int
	fmt.Printf("声明数组但不初始化 %v\n", arr)

	arr = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("初始化数组 %v\n", arr)
	fmt.Printf("数组元素[0] %d\n", arr[0])

	//自动推断长度
	arr1 := [...]int{1, 2, 3}
	fmt.Printf("自动长度:%v,长度%d\n", arr1, len(arr1))

	// var arr2 [5]int = [5]int{1,2,3,4,5}
}

func sliceDemo() {
	fmt.Printf("切片示例")

	//声明一个切片,和声明数组的区别在于声明数组需要初始化长度,而切片不需要
	var slice []int
	fmt.Printf("空切片:%v,nil:%t\n", slice, slice == nil)

	//使用make创建切片
	slice = make([]int, 3, 5)
	fmt.Printf("make创建:%v,长度%d,容量%d\n", slice, len(slice), cap(slice))

	//初始化切片
	slice = []int{1, 2, 3, 4, 5}
	fmt.Printf("初始化slice:%v\n", slice)

	//追加元素
	slice = append(slice, 6)
	fmt.Printf("追加后的切片:%v,长度%d,容量%d\n", slice, len(slice), cap(slice))

	//切片截取
	var subSlice []int = slice[1:3]
	fmt.Printf("截取切片:%v\n", subSlice)

	// 切片共享底层数组， 切片后的subSlice，和slice共享同一个底层数组，所以修改subSlice会影响slice
	subSlice[0] = 999
	fmt.Printf("修改subSlice后: %v\n", slice)

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice1:%v\n", slice1)

}

// 映射示例
func mapDemo() {
	fmt.Printf("\n=== 映射示例 ===\n")
	//声明映射
	var m map[string]int
	fmt.Printf("声明映射但不初始化 m: %v,nil:%t\n", m, m == nil)

	//初始化映射
	m = make(map[string]int)
	m["apple"] = 100
	m["balana"] = 200
	fmt.Printf("初始化映射后 m:%v\n", m)

	//直接初始化映射
	m1 := map[string]int{
		"apple":  100,
		"balana": 200,
		"orange": 300,
	}
	fmt.Printf("m1:%v\n", m1)

	// 读取值
	value := m1["apple"]
	fmt.Printf("apple的值: %d\n", value)
	value1 := m1["orange"]
	fmt.Printf("orange的值: %d\n", value1)

	m1["orange"] = 100
	fmt.Printf("orange的值: %d\n", m1["orange"])

	// // 检查key是否存在
	value, ok := m1["grape"]
	fmt.Printf("检查key是否存在 %t,值是:%d\n", ok, value)
	value1, ok1 := m1["orange"]
	fmt.Printf("检查key是否存在 %t,值是:%d\n", ok1, value1)

	delete(m1, "balana")
	fmt.Printf("删除后: %v\n", m1)

	// // 遍历映射
	fmt.Println("遍历映射:")
	for key, value := range m1 {
		fmt.Printf("  %s: %d\n", key, value)
	}

}

func pointerDemo() {
	fmt.Println("\n=== 指针示例 ===")

	x := 10
	fmt.Printf("x的值为: %d\n", x)

	// 获取地址
	p := &x
	fmt.Printf("x的地址: %p\n", p)
	fmt.Printf("指针的值: %d\n", *p)

	m := p
	fmt.Printf("m: %p\n", m)
	fmt.Printf("m指向的值: %d\n", *m)

	// // 通过指针修改值
	*p = 20
	fmt.Printf("修改后x的值为: %d\n", x)

	// // 指针作为函数参数
	increment(&x)
	fmt.Printf("函数修改后x的值为: %d\n", x)

	// // nil指针
	var ptr *int
	fmt.Printf("nil指针: %v\n", ptr)
	// 解引用nil指针会导致panic
	// fmt.Println(*ptr) // 这行会panic

	// // ===== 值传递 vs 指针传递 =====
	fmt.Println("\n=== 值传递 vs 指针传递 ===")

	// 值传递示例
	a := 10
	fmt.Printf("调用前 a = %d\n", a)
	valuePass(a)
	fmt.Printf("调用后 a = %d (值未改变)\n", a)

	// 指针传递示例
	b := 10
	fmt.Printf("\n调用前 b = %d\n", b)
	pointerPass(&b)
	fmt.Printf("调用后 b = %d (值已改变)\n", b)

	// // 详细说明
	fmt.Println("\n关键区别：")
	fmt.Println("1. 值传递：函数接收的是值的副本，修改副本不影响原值")
	fmt.Println("2. 指针传递：函数接收的是地址，通过地址可以直接修改原值")
}

func increment(p *int) {
	*p++ // 修改指针指向的值
}

// 值传递：函数接收的是值的副本
func valuePass(num int) {
	fmt.Printf("  函数内接收到的值: %d\n", num)
	num = 100 // 修改副本，不影响原值
	fmt.Printf("  函数内修改后: %d\n", num)
}

// 指针传递：函数接收的是地址
func pointerPass(num *int) {
	fmt.Printf("  函数内接收到的地址: %p\n", num)
	fmt.Printf("  函数内接收到的值: %d\n", *num)
	*num = 100 // 通过指针修改原值
	fmt.Printf("  函数内修改后: %d\n", *num)
}
