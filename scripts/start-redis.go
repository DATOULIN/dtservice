package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var out bytes.Buffer
	dir, err := os.Getwd() // 获取当前目录路径
	if err != nil {
		fmt.Println("获取当前目录失败：", err)
		return
	}
	filePath := dir + "\\scripts\\startredis.bat"  // 拼接.bat文件的路径
	cmd := exec.Command("cmd.exe", "/C", filePath) // 执行命令
	cmd.Stdout = &out
	err = cmd.Run() // 执行命令
	if err != nil {
		fmt.Println("打开文件失败：", err)
		return
	}
	fmt.Printf("%s", out.String())
}
