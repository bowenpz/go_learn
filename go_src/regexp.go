package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func learnRegexp() {
	// API 参考：https://www.cnblogs.com/golove/p/3270918.html
	// 梳理正则历史：https://cloud.tencent.com/developer/article/1914673

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

	// —————————————— 其他 ——————————————

	// 将正则表达式转义成普通字符
	quoteMeta := regexp.QuoteMeta(`a(.*)b`)
	fmt.Printf("%s", quoteMeta) // 打印：a\(\.\*\)b

	// ==========================================================================
	//
	// 以下是正则编译对象（*regexp.Regexp）的方法
	// 默认情况下都是贪婪匹配
	//
	// 下面匹配规则，基本都是【a(任意多个小写字母)c】
	// 默认是 leftmost-first 模式（但是 * 符号是贪婪的，所以表现不出来）
	compile = regexp.MustCompile(`a([a-z]*)c`)
	//
	// ==========================================================================

	// —————————————— 是否匹配 ——————————————

	// 字节数组是否匹配正则
	compileMatch := compile.Match([]byte("abc"))
	fmt.Printf("%t", compileMatch) // 打印：true

	// 字符串是否匹配正则
	compileMatchString := compile.MatchString("abc")
	fmt.Printf("%t", compileMatchString) // 打印：true

	// reader 是否匹配正则
	runeReader = new(bytes.Buffer)
	runeReader.WriteString("abc")
	compileMatchReader := compile.MatchReader(runeReader)
	fmt.Printf("%t", compileMatchReader) // 打印：true

	// —————————————— 获取匹配部分 ——————————————

	// 找到字节数组中首个匹配的字符串
	compileFind := compile.Find([]byte("abcde"))
	fmt.Printf("%s", compileFind) // 打印：abc

	// 找到字符串中首个匹配的字符串
	compileFindString := compile.FindString("abcde")
	fmt.Printf("%s", compileFindString) // 打印：abc

	// 找到字节数组中，前 n 个匹配的字符串（如果 n<0，表示查询全部）
	compileFindAll := compile.FindAll([]byte("abc abbc"), -1)
	for _, findBytes := range compileFindAll {
		fmt.Printf("%s\n", findBytes)
	}
	// 打印：
	// abc
	// abbc

	// 找到字节数组中，前 n 个匹配的字符串（如果 n<0，表示查询全部）
	compileFindAllString := compile.FindAllString("abc abbc", -1)
	for _, findBytes := range compileFindAllString {
		fmt.Printf("%s\n", findBytes)
	}
	// 打印：
	// abc
	// abbc

	// —————————————— 获取匹配分组部分 ——————————————

	// 找到字节数组中，首个匹配的字符串，及其分组内容（分组指的是正则表达式中使用小括号括起的部分）
	compileFindSubmatch := compile.FindSubmatch([]byte("abc abbc"))
	for _, submatch := range compileFindSubmatch {
		fmt.Printf("%s\n", submatch)
	}
	// 打印：
	// abc
	// b

	// 找到字符串中，首个匹配的字符串，及其分组内容
	compileFindStringSubmatch := compile.FindStringSubmatch("abc abbc")
	for _, submatch := range compileFindStringSubmatch {
		fmt.Printf("%s\n", submatch)
	}
	// 打印：
	// abc
	// b

	// 找到字节数组中，前 n 个匹配的字节数组，及其分组内容（如果 n<0，表示查询全部）（返回一个三维数组）
	compileFindAllSubmatch := compile.FindAllSubmatch([]byte("abc abbc"), -1)
	for _, submatch := range compileFindAllSubmatch {
		for _, submatchBytes := range submatch {
			fmt.Printf("%s\n", submatchBytes)
		}
	}
	// 打印：
	// abc
	// b
	// abbc
	// bb

	// 找到字节数组中，前 n 个匹配的字符串，及其分组内容（如果 n<0，表示查询全部）（返回一个二维数组）
	compileFindAllStringSubmatch := compile.FindAllStringSubmatch("abc abbc", -1)
	for _, submatch := range compileFindAllStringSubmatch {
		fmt.Printf("%v\n", submatch)
	}
	// 打印：
	// [abc b]
	// [abbc bb]

	// —————————————— 获取匹配起始下标 ——————————————

	// 找到字节数组中，首个匹配部分的起始下标
	compileFindIndex := compile.FindIndex([]byte("aaa abc"))
	fmt.Printf("%v", compileFindIndex) // 打印：[4, 7]

	// 找到字符串中，首个匹配部分的起始下标
	compileFindStringIndex := compile.FindStringIndex("aaa abc")
	fmt.Printf("%v", compileFindStringIndex) // 打印：[4, 7]

	// 找到 reader 中，首个匹配部分的起始下标
	runeReader = new(bytes.Buffer)
	runeReader.WriteString("aaa abc")
	compileFindReaderIndex := compile.FindReaderIndex(runeReader)
	fmt.Printf("%v", compileFindReaderIndex) // 打印：[4, 7]

	// 找到字节数组中，前 n 个匹配字符串的起始下标（如果 n<0，表示查询全部）
	compileFindAllIndex := compile.FindAllIndex([]byte("abc abbc"), -1)
	for _, findIndex := range compileFindAllIndex {
		fmt.Printf("%v\n", findIndex)
	}
	// 打印：
	// [0 3]
	// [4 8]

	// 找到字符串中，前 n 个匹配字符串的起始下标（如果 n<0，表示查询全部）
	compileFindAllStringIndex := compile.FindAllStringIndex("abc abbc", -1)
	for _, findIndex := range compileFindAllStringIndex {
		fmt.Printf("%v\n", findIndex)
	}
	// 打印：
	// [0 3]
	// [4 8]

	// —————————————— 获取匹配分组起始下标 ——————————————

	// 找到字节数组中，首个匹配部分的起始下标，以及分组部分起始下标
	compileFindSubmatchIndex := compile.FindSubmatchIndex([]byte("abc abbc"))
	fmt.Printf("%v", compileFindSubmatchIndex) // 打印：[0 3 1 2]

	// 找到字节数组中，首个匹配部分的起始下标，以及分组部分起始下标
	compileFindStringSubmatchIndex := compile.FindStringSubmatchIndex("abc")
	fmt.Printf("%v", compileFindStringSubmatchIndex) // 打印：[0 3 1 2]

	// 找到 reader 中，首个匹配部分的起始下标，以及分组部分起始下标
	runeReader = new(bytes.Buffer)
	runeReader.WriteString("abc")
	compileFindReaderSubmatchIndex := compile.FindReaderSubmatchIndex(runeReader)
	fmt.Printf("%v", compileFindReaderSubmatchIndex) // 打印：[0 3 1 2]

	// 找到字节数组中，前 n 个匹配部分的起始下标，以及分组部分起始下标（如果 n<0，表示查询全部）（返回二维数组）
	compileFindAllSubmatchIndex := compile.FindAllSubmatchIndex([]byte("abc abbc"), -1)
	for _, submatchIndex := range compileFindAllSubmatchIndex {
		fmt.Printf("%v\n", submatchIndex)
	}
	// 打印：
	// [0 3 1 2]
	// [4 8 5 7]

	// 找到字符串中，前 n 个所有匹配部分的起始下标，以及分组部分起始下标（如果 n<0，表示查询全部）（返回二维数组）
	compileFindAllStringSubmatchIndex := compile.FindAllStringSubmatchIndex("abc abbc", -1)
	for _, submatchIndex := range compileFindAllStringSubmatchIndex {
		fmt.Printf("%v\n", submatchIndex)
	}
	// 打印：
	// [0 3 1 2]
	// [4 8 5 7]

	// —————————————— 替换匹配部分 ——————————————

	// 对字节数组，替换正则匹配的部分（可以使用 $1 等分组引用符）
	compileReplaceAll1 := compile.ReplaceAll([]byte("...abc..."), []byte("xyz"))
	compileReplaceAll2 := compile.ReplaceAll([]byte("...abc..."), []byte("$1"))
	fmt.Printf("%s %s", compileReplaceAll1, compileReplaceAll2) // 打印：...xyz... ...b...

	// 对字符串，替换正则匹配的部分（可以使用 $1 等分组引用符）
	compileReplaceAllString := compile.ReplaceAllString("...abc...", "xyz")
	fmt.Printf("%s", compileReplaceAllString) // 打印：...xyz...

	// 对字节数组，替换正则匹配的部分（字面量替换）
	compileReplaceAllLiteral := compile.ReplaceAllLiteral([]byte("...abc..."), []byte("xyz"))
	fmt.Printf("%s", compileReplaceAllLiteral) // 打印：...xyz...

	// 对字符串，替换正则匹配的部分（字面量替换）
	compileReplaceAllLiteralString := compile.ReplaceAllLiteralString("...abc...", "xyz")
	fmt.Printf("%s", compileReplaceAllLiteralString) // 打印：...xyz...

	// 对字节数组，替换正则匹配的部分，使用方法指定如何替换
	compileReplaceAllFunc := compile.ReplaceAllFunc([]byte("...abc..."), func(bs []byte) []byte {
		return []byte("$1")
	})
	fmt.Printf("%s", compileReplaceAllFunc) // 打印：...replace...

	// 对字符串，替换正则匹配的部分，使用方法指定如何替换
	compileReplaceAllStringFunc := compile.ReplaceAllStringFunc("...abc...", func(s string) string {
		return "replace"
	})
	fmt.Printf("%s", compileReplaceAllStringFunc) // 打印：...replace...

	// Expand 是 replace 的升级，需要配合 FindSubmatchIndex 方法一起使用
	// 它可以把正则匹配的分组字符，填充到模板里
	//
	// dst		: Expand 返回文本的开头
	// template	: 替换模板，匹配分组结果，会按照这个模板填充
	// src		: 源字符串，用于正则匹配，获取分组字符
	// match	: FindSubmatchIndex 的匹配结果
	var (
		src           = []byte("foo")
		dst           = []byte("这是开头，")
		template      = []byte("这是替换后的结果，解析出文本1($1)，解析出文本2($2)")
		expandCompile = regexp.MustCompile("([a-z]*) ([a-z]*)")
	)
	expandIndex := expandCompile.FindSubmatchIndex(src)
	expand := expandCompile.Expand(dst, template, src, expandIndex)
	fmt.Printf("%s", expand) // 打印：这是开头，这是替换后的模板，解析出文本1(foo)，解析出文本2(bar)

	// Expand 方法的字符串版本，需要配合 FindStringSubmatchIndex 方法一起使用
	var (
		srcStr           = "foo bar"
		templateStr      = "这是替换后的结果，解析出文本1($1)，解析出文本2($2)"
		expandStrCompile = regexp.MustCompile("([a-z]*) ([a-z]*)")
	)
	expandStrIndex := expandStrCompile.FindStringSubmatchIndex(srcStr)
	expandStr := expandCompile.ExpandString(dst, templateStr, srcStr, expandStrIndex)
	fmt.Printf("%s", expandStr) // 打印：这是开头，这是替换后的模板，解析出文本1(foo)，解析出文本2(bar)

	// —————————————— 其他 ——————————————

	// 正则匹配字符串本身
	compileString := compile.String()
	fmt.Printf("%s", compileString) // 打印：a([a-z]*)c

	// 以匹配内容为界，将字符串分割成多个字符串，n 表示分割后的字符串数量，n<0 表示不限数量
	compileSplit := compile.Split("...abc,,,", -1)
	fmt.Printf("%v", compileSplit) // 打印：[... ,,,]

	// 返回正则表达式必须匹配到的字面前缀
	// 如果整个正则表达式都是字面值，则 complete 返回 true
	compilePrefix, compileComplete := compile.LiteralPrefix()
	fmt.Printf("%s %t", compilePrefix, compileComplete) // 打印：a false

	// Longest 将正则对象变成 leftmost-longest 模式
	r := regexp.MustCompile("fo+?")
	fmt.Printf("%v", r.FindStringSubmatch("foooooo")) // 打印：[fo]
	r.Longest()
	fmt.Printf("%v", r.FindStringSubmatch("foooooo")) // 打印：[foooooo]

	// 正则表达式分组的数量
	compileNumSubexp := compile.NumSubexp()
	fmt.Printf("%d", compileNumSubexp) // 打印：1

	// 正则表达式分组的名字（第一个永远是空字符串，代表整体）
	r = regexp.MustCompile("(?P<key1>[a-z]+) (?P<key2>[a-z]+)")
	compileSubexpNames := r.SubexpNames()
	for _, name := range compileSubexpNames {
		fmt.Printf("[%s]\n", name)
	}
	// 打印：
	// []
	// [key1]
	// [key2]

	// 正则表达式分组名字的下标
	compileSubexpIndex := r.SubexpIndex("key2")
	fmt.Printf("%d", compileSubexpIndex) // 打印：2

	// 【已废弃】拷贝正则对象，用于避免早期版本多 goroutine 同时使用时上锁造成的开销
	copyCompile := compile.Copy()
	fmt.Printf("%t", copyCompile.MatchString("abc")) // 打印：true
}
