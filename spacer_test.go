// Copyright (c) 2021 startdusk. All rights reserved.
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
			delimiter: snakeDelimiter,
		},
		{
			test:      "JSONStringify",
			expect:    "json_stringify",
			delimiter: snakeDelimiter,
		},
		{
			test:      "MongoDB",
			expect:    "mongo_db",
			delimiter: snakeDelimiter,
		},
		{
			test:      "MySQL",
			expect:    "my_sql",
			delimiter: snakeDelimiter,
		},
		{
			test:      "SQLServer",
			expect:    "sql_server",
			delimiter: snakeDelimiter,
		},
		{
			test:      "Many2Many",
			expect:    "many_2_many",
			delimiter: snakeDelimiter,
		},
		{
			test:      "DB_USER",
			expect:    "db_user",
			delimiter: snakeDelimiter,
		},
		{
			test:      "1A2",
			expect:    "1_a_2",
			delimiter: snakeDelimiter,
		},
		{
			test:      "1A2",
			expect:    "1_a_2",
			delimiter: snakeDelimiter,
		},
		{
			test:      "DB_USER",
			expect:    "DB_USER",
			delimiter: snakeDelimiter,
			screaming: true,
		},
		{
			test:      "AccountID",
			expect:    "account-id",
			delimiter: kebabDelimiter,
		},
		{
			test:      "JSONStringify",
			expect:    "json-stringify",
			delimiter: kebabDelimiter,
		},
		{
			test:      "MongoDB",
			expect:    "mongo-db",
			delimiter: kebabDelimiter,
		},
		{
			test:      "MySQL",
			expect:    "my-sql",
			delimiter: kebabDelimiter,
		},
		{
			test:      "SQLServer",
			expect:    "sql-server",
			delimiter: kebabDelimiter,
		},
		{
			test:      "Many2Many",
			expect:    "many-2-many",
			delimiter: kebabDelimiter,
		},
		{
			test:      "DB_USER",
			expect:    "db-user",
			delimiter: kebabDelimiter,
		},
		{
			test:      "DB_USER",
			expect:    "DB-USER",
			delimiter: kebabDelimiter,
			screaming: true,
		},
		{
			test:      "DbUser",
			expect:    "DB-USER",
			delimiter: kebabDelimiter,
			screaming: true,
		},
		{
			test:      "ben_love@gmail.com",
			expect:    "ben-love@gmail.com",
			delimiter: kebabDelimiter,
			ignores:   []byte{'@', '.'},
		},
		{
			test:      "",
			expect:    "",
			delimiter: kebabDelimiter,
		},
		{
			test:      "JSONData",
			expect:    "user_json_data",
			delimiter: snakeDelimiter,
			prefix:    "user",
		},
		{
			test:      "JSONData",
			expect:    "user-json-data",
			delimiter: kebabDelimiter,
			prefix:    "user",
		},
		{
			test:      "",
			expect:    "",
			delimiter: kebabDelimiter,
			ignores:   []byte{0},
		},
		{
			test:      "abcEFG",
			expect:    "abc_efg",
			delimiter: snakeDelimiter,
			ignores:   []byte{'@', '@'},
		},
		{
			test:      "",
			expect:    "",
			delimiter: kebabDelimiter,
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
			delimiter: snakeDelimiter,
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
			tb.Errorf("expect case %s, but got %s\n", cc.expect, actual)
		}
	}
}
