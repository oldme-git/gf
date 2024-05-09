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

	{struct {
		Name  string
		Place string
	}{
		Name:  "Earth",
		Place: "马里亚纳海沟",
	}, map[string]interface{}{"Name": "Earth", "Place": "马里亚纳海沟"}},
	{&struct {
		Name  string
		Place string
	}{
		Name:  "Earth",
		Place: "马里亚纳海沟",
	}, map[string]interface{}{"Name": "Earth", "Place": "马里亚纳海沟"}},
}

func TestMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for _, test := range mapTests {
			t.Assert(gconv.Map(test.value), test.expect)
		}
	})
}
