// Copyright (c) 2021 startdusk. All rights reserved.
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
		fmt.Println(camel.Convert("camelcase_key")) // camelcaseKey
	}

set preifx
	camel.WithPrefix("My")
	fmt.Println(camel.Convert("user_id")) // MyuserId

set first char upper

	fmt.Println(camel.Convert("user_id")) // userId
	camel.WithUpperFirst(true)
	fmt.Println(camel.Convert("user_id")) // UserId

set prefix

	camel.WithPrefix("My")
	fmt.Println(camel.Convert("user_name")) // MyuserName

set customize delimiter

	camel.WithDelimiter('-')
	fmt.Println(camel.Convert("user-id")) // userId

set cache

	camel.WithCache("user_id", "UserID")
	fmt.Println(camel.Convert("user_id")) // UserID

set style
	package main

	import (
		"fmt"

		"github.com/startdusk/strnaming"
		"github.com/startdusk/strnaming/style"
	)

	func main() {
		// camel
		camel := strnaming.NewCamel()
		camel.WithStyle(style.NewGolang())
		fmt.Println(camel.Convert("http_test")) // HTTPTest
		fmt.Println(camel.Convert("json_data")) // JSONData
	}

*/
package strnaming
