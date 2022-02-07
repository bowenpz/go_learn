package main

import (
	"bytes"
	"fmt"
)

// Package io provides basic interfaces to I/O primitives.
// Its primary job is to wrap existing implementations of such primitives,
// such as those in package os, into shared public interfaces that
// abstract the functionality, plus some other related primitives.
// io 包提供了 I/O 原语的基本接口。它的首要任务是提供 io 原语的包装实现，例如把 os 包中的 io 原语，转换成公共抽象接口，并增加一些相关原语。
//
// Because these interfaces and primitives wrap lower-level operations with
// various implementations, unless otherwise informed clients should not
// assume they are safe for parallel execution.
// 因为这些接口和原语包装的都是原始操作，且实现众多，因此除非有声明，否则使用时不应假设它们并发安全。

// 参考：https://www.cnblogs.com/golove/p/3276678.html
func learnIO() {
	buffer := bytes.NewBuffer([]byte{})

	// 数组 p1 -> 自身🗂
	p1 := []byte("abc")
	writeNum, err := buffer.Write(p1)
	fmt.Printf("【written bytes num: %d】,【buffer: %s】,【err: %s】", writeNum, buffer.String(), err) // 打印：【written bytes num: 3】,【buffer: abc】,【err: %!s(<nil>)】

	// 自身🗂 -> 数组 p2
	p2 := []byte("wxyz")
	readNum, err := buffer.Read(p2)
	fmt.Printf("【read bytes num: %d】,【p2: %s】,【buffer: %s】,【err: %s】", readNum, p2, buffer.String(), err) // 打印：【read bytes num: 3】,【p2: abcz】,【buffer: 】,【err: %!s(<nil>)】

	// Reader r -> 自身🗂
	r := bytes.NewBuffer([]byte("long long long long buffer"))
	readFromNum, err := buffer.ReadFrom(r)
	fmt.Printf("【read from bytes num: %d】,【buffer: %s】,【r: %s】,【err: %s】", readFromNum, buffer.String(), r.String(), err) // 打印：【read from bytes num: 26】,【buffer: long long long long buffer】,【r: 】,【err: %!s(<nil>)】

	// 自身🗂 -> Writer w
	w := bytes.NewBuffer([]byte{})
	writeToNum, err := buffer.WriteTo(w)
	fmt.Printf("【write to bytes num: %d】,【buffer: %s】,【w: %s】,【err: %s】", writeToNum, buffer.String(), w.String(), err) // 打印：【write to bytes num: 26】,【buffer: 】,【w: long long long long buffer】,【err: %!s(<nil>)】

	// 自身🗂 -> 1 byte
	buffer.Write([]byte("hello"))
	for buffer.Len() > 0 {
		readByte, err := buffer.ReadByte()
		fmt.Printf("【read byte: %c】,【err: %s】\n", readByte, err)
		// 打印：
		//【read byte: h】,【err: %!s(<nil>)】
		//【read byte: e】,【err: %!s(<nil>)】
		//【read byte: l】,【err: %!s(<nil>)】
		//【read byte: l】,【err: %!s(<nil>)】
		//【read byte: o】,【err: %!s(<nil>)】
	}

	// 1 byte -> 自身🗂
	for i := 0; i < 5; i++ {
		buffer.WriteByte('a')
	}
	fmt.Printf("【buffer: %s】", buffer.String()) // 打印：【buffer: aaaaa】

	// 自身🗂 -> 1 rune
	buffer.Reset()
	buffer.Write([]byte("你好呀"))
	readRune, size, err := buffer.ReadRune()
	fmt.Printf("【read rune: %c】,【size: %d】,【buffer: %s】,【err: %s】", readRune, size, buffer.String(), err) // 打印：【read rune: 你】,【size: 3】,【buffer: 好呀】,【err: %!s(<nil>)】

	// 字符串 -> 自身🗂
	buffer.Reset()
	writeStringNum, err := buffer.WriteString("hello")
	fmt.Printf("【write string num: %d】,【buffer: %s】,【err: %s】", writeStringNum, buffer.String(), err) // 打印：【write string num: 5】,【buffer: hello】,【err: %!s(<nil>)】
}
