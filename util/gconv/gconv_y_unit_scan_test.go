// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv_test

import (
	"testing"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

var scanValueMapsTest = []map[string]interface{}{
	{"Name": false, "Place": true},
	{"Name": int(0), "Place": int(1)},
	{"Name": int8(0), "Place": int8(1)},
	{"Name": int16(0), "Place": int16(1)},
	{"Name": int32(0), "Place": int32(1)},
	{"Name": int64(0), "Place": int64(1)},
	{"Name": uint(0), "Place": uint(1)},
	{"Name": uint8(0), "Place": uint8(1)},
	{"Name": uint16(0), "Place": uint16(1)},
	{"Name": uint32(0), "Place": uint32(1)},
	{"Name": uint64(0), "Place": uint64(1)},
	{"Name": float32(0), "Place": float32(1)},
	{"Name": float64(0), "Place": float64(1)},
	{"Name": "Mercury", "Place": "卡罗利斯盆地"},
	{"Name": []byte("Saturn"), "Place": []byte("土星环")},
	{"Name": complex64(0), "Place": complex64(1 + 2i)},
	{"Name": complex128(0), "Place": complex128(1 + 2i)},
	{"Name": interface{}(0), "Place": interface{}("1")},
	{"Name": gvar.New("Jupiter"), "Place": gvar.New("大红斑")},
	{"Name": gtime.New("2024-01-01 01:01:01"), "Place": gtime.New("2021-01-01 01:01:01")},
}

type scanStructTest struct {
	Name  string
	Place string
}

var scanValueStructsTest = []scanStructTest{
	{"Venus", "阿佛洛狄特高原"},
}

var scanValueJsonTest = []string{
	`{"Name": "Mars", "Place": "奥林帕斯山"}`,
}

type scanExpectTest struct {
	mapStrStr map[string]string
	mapStrAny map[string]interface{}
	mapAnyAny map[interface{}]interface{}

	structSub    scanStructTest
	structSubPtr *scanStructTest
}

var scanExpects = scanExpectTest{
	mapStrStr: make(map[string]string),
	mapStrAny: make(map[string]interface{}),
	mapAnyAny: make(map[interface{}]interface{}),

	structSub:    scanStructTest{},
	structSubPtr: &scanStructTest{},
}

func TestScan(t *testing.T) {
	// Test for map converting.
	gtest.C(t, func(t *gtest.T) {
		scanValuesTest := scanValueMapsTest
		for _, test := range scanValuesTest {
			var (
				err         error
				scanExpects = scanExpects
			)

			err = gconv.Scan(test, &scanExpects.mapStrStr)
			t.AssertNil(err)
			t.Assert(test["Name"], scanExpects.mapStrStr["Name"])
			t.Assert(test["Place"], scanExpects.mapStrStr["Place"])

			err = gconv.Scan(test, &scanExpects.mapStrAny)
			t.AssertNil(err)
			t.Assert(test["Name"], scanExpects.mapStrAny["Name"])
			t.Assert(test["Place"], scanExpects.mapStrAny["Place"])

			err = gconv.Scan(test, &scanExpects.mapAnyAny)
			t.AssertNil(err)
			t.Assert(test["Name"], scanExpects.mapAnyAny["Name"])
			t.Assert(test["Place"], scanExpects.mapAnyAny["Place"])

			err = gconv.Scan(test, &scanExpects.structSub)
			t.AssertNil(err)
			t.Assert(test["Name"], scanExpects.structSub.Name)
			t.Assert(test["Place"], scanExpects.structSub.Place)

			err = gconv.Scan(test, &scanExpects.structSubPtr)
			t.AssertNil(err)
			t.Assert(test["Name"], scanExpects.structSubPtr.Name)
			t.Assert(test["Place"], scanExpects.structSubPtr.Place)

		}
	})

	// Test for slice map converting.
	gtest.C(t, func(t *gtest.T) {
		scanValuesTest := scanValueMapsTest
		for _, test := range scanValuesTest {
			var (
				err         error
				scanExpects = scanExpects
				mapList     = []map[string]interface{}{test, test}
			)

			var mss = []map[string]string{scanExpects.mapStrStr, scanExpects.mapStrStr}
			err = gconv.Scan(mapList, &mss)
			t.AssertNil(err)
			t.Assert(len(mss), len(mapList))
			for k, _ := range mapList {
				t.Assert(mapList[k]["Name"], mss[k]["Name"])
				t.Assert(mapList[k]["Place"], mss[k]["Place"])
			}

			var msa = []map[string]interface{}{scanExpects.mapStrAny, scanExpects.mapStrAny}
			err = gconv.Scan(mapList, &msa)
			t.AssertNil(err)
			t.Assert(len(msa), len(mapList))
			for k, _ := range mapList {
				t.Assert(mapList[k]["Name"], msa[k]["Name"])
				t.Assert(mapList[k]["Place"], msa[k]["Place"])
			}

			var maa = []map[interface{}]interface{}{scanExpects.mapAnyAny, scanExpects.mapAnyAny}
			err = gconv.Scan(mapList, &maa)
			t.AssertNil(err)
			t.Assert(len(maa), len(mapList))
			for k, _ := range mapList {
				t.Assert(mapList[k]["Name"], maa[k]["Name"])
				t.Assert(mapList[k]["Place"], maa[k]["Place"])
			}

			var ss = []scanStructTest{scanExpects.structSub, scanExpects.structSub}
			err = gconv.Scan(mapList, &ss)
			t.AssertNil(err)
			t.Assert(len(ss), len(mapList))
			for k, _ := range mapList {
				t.Assert(mapList[k]["Name"], ss[k].Name)
				t.Assert(mapList[k]["Place"], ss[k].Place)
			}

			var ssp = []*scanStructTest{scanExpects.structSubPtr, scanExpects.structSubPtr}
			err = gconv.Scan(mapList, &ssp)
			t.AssertNil(err)
			t.Assert(len(ssp), len(mapList))
			for k, _ := range mapList {
				t.Assert(mapList[k]["Name"], ssp[k].Name)
				t.Assert(mapList[k]["Place"], ssp[k].Place)
			}
		}
	})

	// Test for struct converting.
	gtest.C(t, func(t *gtest.T) {
		scanValuesTest := scanValueStructsTest
		for _, test := range scanValuesTest {
			var (
				err         error
				scanExpects = scanExpects
			)

			err = gconv.Scan(test, &scanExpects.mapStrStr)
			t.AssertNil(err)
			t.Assert(test.Name, scanExpects.mapStrStr["Name"])
			t.Assert(test.Place, scanExpects.mapStrStr["Place"])

			err = gconv.Scan(test, &scanExpects.mapStrAny)
			t.AssertNil(err)
			t.Assert(test.Name, scanExpects.mapStrAny["Name"])
			t.Assert(test.Place, scanExpects.mapStrAny["Place"])

			err = gconv.Scan(test, &scanExpects.mapAnyAny)
			t.AssertNil(err)
			t.Assert(test.Name, scanExpects.mapAnyAny["Name"])
			t.Assert(test.Place, scanExpects.mapAnyAny["Place"])

			err = gconv.Scan(test, &scanExpects.structSub)
			t.AssertNil(err)
			t.Assert(test.Name, scanExpects.structSub.Name)
			t.Assert(test.Place, scanExpects.structSub.Place)

			err = gconv.Scan(test, &scanExpects.structSubPtr)
			t.AssertNil(err)
			t.Assert(test.Name, scanExpects.structSubPtr.Name)
			t.Assert(test.Place, scanExpects.structSubPtr.Place)
		}
	})

	// Test for json converting.
	gtest.C(t, func(t *gtest.T) {
		scanValuesTest := scanValueJsonTest
		for _, test := range scanValuesTest {
			var (
				err         error
				scanExpects = scanExpects
			)

			err = gconv.Scan(test, &scanExpects.mapStrStr)
			t.AssertNil(err)
			t.Assert("Mars", scanExpects.mapStrStr["Name"])
			t.Assert("奥林帕斯山", scanExpects.mapStrStr["Place"])

			err = gconv.Scan(test, &scanExpects.mapStrAny)
			t.AssertNil(err)
			t.Assert("Mars", scanExpects.mapStrAny["Name"])
			t.Assert("奥林帕斯山", scanExpects.mapStrAny["Place"])

			err = gconv.Scan(test, &scanExpects.mapAnyAny)
			t.Assert(err, gerror.New(
				"json.UnmarshalUseNumber failed: json: cannot unmarshal object into Go value of type map[interface {}]interface {}",
			))

			err = gconv.Scan(test, &scanExpects.structSub)
			t.AssertNil(err)
			t.Assert("Mars", scanExpects.structSub.Name)
			t.Assert("奥林帕斯山", scanExpects.structSub.Place)

			err = gconv.Scan(test, &scanExpects.structSubPtr)
			t.AssertNil(err)
			t.Assert("Mars", scanExpects.structSubPtr.Name)
			t.Assert("奥林帕斯山", scanExpects.structSubPtr.Place)
		}
	})
}
