// Copyright (c) 2021 startdusk.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package strnaming

import "strings"

// Camel defines a cameler
type Camel struct {
	upperFirst bool
	prefix     string
	delimiters []byte
	cache      map[string]string
}

// NewCamel creates a camel struct
func NewCamel() *Camel {
	return &Camel{}
}

// WithUpperFirst set first char upper
func (c *Camel) WithUpperFirst(upperFirst bool) *Camel {
	c.upperFirst = upperFirst
	return c
}

// WithDelimiter set delimiter char
func (c *Camel) WithDelimiter(delimiters ...byte) *Camel {
	for _, delimiter := range delimiters {
		if c.containsDelimiter(delimiter) {
			continue
		}

		c.delimiters = append(c.delimiters, delimiter)
	}
	return c
}

// WithCache set cache
func (c *Camel) WithCache(key, value string) *Camel {
	if key == "" || value == "" {
		return c
	}
	if len(c.cache) == 0 {
		c.cache = make(map[string]string)
	}
	c.cache[key] = value
	return c
}

// WithPrefix set prefix
func (c *Camel) WithPrefix(prefix string) *Camel {
	prefix = strings.TrimSpace(prefix)
	if prefix != "" {
		c.prefix = prefix
	}
	return c
}

// Convert to camel string
func (c *Camel) Convert(str string) string {
	str = strings.TrimSpace(str)
	if str == "" {
		return str
	}

	if len(c.cache) > 0 {
		if a, ok := c.cache[str]; ok {
			return a
		}
	}

	return c.do(str)
}

func (c *Camel) do(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	// first char will be lower
	// upperFirst default false if not set true
	nextUpper := c.upperFirst
	for i, sl := 0, len(str); i < sl; i++ {
		cur := str[i]
		curUpper, curLower, curNum := isUpper(cur), isLower(cur), isNumber(cur)

		if nextUpper {
			if curLower {
				cur = toUpper(cur)
			}
		} else if b.Len() == 0 {
			if curUpper {
				cur = toLower(cur)
			}
		}

		if curUpper || curLower {
			b.WriteByte(cur)
			nextUpper = false
		} else if curNum {
			b.WriteByte(cur)
			nextUpper = true
		} else if c.isDelimiterChar(cur) && b.Len() != 0 {
			nextUpper = true
		}
	}

	return c.prefix + b.String()
}

func (c *Camel) isDelimiterChar(delimiter byte) bool {
	// set default delimiter char '_' if not set any delimiter char
	if len(c.delimiters) == 0 {
		c.delimiters = append(c.delimiters, snakeDelimiter)
	}

	return c.containsDelimiter(delimiter)
}

func (c *Camel) containsDelimiter(delimiter byte) bool {
	if delimiter == 0 || len(c.delimiters) == 0 {
		return false
	}

	return strings.ContainsAny(string(delimiter), string(c.delimiters))
}
