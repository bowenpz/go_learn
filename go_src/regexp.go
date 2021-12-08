package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func learnRegexp() {
	// API 参考：https://www.cnblogs.com/golove/p/3270918.html
	// 梳理正则历史：http://www.yitb.com/article-6490

	// —————————————— 直接匹配 ——————————————

	// 字节数组是否匹配正则
	match, _ := regexp.Match(`a(.*)c`, []byte("abc"))
	fmt.Printf("%t", match) // 打印：true

	// 字符串是否匹配正则
	matchString, _ := regexp.MatchString(`a(.*)c`, "abc")
	fmt.Printf("%t", matchString) // 打印：true

	// reader 是否匹配正则
	runeReader := new(bytes.Buffer)
	runeReader.WriteString("abc")
	matchReader, _ := regexp.MatchReader(`a(.*)c`, runeReader)
	fmt.Printf("%t", matchReader) // 打印：true

	// —————————————— 编译 ——————————————

	// 将正则表达式编译成一个正则对象（使用 PERL 语法）
	compile, _ := regexp.Compile(`a(.*)c`)
	fmt.Printf("%t", compile.MatchString("abc")) // 打印：true

	// 将正则表达式编译成一个正则对象（使用 POSIX 语法）
	compilePOSIX, _ := regexp.CompilePOSIX(`a(.*)c`)
	fmt.Printf("%t", compilePOSIX.MatchString("abc")) // 打印：true

	// 正则编译，如果编译不成功就 panic（使用 PERL 语法）
	mustCompile := regexp.MustCompile(`a(.*)c`)
	fmt.Printf("%t", mustCompile.MatchString("abc")) // 打印：true

	// 正则编译，如果编译不成功就 panic（使用 POSIX 语法）
	mustCompilePOSIX := regexp.MustCompilePOSIX(`a(.*)c`)
	fmt.Printf("%t", mustCompilePOSIX.MatchString("abc")) // 打印：true

	// —————————————— 使用正则编译对象匹配 ——————————————

}
