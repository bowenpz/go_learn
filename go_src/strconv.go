package main

import (
	"fmt"
	"strconv" // 1.16.5
)

func main() {
	// 这篇文章总结的挺好的: https://www.cnblogs.com/golove/p/3262925.html

	// 格式化 int，给定进制（进制 ∈ [2, 36]）
	formatInt := strconv.FormatInt(10, 2)
	fmt.Printf("%s", formatInt) // 打印：1010

	// 格式化 uint，给定进制（进制 ∈ [2, 36]）
	formatUint := strconv.FormatUint(10, 16)
	fmt.Printf("%s", formatUint) // 打印：a

	// 格式化 bool
	formatBool := strconv.FormatBool(true)
	fmt.Printf("%s", formatBool) // 打印：true

	// 格式化 float
	// 四个参数分别是：数值、格式、精度、数值位长
	// 【格式】有很多种，例如 'f' 表示 float 格式，'e' 表示指数格式
	// 【精度】与格式有关，表示小数点后数字个数，或总数字个数
	// 【数值位长】有两种：32(float32)、64(float64)
	formatFloat := strconv.FormatFloat(1.2, 'e', 2, 64)
	fmt.Printf("%s", formatFloat) // 打印：1.20e+00

	// 格式化复数，规则与 FormatFloat 类似
	formatComplex := strconv.FormatComplex(1+2i, 'f', 2, 64)
	fmt.Printf("%s", formatComplex) // 打印：(1.00+2.00i)

	// 解析 int 字符串
	// 三个参数分别是：数值字符串、进制、数值位长（0-64, 0 默认 strconv.IntSize）
	parseInt, _ := strconv.ParseInt("10", 10, 32)
	fmt.Printf("%d", parseInt) // 打印：10

	// 解析 uint 字符串，参数同 ParseInt
	parseUint, _ := strconv.ParseUint("10", 10, 32)
	fmt.Printf("%d", parseUint) // 打印：10

	// 解析 bool 字符串
	parseBool, _ := strconv.ParseBool("true")
	fmt.Printf("%t", parseBool) // 打印：true

	// 解析 float 字符串，第二个参数是数值位长
	parseFloat, _ := strconv.ParseFloat("1.20", 32)
	fmt.Printf("%.1f", parseFloat) // 打印：1.2

	// 解析复数字符串，第二个参数是数值位长
	parseComplex, _ := strconv.ParseComplex("1+2i", 32)
	fmt.Printf("%.1f", parseComplex) // 打印：(1.0+2.0i)

	// 解析 int 字符串，等同于 strconv.ParseInt(s, 10, 0)
	atoi, _ := strconv.Atoi("10")
	fmt.Printf("%d", atoi) // 打印：10

	// 转 int 为字符串，等同于 strconv.FormatInt(int64(i), 10)
	itoa := strconv.Itoa(10)
	fmt.Printf("%s", itoa) // 打印：10

	// 字节数组追加 int 字节
	// 三个参数分别是：原始字节数组、追加数值、追加数值进制
	appendInt := strconv.AppendInt([]byte("123"), 10, 2)
	fmt.Printf("%s", appendInt) // 打印：1231010

	// 字节数组追加 uint 字节，参数同 AppendInt
	appendUint := strconv.AppendUint([]byte("123"), 10, 2)
	fmt.Printf("%s", appendUint) // 打印：1231010

	// 字节数组追加 bool 字节
	appendBool := strconv.AppendBool([]byte("123"), true)
	fmt.Printf("%s", appendBool) // 打印：123true

	// 字节数组追加 float 字节
	appendFloat := strconv.AppendFloat([]byte("123"), 5.6, 'e', 2, 32)
	fmt.Printf("%s", appendFloat) // 打印：1235.60e+00

	// 字符串引用
	quote := strconv.Quote("123\t456")
	fmt.Printf("%s", quote) // 打印："123\t456"

	// 字符引用
	quoteRune := strconv.QuoteRune('1')
	fmt.Printf("%s", quoteRune) // 打印：'1'

	// 字符串转成 ASCII 引用
	QuoteToASCII := strconv.QuoteToASCII("我")
	fmt.Printf("%s", QuoteToASCII) // 打印："\u6211"

	// 字符转成 ASCII 引用
	quoteRuneToASCII := strconv.QuoteRuneToASCII('我')
	fmt.Printf("%s", quoteRuneToASCII) // 打印：'\u6211'

	// 字符串转成 unicode 图形字符串引用（由 strconv.IsGraphic() 定义）
	quoteToGraphic := strconv.QuoteToGraphic("!\u00a0")
	fmt.Printf("%s", quoteToGraphic) // 打印："! "

	// 字符转成 unicode 图形字符引用（由 strconv.IsGraphic() 定义）
	quoteRuneToGraphic := strconv.QuoteRuneToGraphic('a')
	fmt.Printf("%s", quoteRuneToGraphic) // 打印："a"

	// 字节数组追加引用字符串引用
	appendQuote := strconv.AppendQuote([]byte("123"), "456\t")
	fmt.Printf("%s", appendQuote) // 打印：123"456\t"

	// 字节数组追加引用字符引用
	appendQuoteRune := strconv.AppendQuoteRune([]byte("123"), '4')
	fmt.Printf("%s", appendQuoteRune) // 打印：123'4'

	// 字节数组追加 ASCII 字符引用
	appendQuoteToASCII := strconv.AppendQuoteToASCII([]byte("123"), "我")
	fmt.Printf("%s", appendQuoteToASCII) // 打印：123"\u6211"

	// 字节数组追加 unicode 图形字符串引用
	appendQuoteToGraphic := strconv.AppendQuoteToGraphic([]byte("123"), "4")
	fmt.Printf("%s", appendQuoteToGraphic) // 打印：123"4"

	// 字节数组追加 unicode 图形字符引用
	appendQuoteRuneToGraphic := strconv.AppendQuoteRuneToGraphic([]byte("123"), '4')
	fmt.Printf("%s", appendQuoteRuneToGraphic) // 打印：123'4'

	// 解析引用字符串
	unquote, _ := strconv.Unquote(`"123\t456"`)
	fmt.Printf("%s", unquote) // 打印：123	456

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
}
