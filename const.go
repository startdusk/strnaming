// Copyright (c) 2021 startdusk.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package strnaming

// ascii A -> a
const transNum = 'a' - 'A'

func toUpper(b byte) byte {
	return b - transNum
}

func toLower(b byte) byte {
	return b + transNum
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

func isUpper(b byte) bool {
	return 'A' <= b && b <= 'Z'
}

func isLower(b byte) bool {
	return 'a' <= b && b <= 'z'
}
