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

// Spacer defines a spacer struct
type Spacer struct {
	delimiter byte
	screaming bool
	ignores   []byte
	cache     map[string]string
}

// WithIgnore ignore special char eg $
func (c *Spacer) WithIgnore(b byte) *Spacer {
	if b == 0 || c.containIgnore(b) {
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
		c.cache = make(map[string]string, 0)
	}
	c.cache[k] = v
	return c
}

// Convert to spacer string
func (c *Spacer) Convert(s string) string {
	return c.do(s)
}

func (c *Spacer) do(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	if len(c.cache) > 0 {
		if a, ok := c.cache[s]; ok {
			return a
		}
	}

	var n strings.Builder
	n.Grow(len(s) + 2)

	var prevUpper bool
	var prev byte
	for i, sl := 0, len(s); i < sl; i++ {
		cur := s[i]
		curUpper, curLower := isUpper(cur), isLower(cur)

		if c.screaming && curLower {
			cur = toUpper(cur)
		} else if !c.screaming && curUpper {
			cur = toLower(cur)
		}

		if next, ok := nextVal(i, s); ok {
			curNum := isNumber(cur)
			nextUpper := isUpper(next)
			nextLower := isLower(next)
			nextNum := isNumber(next)
			if (curUpper && (nextLower || nextNum)) || (curLower && (nextUpper || nextNum)) || (curNum && (nextUpper || nextLower)) {
				if !c.containIgnore(prev) {
					if curUpper && nextLower {
						if i > 0 && prevUpper {
							n.WriteByte(c.delimiter)
						}
					}
					n.WriteByte(cur)
					if curLower || curNum || nextNum {
						n.WriteByte(c.delimiter)
					}

					prev, prevUpper = cur, curUpper
					continue
				}
			}
		}

		if !curUpper && !curLower && !c.containIgnore(cur) {
			n.WriteByte(c.delimiter)
		} else {
			n.WriteByte(cur)
		}
		prev, prevUpper = cur, curUpper
	}

	return n.String()
}

func (c *Spacer) containIgnore(b byte) bool {
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
