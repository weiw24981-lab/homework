package main

import (
	"fmt"
)

// func main() {
// 	m := map[string]interface{}{"kkk": 111}
// 	emptyInterfaceDemo(m["kkk"])
// }

func emptyInterfaceDemo(value interface{}) {
	fmt.Printf("\n值: %v, 类型: %T\n", value, value)

	// 类型断言
	if str, ok := value.(string); ok {
		fmt.Println("是字符串:", str)
	}

	// type switch
	switch v := value.(type) {
	case int:
		fmt.Println("整数:", v)
	case string:
		fmt.Println("字符串:", v)
	case []int:
		fmt.Println("整数切片:", v)
	case map[string]int:
		fmt.Println("map[string]int{}:", v)
	case map[string]float32:
		fmt.Println("map[string]int{}:", v)
	case map[string]interface{}:
		fmt.Println("map[string]interface{}:", v)
	default:
		fmt.Println("未知类型")
	}
}
