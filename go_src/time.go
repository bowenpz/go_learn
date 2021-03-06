package main

import (
	"fmt"
	"math/rand"
	"time" // 1.16.5
)

// —————————————————————————————————————————————
// —————————————— 翻译：time 包注释 ——————————————
// —————————————————————————————————————————————

// Package time provides functionality for measuring and displaying time.
// time 包提供了测量和显示时间的能力

// The calendrical calculations always assume a Gregorian calendar, with
// no leap seconds.
// 日期计算时，使用 Gregorian calendar（公历），并且没有闰秒
// Gregorian calendar：https://zh.wikipedia.org/wiki/%E6%A0%BC%E9%87%8C%E6%9B%86
// 闰秒：https://zh.wikipedia.org/wiki/%E9%97%B0%E7%A7%92

// Monotonic Clocks
// 时间单调增长的时钟（后文翻译将继续使用 monotonic clock 这个英文说法）

// Operating systems provide both a “wall clock,” which is subject to
// changes for clock synchronization, and a “monotonic clock,” which is
// not. The general rule is that the wall clock is for telling time and
// the monotonic clock is for measuring time. Rather than split the API,
// in this package the Time returned by time.Now contains both a wall
// clock reading and a monotonic clock reading; later time-telling
// operations use the wall clock reading, but later time-measuring
// operations, specifically comparisons and subtractions, use the
// monotonic clock reading.
// 操作系统提供了两种 clock：一种是 wall clock，它跟真实世界的时间同步；一种是 monotonic clock，它的时间纯单调递增。
// 通常来讲，wall clock 用于展示时间，monotonic clock 用于计算时间。
// 本 time 包并没有割裂成两套 API，而是通过 time.Now 这一同时包含 wall clock 和 monotonic clock 的方法返回时间。
// time.Now 内部使用 wall clock 来展示时间，使用 monotonic clock 来计算时间，尤其是用于比较时间和计算时间差值。

// For example, this code always computes a positive elapsed time of
// approximately 20 milliseconds, even if the wall clock is changed during
// the operation being timed:
// 举个例子，下面这段代码将计算出大概 20 微秒的时间，即使在计算过程中 wall clock 变化了，算出来的结果也不会变化。

//	start := time.Now()
//	... operation that takes 20 milliseconds ...
//	t := time.Now()
//	elapsed := t.Sub(start)
//
// Other idioms, such as time.Since(start), time.Until(deadline), and
// time.Now().Before(deadline), are similarly robust against wall clock
// resets.
// 还有些类似的 API，比如 time.Since(start)、time.Until(deadline)、time.Now().Before(deadline)，都不会随 wall clock 变化而变化

// The rest of this section gives the precise details of how operations
// use monotonic clocks, but understanding those details is not required
// to use this package.
// 接下来要讲的部分，会讲清楚 time 包使用 monotonic clocks 的细节，但是单纯使用 time 包的话，可以不必知道这些细节。

// The Time returned by time.Now contains a monotonic clock reading.
// If Time t has a monotonic clock reading, t.Add adds the same duration to
// both the wall clock and monotonic clock readings to compute the result.
// Because t.AddDate(y, m, d), t.Round(d), and t.Truncate(d) are wall time
// computations, they always strip any monotonic clock reading from their results.
// Because t.In, t.Local, and t.UTC are used for their effect on the interpretation
// of the wall time, they also strip any monotonic clock reading from their results.
// The canonical way to strip a monotonic clock reading is to use t = t.Round(0).
// 使用 time.Now 方法返回的时间对象，内部包含了 monotonic clock 时间。
// 如果一个时间对象 t 内部包含 monotonic clock，那么使用 t.add 方法增加时间时，会同时往 t 的 wall clock 和 monotonic clock 增加这段时间。
// 因为 t.AddDate(y, m, d)、t.Round(d)、t.Truncate(d) 都是在计算 wall time，它们会在结果中删掉 monotonic clock。
// 因为 t.In、t.Local、t.UTC 是用于解释 wall time 的，它们也会在结果中删掉 monotonic clock。
// 删掉 monotonic clock 的规范写法是 t = t.Round(0)。

