package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func FilePathDemo() {
	fmt.Println("=== 文件读写 ===")

	// runtime.Caller(0) 返回当前函数所在的源文件路径，从而可以定位到与该 Go 文件同级的目录。
	// 这样无论从哪里运行程序，都能在源代码所在目录进行读写，方便查看生成的示例文件。
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("获取当前路径失败\n")
		return
	}
	dir := filepath.Dir(filename)
	inputPath := filepath.Join(dir, "myInput.txt")
	outputPath := filepath.Join(dir, "myOutput.txt")

	if err := os.WriteFile(inputPath, []byte("Hello myinput file111"), 0o644); err != nil {
		fmt.Printf("写入文件失败: %v\n", err)
	}
	content, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
	}
	fmt.Printf("读取到的内容: %s\n", string(content))

	outFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
	}
	defer outFile.Close()
	writer := bufio.NewWriter(outFile)
	lines := []string{"射雕英雄传", "神雕侠侣", "倚天屠龙记"}
	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			fmt.Printf("写入文件失败: %v\n", err)
		}
	}
	if err := writer.Flush(); err != nil {
		fmt.Printf("刷新缓冲区失败: %v\n", err)
	}
	fmt.Println("写入文件:", outputPath)

	// 逐行读取文件内容
	file, err := os.Open(outputPath)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fmt.Println("逐行读取文件内容:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("扫描文件失败: %v\n", err)
	}

}
