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
