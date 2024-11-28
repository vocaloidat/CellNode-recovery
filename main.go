package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// CommandResult 结构体存储命令和执行结果
type CommandResult struct {
	Command string
	Output  string
	Error   error
}

// executeCommand 执行给定命令并返回其输出
func executeCommand(command string) CommandResult {
	cmd := exec.Command("bash", "-c", command)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return CommandResult{Command: command, Output: stderr.String(), Error: err}
	}
	return CommandResult{Command: command, Output: out.String(), Error: nil}
}

func main() {
	// 要执行的命令列表
	commands := []string{
		"systemctl status cellframe-node.service",
		"/opt/cellframe-node/bin/cellframe-node-cli node dump -net Backbone",
		"/opt/cellframe-node/bin/cellframe-node-tool cert pkey show node-cert",
		"/opt/cellframe-node/bin/cellframe-node-cli version",
		"/opt/cellframe-node/bin/cellframe-node-cli srv_stake order list -net Backbone | tail -n 50",
		"/opt/cellframe-node/bin/cellframe-node-cli net -net Backbone get status",
		"/opt/cellframe-node/bin/cellframe-node-cli node list -net Backbone",
	}

	results := make([]CommandResult, len(commands))

	// 执行每个命令
	for i, command := range commands {
		results[i] = executeCommand(command)
	}

	// 汇总输出
	fmt.Println("==== 汇总结果 ====")
	for _, result := range results {
		fmt.Printf("执行命令: %s\n", result.Command)
		if result.Error != nil {
			fmt.Printf("错误: %v\n", result.Error)
		}
		fmt.Printf("输出:\n%s\n", strings.TrimSpace(result.Output))
		fmt.Println("=======================================")
	}
}
