package main

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

/**
统计字符串中有多少个子字符串

目标参数说明
s string
	用于查找的字符串
substr string
	要查找统计的子字符串

返回值
int
	子字符串个数

测试命令
	go test -v -test.run TestCount
*/
func TestCount(t *testing.T) {
	Convey("统计字符串中有多少个子字符串:", t, func() {
		Convey("统计字符串中有几个b", func() {
			result := strings.Count("abcnnnabc", "b")
			fmt.Println(result)
			So(result, ShouldEqual, 2)
		})
		Convey("统计字符串中有几个abc:", func() {
			result := strings.Count("abcnnnabc", "abc")
			fmt.Println(result)
			So(result, ShouldEqual, 2)
		})
		Convey("统计字符串中有几个空字符串(这里直接返回字符串长度):", func() {
			result := strings.Count("abcnnnabc", "")
			fmt.Println(result)
			So(result, ShouldEqual, 10)
		})
	})
}

/**
字符串s中是否包含子字符串substr

目标参数说明
s string
	用于查找的字符串
substr string
	子字符串

返回值
bool
	是否包含

测试命令
go test -v -test.run TestContains
*/
func TestContains(t *testing.T) {
	Convey("是否包含子字符串", t, func() {
		Convey("包含子字符串场景:", func() {
			result := strings.Contains("abc", "ab")
			fmt.Println(result)
			So(result, ShouldBeTrue)
		})
		Convey("不包含子字符串场景:", func() {
			result := strings.Contains("abc", "dd")
			fmt.Println(result)
			So(result, ShouldBeFalse)
		})
		Convey("空字符串是否包含:", func() {
			result := strings.Contains("", "dd")
			fmt.Println(result)
			So(result, ShouldBeFalse)
		})
		Convey("空字符串是否包含空字符串:", func() {
			result := strings.Contains("", "")
			fmt.Println(result)
			So(result, ShouldBeTrue)
		})
		Convey("空字符串是否被包含:", func() {
			result := strings.Contains("abc", "")
			fmt.Println(result)
			So(result, ShouldBeTrue)
		})
		Convey("中文测试是否包含:", func() {
			result := strings.Contains("你是猴子请来的救兵吗", "猴子")
			fmt.Println(result)
			So(result, ShouldBeTrue)
		})
		Convey("unicode字符测试1:", func() {
			result := strings.Contains("\u54c8\u54c8\u54c8", "哈")
			fmt.Println(result)
			So(result, ShouldBeTrue)
		})
		Convey("unicode字符测试2:", func() {
			result := strings.Contains("哈哈哈", "\u54c8\u54c8\u54c8")
			fmt.Println(result)
			So(result, ShouldBeTrue)
		})
		Convey("unicode字符测试3:", func() {
			result := strings.Contains("\u54c8\u54c8\u54c8", "\u54c8")
			fmt.Println(result)
			So(result, ShouldBeTrue)
		})
	})
}

/**
字符串s中是否包含子字符串chars中的任意字符

目标参数说明
s string
	用于查找的字符串
chars string
	子字符串

返回值
bool
	是否包含

测试命令
	go test -v -test.run TestContainsany
*/
func TestContainsany(t *testing.T) {
	Convey("是否包含子字符串中的任意字符", t, func() {
		Convey("包含子字符串场景:", func() {
			result := strings.ContainsAny("abc", "cdef")
			fmt.Println(result)
			So(result, ShouldBeTrue)
		})
		Convey("中文测试是否包含:", func() {
			result := strings.ContainsAny("你是猴子请来的救兵吗", "猴哥")
			fmt.Println(result)
			So(result, ShouldBeTrue)
		})
		Convey("空字符串是否包含:", func() {
			result := strings.ContainsAny("", "dd")
			fmt.Println(result)
			So(result, ShouldBeFalse)
		})
		Convey("空字符串是否包含空字符串:", func() {
			result := strings.ContainsAny("", "")
			fmt.Println(result)
			So(result, ShouldBeFalse)
		})
	})
}

/**
字符串s中是否包含子字符串chars中的任意字符

目标参数说明
s string
	用于查找的字符串
r rune
	目标字符

返回值
bool
	是否包含

测试命令
go test -v -test.run TestContainsRune
*/
func TestContainsRune(t *testing.T) {
	Convey("查找是否包含字符", t, func() {
		Convey("包含字符场景", func() {
			result := strings.ContainsRune("abc", 'a')
			fmt.Println(result)
			So(result, ShouldBeTrue)
		})
		Convey("空字符串测试1", func() {
			result := strings.ContainsRune("", 'a')
			fmt.Println(result)
			So(result, ShouldBeFalse)
		})
		Convey("中文测试", func() {
			result := strings.ContainsRune("你是猴子吗", '猴')
			fmt.Println(result)
			So(result, ShouldBeTrue)
		})
	})
}

