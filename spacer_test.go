// Copyright (c) 2021 startdusk.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package strnaming

import (
	"testing"
)

func TestSpacer(t *testing.T) {
	testSpacer(t)
}

func BenchmarkSpacer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testSpacer(b)
	}
}

func testSpacer(tb testing.TB) {
	cases := []struct {
		test      string
		expect    string
		delimiter byte
		screaming bool
		prefix    string
		ignores   []byte
		cacheKV   struct {
			key, val string
		}
	}{
		{
			test:      "AccountID",
			expect:    "account_id",
			delimiter: '_',
		},
		{
			test:      "JSONStringify",
			expect:    "json_stringify",
			delimiter: '_',
		},
		{
			test:      "MongoDB",
			expect:    "mongo_db",
			delimiter: '_',
		},
		{
			test:      "MySQL",
			expect:    "my_sql",
			delimiter: '_',
		},
		{
			test:      "SQLServer",
			expect:    "sql_server",
			delimiter: '_',
		},
		{
			test:      "Many2Many",
			expect:    "many_2_many",
			delimiter: '_',
		},
		{
			test:      "DB_USER",
			expect:    "db_user",
			delimiter: '_',
		},
		{
			test:      "1A2",
			expect:    "1_a_2",
			delimiter: '_',
		},
		{
			test:      "1A2",
			expect:    "1_a_2",
			delimiter: '_',
		},
		{
			test:      "DB_USER",
			expect:    "DB_USER",
			delimiter: '_',
			screaming: true,
		},
		{
			test:      "AccountID",
			expect:    "account-id",
			delimiter: '-',
		},
		{
			test:      "JSONStringify",
			expect:    "json-stringify",
			delimiter: '-',
		},
		{
			test:      "MongoDB",
			expect:    "mongo-db",
			delimiter: '-',
		},
		{
			test:      "MySQL",
			expect:    "my-sql",
			delimiter: '-',
		},
		{
			test:      "SQLServer",
			expect:    "sql-server",
			delimiter: '-',
		},
		{
			test:      "Many2Many",
			expect:    "many-2-many",
			delimiter: '-',
		},
		{
			test:      "DB_USER",
			expect:    "db-user",
			delimiter: '-',
		},
		{
			test:      "DB_USER",
			expect:    "DB-USER",
			delimiter: '-',
			screaming: true,
		},
		{
			test:      "DbUser",
			expect:    "DB-USER",
			delimiter: '-',
			screaming: true,
		},
		{
			test:      "ben_love@gmail.com",
			expect:    "ben-love@gmail.com",
			delimiter: '-',
			ignores:   []byte{'@', '.'},
		},
		{
			test:      "",
			expect:    "",
			delimiter: '-',
		},
		{
			test:      "JSONData",
			expect:    "user_json_data",
			delimiter: '_',
			prefix:    "user",
		},
		{
			test:      "JSONData",
			expect:    "user-json-data",
			delimiter: '-',
			prefix:    "user",
		},
		{
			test:      "",
			expect:    "",
			delimiter: '-',
			ignores:   []byte{0},
		},
		{
			test:      "abcEFG",
			expect:    "abc_efg",
			delimiter: '_',
			ignores:   []byte{'@', '@'},
		},
		{
			test:      "",
			expect:    "",
			delimiter: '-',
			cacheKV: struct {
				key string
				val string
			}{
				key: "",
				val: "",
			},
		},
		{
			test:      "testCase",
			expect:    "TESTCase",
			delimiter: '_',
			cacheKV: struct {
				key string
				val string
			}{
				key: "testCase",
				val: "TESTCase",
			},
		},
	}

	for _, cc := range cases {
		spacer := &Spacer{
			delimiter: cc.delimiter,
		}
		actual := spacer.WithScreaming(cc.screaming).
			WithCache(cc.cacheKV.key, cc.cacheKV.val).
			WithPrefix(cc.prefix).
			WithIgnore(cc.ignores...).
			Convert(cc.test)

		if actual != cc.expect {
			tb.Errorf("expect snake case %s, but got %s\n", cc.expect, actual)
		}
	}
}
