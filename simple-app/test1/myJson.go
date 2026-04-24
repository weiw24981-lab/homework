package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func doJsonDemo() {
	p := Person{Name: "haha", Age: 33}
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("序列化失败")
	} else {
		fmt.Printf("序列化成功%v\n", string(jsonData))
	}

	jsonData2, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		fmt.Printf("序列化失败")
	} else {
		fmt.Printf("序列化成功%v\n", string(jsonData2))
	}

	var p2 Person
	if err := json.Unmarshal(jsonData2, &p2); err != nil {
		fmt.Printf("反序列化失败")
	} else {
		fmt.Printf("反序列化成功%v\n", p2)
	}
}
