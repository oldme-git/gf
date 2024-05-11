// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv_test

import (
	"testing"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

type SubScanTest struct {
	Name  string
	Place string
}

var scanValueTests = []struct {
	value interface{}
}{
	{map[string]bool{"0": false, "1": true}},
	{map[string]int{"0": 0, "1": 1}},
	{map[string]int8{"0": 0, "1": 1}},
	{map[string]int16{"0": 0, "1": 1}},
	{map[string]int32{"0": 0, "1": 1}},
	{map[string]int64{"0": 0, "1": 1}},
	{map[string]uint{"0": 0, "1": 1}},
	{map[string]uint8{"0": 0, "1": 1}},
	{map[string]uint16{"0": 0, "1": 1}},
	{map[string]uint32{"0": 0, "1": 1}},
	{map[string]uint64{"0": 0, "1": 1}},
	{map[string]float32{"0": 0, "1": 1}},
	{map[string]float64{"0": 0, "1": 1}},
	{map[string]string{"0": "0", "1": "1"}},
	{map[string][]byte{"0": []byte("0"), "1": []byte("1")}},
	{map[string]complex64{"0": 0, "1": 1 + 2i}},
	{map[string]complex128{"0": 0, "1": 1 + 2i}},
	{map[string]interface{}{"0": 0, "1": "1"}},
	{map[string]*gvar.Var{"0": gvar.New(0), "1": gvar.New("1")}},
	// TODO it should be failed at `gconv.Scan(test.value, &scanExpect.mapStrStruct)`.
	//{map[string]*gtime.Time{"0": gtime.New("2021-01-01 01:01:01"),
	//	"1": gtime.New("2021-01-01 01:01:01")}},
	{map[string]SubScanTest{"0": {"Mars", "马沃斯谷"}, "1": {"Venus", "伊师塔地"}}},
	{map[string]*SubScanTest{"0": {"Mars", "希腊平原"}, "1": {"Venus", "阿佛洛狄忒陆"}}},

	{map[int]interface{}{0: 0, 1: 1}},
	{map[uint]interface{}{0: 0, 1: 1}},
	{map[float32]interface{}{0: 0.12, 1: 1.23}},
	{map[float64]interface{}{0: 0.12, 1: 1.23}},
}

var scanExpectTests = struct {
	mapStrStr       map[string]string
	mapStrAny       map[string]interface{}
	mapStrStruct    map[string]SubScanTest
	mapStrStructPtr map[string]*SubScanTest
	mapAnyAny       map[interface{}]interface{}
}{
	mapStrStr:       make(map[string]string),
	mapStrAny:       make(map[string]interface{}),
	mapStrStruct:    make(map[string]SubScanTest),
	mapStrStructPtr: make(map[string]*SubScanTest),
	mapAnyAny:       make(map[interface{}]interface{}),
}

func TestScan(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for _, test := range scanValueTests {
			var (
				err        error
				scanExpect = scanExpectTests
			)

			err = gconv.Scan(test.value, &scanExpect.mapStrStr)
			t.AssertNil(err)
			t.Assert(test.value, scanExpect.mapStrStr)

			err = gconv.Scan(test.value, &scanExpect.mapStrAny)
			t.AssertNil(err)
			t.Assert(test.value, scanExpect.mapStrAny)

			err = gconv.Scan(test.value, &scanExpect.mapStrStruct)
			if err == nil {
				t.Assert(test.value, scanExpect.mapStrStruct)
			}

			err = gconv.Scan(test.value, &scanExpect.mapStrStructPtr)
			if err == nil {
				t.Assert(test.value, scanExpect.mapStrStructPtr)
			}

			err = gconv.Scan(test.value, &scanExpect.mapAnyAny)
			t.AssertNil(err)
			t.Assert(test.value, scanExpect.mapAnyAny)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		for _, test := range scanValueTests {
			var (
				err        error
				scanExpect = scanExpectTests
				maps       = []interface{}{
					test.value,
					test.value,
				}
			)

			mss := []map[string]string{
				scanExpect.mapStrStr,
				scanExpect.mapStrStr,
			}
			err = gconv.Scan(maps, &mss)
			t.AssertNil(err)
			for k, _ := range maps {
				t.Assert(maps[k], mss[k])
			}

			msa := []map[string]interface{}{
				scanExpect.mapStrAny,
				scanExpect.mapStrAny,
			}
			err = gconv.Scan(maps, &msa)
			t.AssertNil(err)
			for k, _ := range maps {
				t.Assert(maps[k], msa[k])
			}

			msst := []map[string]SubScanTest{
				scanExpect.mapStrStruct,
				scanExpect.mapStrStruct,
			}
			err = gconv.Scan(maps, &msst)
			if err == nil {
				for k, _ := range maps {
					t.Assert(maps[k], msst[k])
				}
			}

			msstp := []map[string]*SubScanTest{
				scanExpect.mapStrStructPtr,
				scanExpect.mapStrStructPtr,
			}
			err = gconv.Scan(maps, &msstp)
			if err == nil {
				for k, _ := range maps {
					t.Assert(maps[k], msstp[k])
				}
			}

			maa := []map[interface{}]interface{}{
				scanExpect.mapAnyAny,
				scanExpect.mapAnyAny,
			}
			err = gconv.Scan(maps, &maa)
			t.AssertNil(err)
			for k, _ := range maps {
				t.Assert(maps[k], maa[k])
			}
		}
	})
}
