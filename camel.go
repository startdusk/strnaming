package strnaming

import "strings"

// ToCamel("abc_efg").String() // AbcEfg
// ToCamel("abc_efg").WithLowerFirst(true).String() // abcEfg
// camel := NewCamel()
// for _, v := range []string{"abc_sa", "sdaf fsda", "sae.sdad"} {
// 	   fmt.Println(camel.Do(v))
// }

// ascii a -> A
const transNum = 'a' - 'A'

// CamelRes defines a camel struct
type CamelRes struct {
	str        string
	lowerFirst bool
	splits     []byte
}

// NewCamel creates a CamelRes
func NewCamel() *CamelRes {
	return &CamelRes{}
}

// ToCamel wrapper string to camel struct
func ToCamel(s string) *CamelRes {
	return &CamelRes{
		str: strings.TrimSpace(s),
	}
}

// WithLowerFirst set first char lower
func (c *CamelRes) WithLowerFirst(b bool) *CamelRes {
	c.lowerFirst = b
	return c
}

// WithSplit set split char
func (c *CamelRes) WithSplit(b byte) *CamelRes {
	if b == 0 || c.contains(b) {
		return c
	}

	c.splits = append(c.splits, b)
	return c
}

// Do multi use
func (c *CamelRes) Do(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	strArr := []byte(s)
	n := &strings.Builder{}
	n.Grow(len(strArr))
	// set first char defualt upper
	nextUpper := !c.lowerFirst
	for _, v := range strArr {
		isUpper := v >= 'A' && v <= 'Z'
		isLower := v >= 'a' && v <= 'z'

		if nextUpper {
			if isLower {
				v = toUpper(v)
			}
		} else if n.Len() == 0 {
			if isUpper {
				v = toLower(v)
			}
		}

		if isUpper || isLower {
			n.WriteByte(v)
			nextUpper = false
		} else if isNumberChar(v) {
			n.WriteByte(v)
			nextUpper = true
		} else if c.isSplitChar(v) && n.Len() != 0 {
			nextUpper = true
		}
	}
	return n.String()
}

func (c *CamelRes) String() string {
	return c.Do(c.str)
}

func (c *CamelRes) isSplitChar(b byte) bool {
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

func (c *CamelRes) contains(b byte) bool {
	for _, v := range c.splits {
		if v == b {
			return true
		}
	}
	return false
}

func toUpper(b byte) byte {
	return b - transNum
}

func toLower(b byte) byte {
	return b + transNum
}

func isNumberChar(b byte) bool {
	return b >= '0' && b <= '9'
}
