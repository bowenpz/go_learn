package main

import (
	"fmt"
	"strconv" // 1.16.5
)

func main() {
	// 这篇文章总结的挺好的: https://www.cnblogs.com/golove/p/3262925.html

	// 反转义解码字符串的首个字符
	//
	// s	: 待解码的字符串
	// quote: 转义字符，可以是单引号（'\''）或双引号（'"'），否则表示没有转义字符
	//
	// value	: 首个解码出的字符
	// multibyte: 该字符是否是多字节的（比如'a'是单字节，'我'是多字节）
	// tail		: 还未解码的剩余字符串
	// err		: error，目前只有一种，就是解析错误（invalid syntax）
	//
	// 下面代码以解析【\"我\"】为例
	// 将这个转义字符串，逐个字符地解析出来
	var (
		value     rune
		multibyte bool
		tail      string = `\"我\"`
		err       error
	)
	for len(tail) > 0 {
		value, multibyte, tail, err = strconv.UnquoteChar(tail, '"')
		if err != nil {
			fmt.Printf("err: %s", err)
			break
		}
		fmt.Printf("解码字符：%c\t是否多字节：%t\t剩余字符串：%s\n", value, multibyte, tail)
	}
	// 打印：
	// 解码字符："     是否多字节：false       剩余字符串：我\"
	// 解码字符：我    是否多字节：true        剩余字符串：\"
	// 解码字符："     是否多字节：false       剩余字符串：

	// 字符是否可显示（可显示的定义，跟通常所想并不一样）
	isPrint1 := strconv.IsPrint('a')
	isPrint2 := strconv.IsPrint(' ')
	isPrint3 := strconv.IsPrint('\t')
	fmt.Printf("%t %t %t", isPrint1, isPrint2, isPrint3) // 打印：true true false

	// 字符是否是一个 Unicode 图形字符
	isGraphic1 := strconv.IsGraphic('a')
	isGraphic2 := strconv.IsGraphic('\t')
	fmt.Printf("%t %t", isGraphic1, isGraphic2) // 打印：true false

	// 字符串是否可以不被修改地表示为一个单行的反引号字符串
	canBackquote1 := strconv.CanBackquote("abc")
	canBackquote2 := strconv.CanBackquote("`abc`")
	fmt.Printf("%t %t", canBackquote1, canBackquote2) // 打印：true, false

	itoa := strconv.Itoa(10)
	fmt.Printf("%s", itoa) // 打印：10
}
