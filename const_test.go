// Copyright (c) 2021 startdusk. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package strnaming

import "testing"

func TestToLower(t *testing.T) {
	cases := []struct {
		test   byte
		expect byte
		result bool
	}{
		{
			test:   'a',
			expect: 'a',
			result: false,
		},
		{
			test:   'A',
			expect: 'a',
			result: true,
		},
		{
			test:   '1',
			expect: 'a',
			result: false,
		},
		{
			test:   ' ',
			expect: 'a',
			result: false,
		},
		{
			test:   '1',
			expect: '2',
			result: false,
		},
		{
			test:   0,
			expect: 0,
			result: false,
		},
		{
			test:   'a',
			expect: 'A',
			result: false,
		},
	}

	for _, cc := range cases {
		t.Run(string(cc.test), func(t *testing.T) {
			actual := toLower(cc.test)
			if !cc.result && actual == cc.expect {
				t.Errorf("toLower test %c want to %t, expect: %c but got: %c", cc.test, cc.result, cc.expect, actual)
			}

			if cc.result && actual != cc.expect {
				t.Errorf("toLower test %c want to %t, expect: %c but got: %c", cc.test, cc.result, cc.expect, actual)
			}
		})
	}
}

func TestToUpper(t *testing.T) {
	cases := []struct {
		test   byte
		expect byte
		result bool
	}{
		{
			test:   'a',
			expect: 'a',
			result: false,
		},
		{
			test:   'A',
			expect: 'a',
			result: false,
		},
		{
			test:   '1',
			expect: 'a',
			result: false,
		},
		{
			test:   ' ',
			expect: 'a',
			result: false,
		},
		{
			test:   '1',
			expect: '2',
			result: false,
		},
		{
			test:   0,
			expect: 0,
			result: false,
		},
		{
			test:   'a',
			expect: 'A',
			result: true,
		},
	}

	for _, cc := range cases {
		t.Run(string(cc.test), func(t *testing.T) {
			actual := toUpper(cc.test)
			if !cc.result && actual == cc.expect {
				t.Errorf("toUpper test %c result want to %t, expect: %c but got: %c", cc.test, cc.result, cc.expect, actual)
			}

			if cc.result && actual != cc.expect {
				t.Errorf("toUpper test %c result want to %t, expect: %c but got: %c", cc.test, cc.result, cc.expect, actual)
			}
		})
	}
}

func TestIsNumber(t *testing.T) {
	cases := []struct {
		test   byte
		result bool
	}{
		{
			test:   '1',
			result: true,
		},
		{
			test:   '9',
			result: true,
		},
		{
			test:   'b',
			result: false,
		},
		{
			test:   '@',
			result: false,
		},
		{
			test:   0,
			result: false,
		},
		{
			test:   '0',
			result: true,
		},
	}

	for _, cc := range cases {
		t.Run(string(cc.test), func(t *testing.T) {
			actual := isNumber(cc.test)
			if cc.result != actual {
				t.Errorf("isNumber test %c want to %t, but got: %t", cc.test, cc.result, actual)
			}
		})
	}
}

func TestIsLower(t *testing.T) {
	cases := []struct {
		test   byte
		result bool
	}{
		{
			test:   '1',
			result: false,
		},
		{
			test:   '9',
			result: false,
		},
		{
			test:   'b',
			result: true,
		},
		{
			test:   '@',
			result: false,
		},
		{
			test:   0,
			result: false,
		},
		{
			test:   '0',
			result: false,
		},
		{
			test:   'A',
			result: false,
		},
		{
			test:   'z',
			result: true,
		},
	}

	for _, cc := range cases {
		t.Run(string(cc.test), func(t *testing.T) {
			actual := isLower(cc.test)
			if cc.result != actual {
				t.Errorf("isLower test %c want to %t, but got: %t", cc.test, cc.result, actual)
			}
		})
	}
}

func TestIsUpper(t *testing.T) {
	cases := []struct {
		test   byte
		result bool
	}{
		{
			test:   '1',
			result: false,
		},
		{
			test:   '9',
			result: false,
		},
		{
			test:   'b',
			result: false,
		},
		{
			test:   '@',
			result: false,
		},
		{
			test:   0,
			result: false,
		},
		{
			test:   '0',
			result: false,
		},
		{
			test:   'A',
			result: true,
		},
		{
			test:   'z',
			result: false,
		},
		{
			test:   'Z',
			result: true,
		},
	}

	for _, cc := range cases {
		t.Run(string(cc.test), func(t *testing.T) {
			actual := isUpper(cc.test)
			if cc.result != actual {
				t.Errorf("isUpper test %c want to %t, but got: %t", cc.test, cc.result, actual)
			}
		})
	}
}
