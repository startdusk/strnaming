// Copyright (c) 2021 startdusk. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package strnaming

// Style defines style of language
type Style interface {
	// Cut according to the separator
	// enter a complete string (elem) and change it to your favorite code style
	Transformation(elem string) string
}
