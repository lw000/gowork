// pholcus_test project main.go
package main

import (
	"fmt"

	"github.com/henrylee2cn/pholcus/exec"
)

func main() {
	fmt.Println("Hello World!")
	// 设置运行时默认操作界面，并开始运行
	// 运行软件前，可设置 -a_ui 参数为"web"、"gui"或"cmd"，指定本次运行的操作界面
	// 其中"gui"仅支持Windows系统
	exec.DefaultRun("web")
}
