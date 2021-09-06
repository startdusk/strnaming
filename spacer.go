// Copyright (c) 2021 startdusk.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package strnaming

import "strings"

// Spacer defines a spacer struct
type Spacer struct {
	delimiter byte
	screaming bool
	prefix    string
	ignores   []byte
	cache     map[string]string
}

// WithIgnore ignore special char eg $
func (c *Spacer) WithIgnore(b byte) *Spacer {
	if b == 0 || c.containsIgnore(b) {
		return c
	}
	c.ignores = append(c.ignores, b)
	return c
}

// WithScreaming convert all char for upper
func (c *Spacer) WithScreaming(b bool) *Spacer {
	c.screaming = b
	return c
}

// WithCache set cache
func (c *Spacer) WithCache(k, v string) *Spacer {
	if k == "" || v == "" {
		return c
	}
	if len(c.cache) == 0 {
		c.cache = make(map[string]string)
	}
	c.cache[k] = v
	return c
}

// WithPrefix set prefix
func (c *Spacer) WithPrefix(s string) *Spacer {
	s = strings.TrimSpace(s)
	if s != "" {
		c.prefix = s
	}
	return c
}

// Convert to spacer string
func (c *Spacer) Convert(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	if len(c.cache) > 0 {
		if a, ok := c.cache[s]; ok {
			return a
		}
	}

	return c.do(s)
}

func (c *Spacer) do(s string) string {

	var n strings.Builder
	// Normally, most underscore named strings have 1 to 2 separators, so 2 is added here
	n.Grow(len(s) + 2)

	var prev byte
	var prevUpper bool
	for i, sl := 0, len(s); i < sl; i++ {
		cur := s[i]
		curUpper, curLower, curNum := isUpper(cur), isLower(cur), isNumber(cur)

		if c.screaming && curLower {
			cur = toUpper(cur)
		} else if !c.screaming && curUpper {
			cur = toLower(cur)
		}

		next, ok := nextVal(i, s)
		if !c.containsIgnore(prev) && ok {
			nextUpper, nextLower, nextNum := isUpper(next), isLower(next), isNumber(next)
			if (curUpper && (nextLower || nextNum)) || (curLower && (nextUpper || nextNum)) || (curNum && (nextUpper || nextLower)) {
				if prevUpper && curUpper && nextLower {
					n.WriteByte(c.delimiter)
				}
				n.WriteByte(cur)
				if curLower || curNum || nextNum {
					n.WriteByte(c.delimiter)
				}

				prev, prevUpper = cur, curUpper
				continue
			}
		}

		if !c.containsIgnore(cur) && !curUpper && !curLower && !curNum {
			n.WriteByte(c.delimiter)
		} else {
			n.WriteByte(cur)
		}
		prev, prevUpper = cur, curUpper
	}

	res := n.String()
	if c.prefix != "" {
		return c.prefix + string(c.delimiter) + res
	}

	return res
}

func (c *Spacer) containsIgnore(b byte) bool {
	if b == 0 || len(c.ignores) == 0 {
		return false
	}
	return strings.ContainsAny(string(b), string(c.ignores))
}

func nextVal(i int, s string) (byte, bool) {
	var b byte
	next := i + 1
	if next < len(s) {
		b = s[next]
		return b, true
	}
	return b, false
}
