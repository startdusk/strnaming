// Copyright (c) 2021 startdusk.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package strnaming

import "testing"

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
		lowerFirst bool
		delimiters []byte
		prefix     string
		cacheKV    struct {
			key, val string
		}
	}{
		{
			test:   "_abc_def",
			expect: "AbcDef",
		},
		{
			test:   "           ***_abc_def",
			expect: "AbcDef",
		},
		{
			test:   "   ***_abc_DB_",
			expect: "AbcDB",
		},
		{
			test:       "DNV.abc",
			expect:     "DNVAbc",
			delimiters: []byte{'.'},
		},
		{
			test:       "test_case",
			expect:     "TestCase",
			delimiters: []byte{'_', '_'},
		},
		{
			test:   "many2many",
			expect: "Many2Many",
		},
		{
			test:   "department_id",
			expect: "DepartmentId",
		},
		{
			test:       "department_id",
			expect:     "departmentId",
			lowerFirst: true,
		},
		{
			test:       "_Department_id",
			expect:     "departmentId",
			lowerFirst: true,
		},
		{
			test:   "_Department_id",
			expect: "DepartmentId",
		},
		{
			test:       "__prod__",
			expect:     "prod",
			lowerFirst: true,
		},
		{
			test:   "",
			expect: "",
		},
		{
			test:       "Golang",
			expect:     "golang",
			lowerFirst: true,
		},
		{
			test:   "__foo_ Fzz**",
			expect: "FooFzz",
		},
		{
			test:       "__foo_ Fzz**",
			expect:     "fooFzz",
			lowerFirst: true,
		},
		{
			test:       "account_id",
			expect:     "accountId",
			lowerFirst: true,
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
			expect:     "myaccountId",
			lowerFirst: true,
			prefix:     "my",
		},
		{
			test:   "account_id",
			expect: "myAccountId",
			prefix: "my",
		},
	}

	for _, cc := range cases {
		camel := NewCamel()
		camel.WithLowerFirst(cc.lowerFirst)
		camel.WithCache(cc.cacheKV.key, cc.cacheKV.val)
		camel.WithPrefix(cc.prefix)
		for _, v := range cc.delimiters {
			camel.WithDelimiter(v)
		}

		actual := camel.Convert(cc.test)
		if actual != cc.expect {
			tb.Errorf("expect camel case %s, but got %s\n", cc.expect, actual)
		}
	}
}
