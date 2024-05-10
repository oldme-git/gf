// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

var mapTests = []struct {
	value  interface{}
	expect map[string]interface{}
}{
	{map[int]int{1: 1}, map[string]interface{}{"1": 1}},
	{map[float64]int{1.1: 1}, map[string]interface{}{"1.1": 1}},
	{map[string]int{"k1": 1}, map[string]interface{}{"k1": 1}},

	{map[string]int{"k1": 1}, map[string]interface{}{"k1": 1}},
	{map[string]float64{"k1": 1.1}, map[string]interface{}{"k1": 1.1}},
	{map[string]string{"k1": "v1"}, map[string]interface{}{"k1": "v1"}},

	{`{"earth": "亚马逊雨林"}`,
		map[string]interface{}{"earth": "亚马逊雨林"}},
	{[]byte(`{"earth": "撒哈拉沙漠"}`),
		map[string]interface{}{"earth": "撒哈拉沙漠"}},

	{struct {
		Earth string
	}{
		Earth: "大峡谷",
	}, map[string]interface{}{"Earth": "大峡谷"}},
	{&struct {
		Earth string
	}{
		Earth: "马里亚纳海沟",
	}, map[string]interface{}{"Earth": "马里亚纳海沟"}},

	{[]int{}, map[string]interface{}{}},
	{[]int{1, 2, 3}, map[string]interface{}{"1": "2", "3": nil}},
	{[]int{1, 2, 3, 4}, map[string]interface{}{"1": "2", "3": "4"}},
}

func TestMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for _, test := range mapTests {
			t.Assert(gconv.Map(test.value), test.expect)
		}
	})
}
