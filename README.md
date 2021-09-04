# strnaming

[![Godoc Reference](https://godoc.org/github.com/startdusk/strnaming?status.svg)](https://godoc.org/github.com/startdusk/strnaming)&nbsp;[![Go Report Card](https://goreportcard.com/badge/github.com/startdusk/strnaming)](https://goreportcard.com/report/github.com/startdusk/strnaming)&nbsp;![](https://img.shields.io/github/license/startdusk/strnaming)

Reference from [https://github.com/iancoleman/strcase](https://github.com/iancoleman/strcase) and changed a lot.

## Install

To start using strnaming, install Go and run `go get`:

```bash
$ go get -u github.com/startdusk/strnaming
```

This will retrieve the library.

## Example

```go
package main

import (
	"fmt"

	"github.com/startdusk/strnaming"
)

func main() {
	// camel
	camel := strnaming.NewCamel()
	fmt.Println(camel.Convert("camelcase_key")) // CamelcaseKey

	fmt.Println(camel.Convert("user_id")) // UserId
	camel.WithLowerFirst(true)
	fmt.Println(camel.Convert("user_id")) // userId

	camel.WithDelimiter('-')
	fmt.Println(camel.Convert("user-id")) // UserId

	camel.WithCache("user_id", "UserID")
	fmt.Println(camel.Convert("user_id")) // UserID

	// snake
	snake := strnaming.NewSnake()
	fmt.Println(snake.Convert("SnakeKey")) // snake_key

	snake.WithIgnore('-')
	fmt.Println(snake.Convert("My-IDCard")) // my-id_card

	snake.WithScreaming(true)
	fmt.Println(snake.Convert("SnakeKey")) // SNAKE_KEY

	snake.WithCache("UserID", "userid")
	fmt.Println(snake.Convert("UserID")) // userid

	// kebab
	kebab := strnaming.NewKebab()
	fmt.Println(kebab.Convert("KebabKey")) // kebab-key

	kebab.WithIgnore('@')
	kebab.WithIgnore('.')
	fmt.Println(kebab.Convert("ben_love@gmail.com")) // ben-love@gmail.com

	kebab.WithScreaming(true)
	fmt.Println(kebab.Convert("KebabKey")) // KEBAB_KEY

	kebab.WithCache("UserID", "User-Id")
	fmt.Println(kebab.Convert("UserID")) // User-Id
}


```

## TODO

- [x] Add prefix for string
- [ ] Cli for command line access
