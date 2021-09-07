# strnaming

[![Godoc Reference](https://godoc.org/github.com/startdusk/strnaming?status.svg)](https://godoc.org/github.com/startdusk/strnaming)&nbsp;[![Go Report Card](https://goreportcard.com/badge/github.com/startdusk/strnaming)](https://goreportcard.com/report/github.com/startdusk/strnaming)&nbsp;[![](https://img.shields.io/github/license/startdusk/strnaming)](https://github.com/startdusk/strnaming/blob/main/LICENSE)&nbsp;[![Goproxy.cn](https://goproxy.cn/stats/github.com/startdusk/strnaming/badges/download-count.svg)](https://goproxy.cn/stats/github.com/startdusk/strnaming/badges/download-count.svg)

Reference from [https://github.com/iancoleman/strcase](https://github.com/iancoleman/strcase) and changed a lot.

## Contents

- [strnaming](#strnaming)
  - [Contents](#contents)
  - [API Examples](#api-examples)
    - [Install](#install)
    - [Quick start](#quick-start)
      - [camelcase](#camelcase)
      - [snake](#snake)
      - [kebab](#kebab)
  - [CLI Examples](#cli-examples)
    - [Install](#install-1)
    - [Quick start](#quick-start-1)
  - [TODO](#todo)

## API Examples

### Install

To start using strnaming, install Go and run `go get`:

```bash
$ go get -u github.com/startdusk/strnaming
```

This will retrieve the library.

### Quick start

#### camelcase

```go
package main

import (
	"fmt"

	"github.com/startdusk/strnaming"
)

func main() {
	// camelcase
	camel := strnaming.NewCamel()
	fmt.Println(camel.Convert("camelcase_key")) // CamelcaseKey

	fmt.Println(camel.Convert("user_id")) // UserId

	camel.WithDelimiter('-')
	fmt.Println(camel.Convert("user-id")) // UserId

	camel.WithLowerFirst(true)
	fmt.Println(camel.Convert("user_id")) // userId

	camel.WithCache("user_id", "UserID")
	fmt.Println(camel.Convert("user_id")) // UserID

	camel.WithPrefix("My")
	camel.WithLowerFirst(false)
	fmt.Println(camel.Convert("user_name")) // MyUserName
}

```

#### snake

```go
package main

import (
	"fmt"

	"github.com/startdusk/strnaming"
)

func main() {
	// snake
	snake := strnaming.NewSnake()
	fmt.Println(snake.Convert("SnakeKey")) // snake_key

	snake.WithIgnore('-')
	fmt.Println(snake.Convert("My-IDCard")) // my-id_card

	snake.WithScreaming(true)
	fmt.Println(snake.Convert("SnakeKey")) // SNAKE_KEY

	snake.WithCache("UserID", "userid")
	fmt.Println(snake.Convert("UserID")) // userid

	snake.WithPrefix("go")
	snake.WithScreaming(false)
	fmt.Println(snake.Convert("PageSize")) // go_page_size
}

```

#### kebab

```go
package main

import (
	"fmt"

	"github.com/startdusk/strnaming"
)

func main() {
	// kebab
	kebab := strnaming.NewKebab()
	fmt.Println(kebab.Convert("KebabKey")) // kebab-key

	kebab.WithIgnore('@', '.')
	fmt.Println(kebab.Convert("ben_love@gmail.com")) // ben-love@gmail.com

	kebab.WithScreaming(true)
	fmt.Println(kebab.Convert("KebabKey")) // KEBAB_KEY

	kebab.WithCache("UserID", "User-Id")
	fmt.Println(kebab.Convert("UserID")) // User-Id

	kebab.WithPrefix("go")
	kebab.WithScreaming(false)
	fmt.Println(kebab.Convert("PageSize")) // go-page-size
}

```

## CLI Examples

### Install

To start using strnaming in command line, install Go and run `go get`:

```bash
$ go get -u github.com/startdusk/strnaming/cmd/strnaming
```

### Quick start

convert json keys to camelcase key, eg:

```json
// ./testdata/test.json
{
  "test_url": "http://json-schema.org/draft-04/schema",
  "another_url": [
    {
      "sub_url": [
        {
          "for_lady": 1234,
          "bba_media": "hahahaha"
        }
      ]
    },
    {
      "sub_url2": [
        {
          "for_lady2": "ben",
          "bba_media2": 2021,
          "key_space": [
            [
              [
                {
                  "low_code": true
                }
              ]
            ]
          ]
        }
      ]
    }
  ]
}
```

command:

```bash
$ strnaming c -f=./testdata/test.json
```

output:

```json
{
  "AnotherUrl": [
    {
      "SubUrl": [
        {
          "BbaMedia": "hahahaha",
          "ForLady": 1234
        }
      ]
    },
    {
      "SubUrl2": [
        {
          "BbaMedia2": 2021,
          "ForLady2": "ben",
          "KeySpace": [
            [
              [
                {
                  "LowCode": true
                }
              ]
            ]
          ]
        }
      ]
    }
  ],
  "TestUrl": "http://json-schema.org/draft-04/schema"
}
```

## TODO

- [x] Add prefix for string
- [x] Cli for command line access
