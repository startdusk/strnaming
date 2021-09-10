// Copyright (c) 2021 startdusk. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package style

// DefaultLang for default
type DefaultLang struct{}

// NewDefault default code style
func NewDefault() *DefaultLang {
	return &DefaultLang{}
}

// Transformation transform to default code style
func (d DefaultLang) Transformation(elem string) string {
	return elem
}
