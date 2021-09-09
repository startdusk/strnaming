// Copyright (c) 2021 startdusk.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package codes

// Style defines language naming style type
type Style string

const (
	// Golang naming style
	Golang Style = "Golang"
)

func (s Style) String() string {
	return string(s)
}