// If Times t and u both contain monotonic clock readings, the operations
// t.After(u), t.Before(u), t.Equal(u), and t.Sub(u) are carried out
// using the monotonic clock readings alone, ignoring the wall clock
// readings. If either t or u contains no monotonic clock reading, these
// operations fall back to using the wall clock readings.
// 如果时间对象 t 和 u 都包含 monotonic clock 时间，那么在使用 t.After(u)、t.Before(u)、t.Equal(u)、t.Sub(u) 时，都只会使用 monotonic clock 计算。
// 如果时间对象 t 和 u 只要有一个没有 monotonic clock 时间，那么将使用 wall clock 计算。

// On some systems the monotonic clock will stop if the computer goes to sleep.
// On such a system, t.Sub(u) may not accurately reflect the actual
// time that passed between t and u.
// 有些系统在睡眠时会停止 monotonic clock。在这些系统上，t.Sub(u) 算出的结果不准。

// Because the monotonic clock reading has no meaning outside
// the current process, the serialized forms generated by t.GobEncode,
// t.MarshalBinary, t.MarshalJSON, and t.MarshalText omit the monotonic
// clock reading, and t.Format provides no format for it. Similarly, the
// constructors time.Date, time.Parse, time.ParseInLocation, and time.Unix,
// as well as the unmarshalers t.GobDecode, t.UnmarshalBinary.
// t.UnmarshalJSON, and t.UnmarshalText always create times with
// no monotonic clock reading.
// 因为 monotonic clock 时间在当前进程外没有任何意义，因此 t.GobEncode、t.MarshalBinary、t.MarshalJSON、t.MarshalText 这些序列化方法，将会省略 monotonic clock。
// 同样地，time.Date、time.Parse、time.ParseInLocation、time.Unix 这些构造方法，以及 t.GobDecode、t.UnmarshalBinary、t.UnmarshalJSON、t.UnmarshalText
// 这些反序列化方法，创建出的时间对象也不会包含 monotonic clock 时间。

// Note that the Go == operator compares not just the time instant but
// also the Location and the monotonic clock reading. See the documentation for the Time type for a discussion of equality testing for Time values.
// documentation for the Time type for a discussion of equality
// testing for Time values.
// 需要注意的是，在 Go 中 == 操作符不光会比较时间戳，也会比较时区和 monotonic clock 时间。
// 跟时间相等相关的部分，参阅 time.Time 类型的介绍。

// For debugging, the result of t.String does include the monotonic
// clock reading if present. If t != u because of different monotonic clock readings,
// that difference will be visible when printing t.String() and u.String().
// 在调试时，如果时间对象 t 内部存在 monotonic clock 时间，t.String 方法会返回该时间。
// 如果两个时间对象 t 和 u 因为 monotonic clock 不同而不同，可以通过 t.String() 和 u.String() 方法显示出来。

// ——————————————————————————————————————————————————
// —————————————— 翻译：time.Time 类注释 ——————————————
// ——————————————————————————————————————————————————

// A Time represents an instant in time with nanosecond precision.
// time.Time 类代表着一个精度到纳秒的时间戳（time instant）。

// Programs using times should typically store and pass them as values,
// not pointers. That is, time variables and struct fields should be of
// type time.Time, not *time.Time.
// 程序在存储和传递时间对象时，应该使用值类型，而非指针类型。
// 因此，在使用 time 变量和结构体时，应当使用 time.Time，而非 *time.Time。

// A Time value can be used by multiple goroutines simultaneously except
// that the methods GobDecode, UnmarshalBinary, UnmarshalJSON and
// UnmarshalText are not concurrency-safe.
// time.Time 对象在多 goroutine 环境下，除了 GobDecode、UnmarshalBinary、UnmarshalJSON 和 UnmarshalText 这四个方法下不是并发安全之外，其余都并发安全。

