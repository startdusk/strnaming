// Copyright (c) 2021 startdusk. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package strnaming

import (
	"testing"

	"github.com/startdusk/strnaming/style"
)

func TestCamel(t *testing.T) {
	testCamel(t)
}

func BenchmarkCamel(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testCamel(b)
	}
}

func testCamel(tb testing.TB) {
	cases := []struct {
		test       string
		expect     string
		upperFirst bool
		style      Style
		delimiters []byte
		prefix     string
		cacheKV    struct {
			key, val string
		}
	}{
		{
			test:   "_abc_def",
			expect: "abcDef",
		},
		{
			test:   "           _abc_def",
			expect: "abcDef",
		},
		{
			test:   "   _abc_DB_",
			expect: "abcDB",
		},
		{
			test:       "DNV.abc",
			expect:     "dNVAbc",
			delimiters: []byte{'.'},
		},
		{
			test:       "test_case",
			expect:     "testCase",
			delimiters: []byte{'_', '_'},
		},
		{
			test:   "many2many",
			expect: "many2Many",
		},
		{
			test:   "department_id",
			expect: "departmentId",
		},
		{
			test:       "department_id",
			expect:     "DepartmentId",
			upperFirst: true,
		},
		{
			test:       "_Department_id",
			expect:     "DepartmentId",
			upperFirst: true,
		},
		{
			test:   "_Department_id",
			expect: "departmentId",
		},
		{
			test:       "__prod__",
			expect:     "Prod",
			upperFirst: true,
		},
		{
			test:   "",
			expect: "",
		},
		{
			test:       "Golang",
			expect:     "Golang",
			upperFirst: true,
		},
		{
			test:   "__foo_Fzz",
			expect: "fooFzz",
		},
		{
			test:       "__foo_Fzz",
			expect:     "FooFzz",
			upperFirst: true,
		},
		{
			test:       "account_id",
			expect:     "AccountId",
			upperFirst: true,
		},
		{
			test:   "account_id",
			expect: "accountID",
			cacheKV: struct {
				key string
				val string
			}{
				key: "account_id",
				val: "accountID",
			},
		},
		{
			test:       "account_id",
			expect:     "myAccountId",
			upperFirst: true,
			prefix:     "my",
		},
		{
			test:   "account_id",
			expect: "myaccountId",
			prefix: "my",
		},
		{
			test:   "AccountID",
			expect: "accountID",
		},
		{
			test:   "http_test",
			expect: "httpTest",
		},
		{
			test:       "http_test",
			expect:     "HttpTest",
			upperFirst: true,
		},
		{
			test:   "http_test",
			expect: "HTTPTest",
			style:  style.NewGolang(),
		},
		{
			test:   "oauth",
			expect: "OAuth",
			style:  style.NewGolang(),
		},
		{
			test:   "oi_Auth",
			expect: "oiAuth",
		},
		{
			test:   "json_data",
			expect: "JSONData",
			style:  style.NewGolang(),
		},
	}

	for _, cc := range cases {
		actual := NewCamel().
			WithUpperFirst(cc.upperFirst).
			WithCache(cc.cacheKV.key, cc.cacheKV.val).
			WithPrefix(cc.prefix).
			WithDelimiter(cc.delimiters...).
			WithStyle(cc.style).
			Convert(cc.test)
		if actual != cc.expect {
			tb.Errorf("test camel case %s expect %s, but got %s\n", cc.test, cc.expect, actual)
		}
	}
}
