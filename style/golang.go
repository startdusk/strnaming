// Copyright (c) 2021 startdusk. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package style

import "strings"

// commonInitialisms, taken from
// https://github.com/golang/lint/blob/206c0f020eba0f7fbcfbc467a5eb808037df2ed6/lint.go#L731
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"ETA":   true,
	"GPU":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"OS":    true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
	"OAuth": true,
}

// add exceptions here for things that are not automatically convertable
// toCamelSpecial, taken from
// https://github.com/serenize/snaker/blob/a683aaf2d516deecd70cad0c72e3ca773ecfcef0/snaker.go#L147-L150
var toCamelSpecial = map[string]string{
	"oauth": "OAuth",
}

// Golang implementations Style interface
type Golang struct{}

// NewGolang creates code style for golang
func NewGolang() *Golang {
	return &Golang{}
}

// Transformation transform to golang code style
func (g Golang) Transformation(elem string) string {
	// special case
	if special, ok := toCamelSpecial[strings.ToLower(elem)]; ok {
		return special
	} else if upper := strings.ToUpper(elem); commonInitialisms[upper] {
		return upper
	}
	return elem
}