// Time instants can be compared using the Before, After, and Equal methods.
// The Sub method subtracts two instants, producing a Duration.
// The Add method adds a Time and a Duration, producing a Time.
// 时间戳可以使用 Before、After、Equal 方法进行比较。
// Sub 方法对两个 time.Time 相减，得到一个 time.Duration。
// Add 方法把 time.Time 和 time.Duration 相加，得到一个 time.Time。

// The zero value of type Time is January 1, year 1, 00:00:00.000000000 UTC.
// As this time is unlikely to come up in practice, the IsZero method gives
// a simple way of detecting a time that has not been initialized explicitly.
// time.Time 的零值是公元 1 年 1 月 1 日，00:00:00.000000000 UTC。
// 由于这个时间在实际场景中不太可能出现，因此提供 IsZero 方法检测 time.Time 对象是否被显式初始化。

// Each Time has associated with it a Location, consulted when computing the
// presentation form of the time, such as in the Format, Hour, and Year methods.
// The methods Local, UTC, and In return a Time with a specific location.
// Changing the location in this way changes only the presentation; it does not
// change the instant in time being denoted and therefore does not affect the
// computations described in earlier paragraphs.
// 每一个 time.Time 对象会关联时区（Location）信息，用于表示时间格式，例如在 Format、Hour、Year 方法中会使用。
// Local、UTC、In 方法会返回一个包含指定时区的 time.Time 对象。
// 以这些方法修改时区，只会修改时间显示，不会修改时间戳，也不会影响之前已有的时间。

// Representations of a Time value saved by the GobEncode, MarshalBinary,
// MarshalJSON, and MarshalText methods store the Time.Location's offset, but not
// the location name. They therefore lose information about Daylight Saving Time.
// 使用 GobEncode、MarshalBinary、MarshalJSON 和 MarshalText 方法存储时间，只会包含时区偏移量（Time.Location's offset），不会存储时区名称。
// 因此这些方法会丢失夏令时信息。

// In addition to the required “wall clock” reading, a Time may contain an optional
// reading of the current process's monotonic clock, to provide additional precision
// for comparison or subtraction.
// See the “Monotonic Clocks” section in the package documentation for details.
// time.Time 除了包含必需的 wall clock 之外，还可能会包含当前程序的 monotonic clock，用于精确地比较时间、计算时间差值。
// 有关 monotonic clock 的部分可以阅读 time 包的文档。

// Note that the Go == operator compares not just the time instant but also the
// Location and the monotonic clock reading. Therefore, Time values should not
// be used as map or database keys without first guaranteeing that the identical Location has been set for all values
// identical Location has been set for all values, which can be achieved
// through use of the UTC or Local method, and that the monotonic clock reading
// has been stripped by setting t = t.Round(0). In general, prefer t.Equal(u)
// to t == u, since t.Equal uses the most accurate comparison available and
// correctly handles the case when only one of its arguments has a monotonic
// clock reading.
// 需要注意的是，在 Go 中 == 操作符不光会比较时间戳，也会比较时区和 monotonic clock 时间。
// 因此，如果没有使用 UTC 或 Local 方法确保时区相同，或者没有使用 t.Round(0) 去除掉 monotonic clock 时间，不应该使用 time.Time 作为 map 或数据库的 key。
// 通常情况下，优先使用 t.Equal(u)，而不是 t == u，因为 t.Equal 方法更精确，并且能正确处理单 monotonic clock 的 case。

