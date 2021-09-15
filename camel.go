// Copyright (c) 2021 startdusk. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package strnaming

import (
	"strings"

	"github.com/startdusk/strnaming/style"
)

// Camel defines a cameler
type Camel struct {
	style      Style
	upperFirst bool
	prefix     string
	delimiters []byte
	cache      map[string]string
}

// NewCamel creates a camel struct
func NewCamel() *Camel {
	return &Camel{
		style: style.NewDefault(),
	}
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

// WithStyle using lang style
func (c *Camel) WithStyle(style Style) *Camel {
	if style != nil {
		c.style = style
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

	for {
		var word []byte
		i, sl := 0, len(str)
		for i < sl {
			cur := str[i]
			if c.isDelimiterChar(cur) {
				break
			}

			curUpper, curLower, curNum := isUpper(cur), isLower(cur), isNumber(cur)
			// special first char
			if i == 0 && b.Len() == 0 {
				if c.upperFirst {
					if curLower {
						cur = toUpper(cur)
					}
				} else if curUpper {
					cur = toLower(cur)
				}
			} else if i == 0 && curLower {
				cur = toUpper(cur)
			}
			word = append(word, cur)
			if curNum {
				break
			}
			i++
		}

		elem := c.style.Transformation(string(word))
		b.WriteString(elem)

		if i == len(str) {
			break
		}

		str = str[i+1:]
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
