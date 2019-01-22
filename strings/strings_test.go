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
s
	用于查找的字符串
substr
	要查找统计的子字符串

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
s
	用于查找的字符串
substr
	子字符串

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
			result := strings.Contains("我是最帅的aaa", "帅")
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
s
	用于查找的字符串
chars
	子字符串

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
			result := strings.Contains("aaaa", "ab")
			fmt.Println(result)
			So(result, ShouldBeTrue)
		})
	})
}
