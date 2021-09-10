// Copyright (c) 2021 startdusk. All rights reserved.
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
func (c *Spacer) WithIgnore(ignores ...byte) *Spacer {
	for _, ignore := range ignores {
		if c.containsIgnore(ignore) {
			continue
		}
		c.ignores = append(c.ignores, ignore)
	}
	return c
}

// WithScreaming convert all char for upper
func (c *Spacer) WithScreaming(screaming bool) *Spacer {
	c.screaming = screaming
	return c
}

// WithCache set cache
func (c *Spacer) WithCache(key, value string) *Spacer {
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
func (c *Spacer) WithPrefix(prefix string) *Spacer {
	prefix = strings.TrimSpace(prefix)
	if prefix != "" {
		c.prefix = prefix
	}
	return c
}

// Convert to spacer string
func (c *Spacer) Convert(str string) string {
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

func (c *Spacer) do(str string) string {
	var b strings.Builder
	// Normally, most underscore named strings have 1 to 2 separators, so 2 is added here
	b.Grow(len(str) + 2)

	var prev byte
	var prevUpper bool
	for i, sl := 0, len(str); i < sl; i++ {
		cur := str[i]
		curUpper, curLower, curNum := isUpper(cur), isLower(cur), isNumber(cur)

		if c.screaming && curLower {
			cur = toUpper(cur)
		} else if !c.screaming && curUpper {
			cur = toLower(cur)
		}

		next, ok := nextVal(i, str)
		if !c.containsIgnore(prev) && ok {
			nextUpper, nextLower, nextNum := isUpper(next), isLower(next), isNumber(next)
			if (curUpper && (nextLower || nextNum)) || (curLower && (nextUpper || nextNum)) || (curNum && (nextUpper || nextLower)) {
				if prevUpper && curUpper && nextLower {
					b.WriteByte(c.delimiter)
				}
				b.WriteByte(cur)
				if curLower || curNum || nextNum {
					b.WriteByte(c.delimiter)
				}

				prev, prevUpper = cur, curUpper
				continue
			}
		}

		if !c.containsIgnore(cur) && !curUpper && !curLower && !curNum {
			b.WriteByte(c.delimiter)
		} else {
			b.WriteByte(cur)
		}
		prev, prevUpper = cur, curUpper
	}

	res := b.String()
	if c.prefix != "" {
		return c.prefix + string(c.delimiter) + res
	}

	return res
}

func (c *Spacer) containsIgnore(ignore byte) bool {
	if ignore == 0 || len(c.ignores) == 0 {
		return false
	}
	return strings.ContainsAny(string(ignore), string(c.ignores))
}

func nextVal(index int, str string) (byte, bool) {
	var b byte
	next := index + 1
	if next < len(str) {
		b = str[next]
		return b, true
	}
	return b, false
}
