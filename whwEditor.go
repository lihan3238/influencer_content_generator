package main

import (
	"fmt"

	"github.com/lihan3238/influencer_content_generator/Tools"
)

func main() {
	fmt.Println("\x1b[36m网红文生成器\x1b[0m\n") // Cyan color for the title

	fmt.Println("模板:\n")
	fmt.Println("\x1b[32m" + Tools.ShowTemplate() + "\x1b[0m\n") // Green color for the template

	fmt.Println("原文:\n")
	fmt.Println("\x1b[34m" + Tools.ShowOriginal() + "\x1b[0m\n") // Blue color for the original text

	fmt.Println("编辑:\n")

	var replacements = map[string]*string{
		"[人物A]":    new(string),
		"[某事]":     new(string),
		"[地点]":     new(string),
		"[某人/物]":   new(string),
		"[人物C]":    new(string),
		"[人物D]":    new(string),
		"[人物E]":    new(string),
		"[某文件/内容]": new(string),
		"[身份]":     new(string),
		"[人物B]":    new(string),
		"[人物F]":    new(string),
		"[人物A的组织]": new(string),
		"[刑罚]":     new(string),
	}

	fmt.Println("请依次填充内容:\n")
	fmt.Print("\x1b[33m人物A:\x1b[0m\n") // Yellow color for user prompts
	fmt.Scanln(replacements["[人物A]"])
	fmt.Print("\x1b[33m某事:\x1b[0m\n")
	fmt.Scanln(replacements["[某事]"])
	fmt.Print("\x1b[33m地点:\x1b[0m\n")
	fmt.Scanln(replacements["[地点]"])
	fmt.Print("\x1b[33m某人/物:\x1b[0m\n")
	fmt.Scanln(replacements["[某人/物]"])
	fmt.Print("\x1b[33m人物C:\x1b[0m\n")
	fmt.Scanln(replacements["[人物C]"])
	fmt.Print("\x1b[33m人物D:\x1b[0m\n")
	fmt.Scanln(replacements["[人物D]"])
	fmt.Print("\x1b[33m人物E:\x1b[0m\n")
	fmt.Scanln(replacements["[人物E]"])
	fmt.Print("\x1b[33m某文件/内容:\x1b[0m\n")
	fmt.Scanln(replacements["[某文件/内容]"])
	fmt.Print("\x1b[33m身份:\x1b[0m\n")
	fmt.Scanln(replacements["[身份]"])
	fmt.Print("\x1b[33m人物B:\x1b[0m\n")
	fmt.Scanln(replacements["[人物B]"])
	fmt.Print("\x1b[33m人物F:\x1b[0m\n")
	fmt.Scanln(replacements["[人物F]"])
	fmt.Print("\x1b[33m人物A的组织:\x1b[0m\n")
	fmt.Scanln(replacements["[人物A的组织]"])
	fmt.Print("\x1b[33m刑罚:\x1b[0m\n")
	fmt.Scanln(replacements["[刑罚]"])

	replacementStrings := make(map[string]string)
	for key, value := range replacements {
		replacementStrings[key] = *value
	}

	fmt.Println("\x1b[31m" + Tools.Edit(replacementStrings) + "\x1b[0m\n") // Red color for the edited text
	fmt.Print("按下回车键继续...")
	fmt.Scanln()
}
