// Copyright (c) 2021 startdusk.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package strnaming

import "strings"

// Camel defines a cameler
type Camel struct {
	lowerFirst bool
	splits     []byte
	cache      map[string]string
}

// NewCamel creates a camel struct
func NewCamel() *Camel {
	return &Camel{}
}

// WithLowerFirst set first char lower
func (c *Camel) WithLowerFirst(b bool) *Camel {
	c.lowerFirst = b
	return c
}

// WithSplit set split char
func (c *Camel) WithSplit(b byte) *Camel {
	if b == 0 || c.containSplit(b) {
		return c
	}

	c.splits = append(c.splits, b)
	return c
}

// WithCache set cache
func (c *Camel) WithCache(k, v string) *Camel {
	if k == "" || v == "" {
		return c
	}
	if len(c.cache) == 0 {
		c.cache = make(map[string]string)
	}
	c.cache[k] = v
	return c
}

// Convert to camel string
func (c *Camel) Convert(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	if len(c.cache) > 0 {
		if a, ok := c.cache[s]; ok {
			return a
		}
	}

	var b strings.Builder
	b.Grow(len(s))
	// set first char defualt upper
	nextUpper := !c.lowerFirst
	for i, sl := 0, len(s); i < sl; i++ {
		v := s[i]
		curUpper, curLower := isUpper(v), isLower(v)

		if nextUpper {
			if curLower {
				v = toUpper(v)
			}
		} else if b.Len() == 0 {
			if curUpper {
				v = toLower(v)
			}
		}

		if curUpper || curLower {
			b.WriteByte(v)
			nextUpper = false
		} else if isNumber(v) {
			b.WriteByte(v)
			nextUpper = true
		} else if c.isSplitChar(v) && b.Len() != 0 {
			nextUpper = true
		}
	}

	return b.String()
}

func (c *Camel) isSplitChar(b byte) bool {
	// set defualt split char '_' if not set any split char
	if len(c.splits) == 0 {
		c.splits = append(c.splits, '_')
	}
	for _, v := range c.splits {
		if v == b {
			return true
		}
	}
	return false
}

func (c *Camel) containSplit(b byte) bool {
	for _, v := range c.splits {
		if v == b {
			return true
		}
	}
	return false
}