func learnTime() {
	// —————————————— 获取 time.Time 对象 ——————————————

	// 当前时间
	now := time.Now()
	fmt.Printf("%v", now) // 打印：2021-12-10 16:44:10.337708 +0800 CST m=+0.000118756

	// 根据 unix 时间戳获取时间
	unix := time.Unix(1639094400, 0)
	fmt.Printf("%v", unix) // 打印：2021-12-10 08:00:00 +0800 CST

	// 根据日期获取时间
	date := time.Date(2021, time.December, 10, 0, 0, 0, 0, time.UTC)
	fmt.Printf("%v", date) // 打印：2021-12-10 08:00:00 +0800 CST

	// 根据字符串文本获取时间（默认本地时区）
	parse, _ := time.Parse("2006-01-02 15:04:05", "2021-12-10 00:00:00")
	fmt.Printf("%v", parse) // 打印：2021-12-10 08:00:00 +0800 CST

	// 根据字符串文本获取时间
	parseInLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-12-10 00:00:00", time.Local)
	fmt.Printf("%v", parseInLocation) // 打印：2021-12-10 00:00:00 +0800 CST

	// —————————————— 使用 time.Time 对象：展示 ——————————————

	// 接下来都使用 2021-01-01 00:00:00 这个时间
	t := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)

	// 日期与时刻
	var (
		year                          = t.Year()
		month                         = t.Month()
		weekday                       = t.Weekday()
		yearDay                       = t.YearDay()
		day                           = t.Day()
		hour                          = t.Hour()
		minute                        = t.Minute()
		second                        = t.Second()
		nanosecond                    = t.Nanosecond()
		isoYear, isoWeek              = t.ISOWeek()
		dateYear, dateMonth, dateDay  = t.Date()
		clockHour, clockMin, clockSec = t.Clock()
	)
	fmt.Printf("%d %d %d %d %d %d %d %d %d\n", year, month, weekday, yearDay, day, hour, minute, second, nanosecond)     // 打印：2021 1 5 1 1 0 0 0 0
	fmt.Printf("%d %d %d %d %d %d %d %d", isoYear, isoWeek, dateYear, dateMonth, dateDay, clockHour, clockMin, clockSec) // 打印：2020 53 2021 1 1 0 0 0

	// unix 时间（单位：秒）
	tUnix := t.Unix()
	fmt.Printf("%d ", tUnix) // 打印：1609459200

	// unix 时间（单位：纳秒）
	tUnixNano := t.UnixNano()
	fmt.Printf("%d", tUnixNano) // 打印：1609459200000000000

	// 所在时区
	location := t.Location()
	fmt.Printf("%v", location) // 打印：UTC

	// 所在时区与偏移量（秒）
	zone, offset := t.Zone()
	fmt.Printf("%s %d", zone, offset) // 打印：UTC 0

	// 设定时区
	var (
		inUTC       = t.UTC()                                  // 时区设置为 UTC
		inLocal     = t.Local()                                // 时区设置为本地
		inFixedZone = t.In(time.FixedZone("Hangzhou", 8*3600)) // 时区设置为指定时区
	)
	fmt.Printf("%v\n%v\n%v", inUTC, inLocal, inFixedZone)
	// 打印：
	// 2021-01-01 00:00:00 +0000 UTC
	// 2021-01-01 08:00:00 +0800 CST
	// 2021-01-01 08:00:00 +0800 Hangzhou

	// 格式化时间
	// go 格式化时间非常有病，并不使用 YYYY 之类的占位符，而是使用 go 出生的那一刻作为时间格式：Mon Jan 2 15:04:05 -0700 MST 2006
	format := t.Format("2006-01-02 15:04:05 MST")
	fmt.Printf("%s", format) // 打印：2021-01-01 00:00:00 UTC

	// （带前缀）格式化时间
	appendFormat := t.AppendFormat([]byte("Time: "), "15:04:05 PM")
	fmt.Printf("%s", appendFormat) // 打印：Time: 00:00:00 AM

	// String 方法
	// 等同于 t.Format("2006-01-02 15:04:05.999999999 -0700 MST")
	// 如果有 monotonic clock（例如 time.Now()）按照 m=±ddd.nnnnnnnnn 格式打印
	tStr := t.String()
	fmt.Printf("%s", tStr) // 打印：2021-01-01 00:00:00 +0000 UTC

	// 序列化
	var (
		encode, _        = t.GobEncode()
		marshalBinary, _ = t.MarshalBinary()
		marshalJSON, _   = t.MarshalJSON()
		marshalText, _   = t.MarshalText()
	)
	fmt.Printf("%v\n%v\n%s\n%s", encode, marshalBinary, marshalJSON, marshalText)
	// 打印：
	// [1 0 0 0 14 215 128 93 0 0 0 0 0 255 255]
	// [1 0 0 0 14 215 128 93 0 0 0 0 0 255 255]
	// "2021-01-01T00:00:00Z"
	// 2021-01-01T00:00:00Z2021-01-01 00:00:00 +0000 UTC

	// 反序列化
	decode, m1, m2, m3 := time.Now(), time.Now(), time.Now(), time.Now()
	decode.GobDecode(encode)
	m1.UnmarshalBinary(marshalBinary)
	m2.UnmarshalJSON(marshalJSON)
	m3.UnmarshalText(marshalText)
	fmt.Printf("%v\n%v\n%v\n%v", decode, m1, m2, m3)
	// 打印：
	// 2021-01-01T00:00:00Z2021-01-01 00:00:00 +0000 UTC
	// 2021-01-01 00:00:00 +0000 UTC
	// 2021-01-01 00:00:00 +0000 UTC
	// 2021-01-01 00:00:00 +0000 UTC

	// —————————————— 使用 time.Time 对象：比较与计算 ——————————————

	// 增加时间
	add := t.Add(time.Hour * 24)
	fmt.Printf("%v", add) // 打印：2021-01-02 00:00:00 +0000 UTC

	// 增加日期
	addDate := t.AddDate(1, 1, 1)
	fmt.Printf("%v", addDate) // 打印：2022-02-02 00:00:00 +0000 UTC

	// 相减得到时间间隔
	sub := t.Sub(add)
	fmt.Printf("%v", sub) // 打印：-24h0m0s

	// 是否在指定时间之前
	isBefore := t.Before(time.Now())
	fmt.Printf("%t", isBefore) // 打印：true

	// 是否在指定时间之后
	isAfter := t.After(time.Now())
	fmt.Printf("%t", isAfter) // 打印：false

	// 时间是否相等（即使时区不同也可以相等，因此不建议使用 ==，优先使用 Equal）
	t1 := t.In(time.FixedZone("zone1", 0))
	t2 := t.In(time.FixedZone("zone2", 8*3600))
	fmt.Printf("%t", t1.Equal(t2)) // 打印：true
	fmt.Printf("%t", t1 == t2)     // 打印：false

	// 是否是时间零值（公元 1 年第 0 秒是零值）
	isZero1 := t.IsZero()
	isZero2 := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC).IsZero()
	fmt.Printf("%t", isZero1) // 打印：false
	fmt.Printf("%t", isZero2) // 打印：true

	// 时间取整
	var (
		randTime = time.Date(2021, time.January, 1, 0, 30+rand.Intn(30), rand.Intn(60), rand.Intn(60), time.UTC)
		round    = randTime.Round(time.Hour)    // 四舍五入取整
		truncate = randTime.Truncate(time.Hour) // 向下取整
	)
	fmt.Printf("%v\n%v", round, truncate)
	// 打印：
	// 2021-01-01 01:00:00 +0000 UTC
	// 2021-01-01 00:00:00 +0000 UTC

	// —————————————— 获取 *time.Location 指针 ——————————————

	// 获取 UTC 时区
	utc := time.UTC
	fmt.Printf("%v", utc) // 打印：UTC

	// 获取当前时区
	local := time.Local
	fmt.Printf("%v", local) // 打印：Local

	// 获取指定时区
	fixedZone := time.FixedZone("Hangzhou", 8*3600)
	fmt.Printf("%v", fixedZone) // 打印：Hangzhou

	// 加载时区
	loadLocation, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Printf("%v", loadLocation) // 打印：Asia/Shanghai

	// 根据 TZData 加载时区（感觉基本用不上）
	loadLocationFromTZData, err := time.LoadLocationFromTZData("Asia/Beijing", []byte{})
	fmt.Printf("%v, %s", loadLocationFromTZData, err) // 打印：UTC, malformed time zone information

	// —————————————— 时间区间 ——————————————

	// 指定时间距今有多长时间（等同于 t.Sub(now)）
	until := time.Until(time.Date(2021, time.December, 10, 0, 0, 0, 0, time.UTC))
	fmt.Printf("%v", until) // 打印：-105h50m3.370672s

	// 至今过去了多长时间（等同于 now.Sub(t)）
	since := time.Since(time.Date(2021, time.December, 10, 0, 0, 0, 0, time.UTC))
	fmt.Printf("%v", since) // 打印：105h50m3.37071s

	// 根据字符串解析成时间区间
	parseDuration, _ := time.ParseDuration("1h0m0s")
	fmt.Printf("%v", parseDuration) // 打印：1h0m0s

	// 除此之外还有一些枚举：
	// time.Hour
	// time.Minute
	// time.Second
	// time.Millisecond
	// time.Microsecond
	// time.Nanosecond

	// —————————————— 直接休眠 ——————————————

	// 让 Goroutine 休眠不少于多少时间
	time.Sleep(time.Second)

	// —————————————— 计时器 time.Timer ——————————————

	// 休眠一段时间（只暂停一次）
	timer := time.NewTimer(time.Second)
	select {
	case <-timer.C:
		// ...
	}
	// 重新计时
	timer.Reset(time.Second)
	// 终止
	timer.Stop()

	// 计时器 time.Timer 的简易版本，但是内部的 Time 不会被 GC，会造成内存泄漏
	after := time.After(time.Second)
	select {
	case <-after:
		// ...
	}

	// 延迟执行方法
	time.AfterFunc(time.Second, func() {
		// ...
	})

	// —————————————— 打点器 time.Ticker ——————————————

	// 多次休眠一段时间（无限次暂停）
	ticker := time.NewTicker(time.Second)
	for i := 0; i < 10; i++ {
		select {
		case <-ticker.C:
			// ...
		}
	}
	// 重新计时
	ticker.Reset(time.Second)
	// 终止
	ticker.Stop()

	// 打点器 time.Ticker 的简易版本，适合无限循环（如果不是无限循环，内部的 Ticker 会无法 GC 造成内存泄漏）
	tick := time.Tick(time.Second)
	for i := 0; i < 10; i++ {
		select {
		case <-tick:
			// ...
		}
	}

	// —————————————— 一些枚举 ——————————————

	// 月份
	months := []time.Month{
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December,
	}
	fmt.Printf("%v", months) // 打印：[January February March April May June July August September October November December]

	// 周几
	weekdays := []time.Weekday{
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
		time.Sunday,
	}
	fmt.Printf("%v", weekdays) // 打印：[Monday Tuesday Wednesday Thursday Friday Saturday Sunday]

	// 持续时间（最大只到小时，避免夏令时造成的时间错乱）
	durations := []time.Duration{
		time.Hour,
		time.Minute,
		time.Second,
		time.Millisecond,
		time.Microsecond,
		time.Nanosecond,
	}
	fmt.Printf("%v", durations) // 打印：[1h0m0s 1m0s 1s 1ms 1µs 1ns]

	// 还有一些时间格式的字符串枚举，例如 time.RFC822，就不粘贴了
}

func separate() {
	fmt.Printf("\n\n\n")
}
func separateInterface(i interface{}) {
	fmt.Printf("\n\n\n")
	fmt.Printf("%v", i)
	fmt.Printf("\n\n\n")
}
