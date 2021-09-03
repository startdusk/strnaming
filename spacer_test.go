/*
 * MIT License
 *
 * Copyright (c) 2021 startdusk
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package strnaming

import "testing"

func TestSpacer(t *testing.T) {
	testSpacer(t)
}

func BenchmarkSpacer(b *testing.B) {
	testSpacer(b)
}

func testSpacer(tb testing.TB) {
	cases := []struct {
		test      string
		expect    string
		delimiter byte
		screaming bool
		ignores   []byte
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
	}

	for _, cc := range cases {
		spacer := &Spacer{
			delimiter: cc.delimiter,
			screaming: cc.screaming,
			ignores:   cc.ignores,
		}
		actual := spacer.Convert(cc.test)
		if actual != cc.expect {
			tb.Errorf("expect snake case %s, but got %s\n", cc.expect, actual)
		}
	}
}
