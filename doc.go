// Copyright (c) 2021 startdusk.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

/*
Package strnaming is used to replace the naming method of strings.

For example:
	package main

	import (
		"fmt"

		"github.com/startdusk/strnaming"
	)

	func main() {
		// camel
		camel := strnaming.NewCamel()
		fmt.Println(camel.Convert("camelcase_key")) // CamelcaseKey
	}

set first char lower

	fmt.Println(camel.Convert("user_id")) // UserId
	camel.WithLowerFirst(true)
	fmt.Println(camel.Convert("user_id")) // userId

set customize split

	camel.WithSplit('-')
	fmt.Println(camel.Convert("user-id")) // UserId

set cache

	camel.WithCache("user_id", "UserID")
	fmt.Println(camel.Convert("user_id")) // UserID

*/
package strnaming