/**
获取子字符串出现的最后一次位置

目标参数说明
s string
	用于查找的字符串
substr
	目标子字符串

返回值
int
	索引位置，如果substr为空，直接返回s长度，否则返回自字符串所在位置
	如果不存在返回-1

测试命令
go test -v -test.run TestLastIndex
*/
func TestLastIndex(t *testing.T) {
	Convey("查找子字符串最后一次出现的位置", t, func() {
		Convey("查找子字符串最后一次位置:", func() {
			result := strings.LastIndex("abcabc", "abc")
			fmt.Println(result)
			So(result, ShouldEqual, 3)
		})
		Convey("不存在情况测试:", func() {
			result := strings.LastIndex("abcabc", "ddd")
			fmt.Println(result)
			So(result, ShouldEqual, -1)
		})
		Convey("空字符串测试1:", func() {
			result := strings.LastIndex("", "abc")
			fmt.Println(result)
			So(result, ShouldEqual, -1)
		})
		Convey("空字符串测试2:", func() {
			result := strings.LastIndex("", "")
			fmt.Println(result)
			So(result, ShouldEqual, 0)
		})
		Convey("空字符串测试3:", func() {
			s := "abc"
			result := strings.LastIndex(s, "")
			fmt.Println(result)
			So(result, ShouldEqual, len(s))
		})
		Convey("中文测试:", func() {
			s := "我是猴子请来的猴子兵"
			result := strings.LastIndex(s, "猴子")
			fmt.Println(result)
			// 因为中文在utf-8编码下占3个字节
			So(result, ShouldEqual, 21)
		})

	})
}

/**
查找字符在字符串中第一次出现的位置

目标参数说明
s string
	用于查找的目标字符串
r rune
	字符用于查找的字符串

返回值
int
	字符第一次出现的索引位置

测试命令
go test -v -test.run TestIndexRune
*/
func TestIndexRune(t *testing.T) {
	Convey("查找字符r在字符串s中第一次出现的位置", t, func() {
		Convey("查找第一次出现的位置:", func() {
			result := strings.IndexRune("abc", 'b')
			fmt.Println(result)
			So(result, ShouldEqual, 1)
		})
		Convey("找不到的情况:", func() {
			result := strings.IndexRune("abc", 'f')
			fmt.Println(result)
			So(result, ShouldEqual, -1)
		})
		Convey("空字符串的情况:", func() {
			result := strings.IndexRune("", 'f')
			fmt.Println(result)
			So(result, ShouldEqual, -1)
		})
		Convey("中文测试:", func() {
			result := strings.IndexRune("又是猴子", '猴')
			fmt.Println(result)
			So(result, ShouldEqual, 6)
		})
	})
}

/**
查找子字符串中任意一个字符第一次出现的位置

目标参数说明
s string
	用于查找的目标字符串
chars string
	用于查找的字符串

返回值
int
	chars所在索引位置

测试命令
go test -v -test.run TestIndexAny
*/
func TestIndexAny(t *testing.T) {
	Convey("查找子字符串中任意一个字符出现的位置", t, func() {
		Convey("测试存在的情况:", func() {
			result := strings.IndexAny("abc", "fbf")
			fmt.Println(result)
			So(result, ShouldEqual, 1)
		})
		Convey("测试不存在的情况:", func() {
			result := strings.IndexAny("abc", "dd")
			fmt.Println(result)
			So(result, ShouldEqual, -1)
		})
		Convey("测试中文:", func() {
			result := strings.IndexAny("测试中文", "语文")
			fmt.Println(result)
			So(result, ShouldEqual, 9)
		})
		// 如果有多个只返回第一个找到的位置
		Convey("测试中文有多个的情况:", func() {
			result := strings.IndexAny("测试中文，我是文化人", "语文")
			fmt.Println(result)
			So(result, ShouldEqual, 9)
		})
	})
}

/**


目标参数说明
s string
	用于查找的目标字符串
chars string
	用于查找的子字符串

返回值
int
	找到的最后一个字符串的位置


测试命令
go test -v -test.run
*/
func TestLastIndexAny(t *testing.T) {
}
