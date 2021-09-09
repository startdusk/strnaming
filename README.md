# strnaming

[![Godoc Reference](https://godoc.org/github.com/startdusk/strnaming?status.svg)](https://godoc.org/github.com/startdusk/strnaming)&nbsp;[![Go Report Card](https://goreportcard.com/badge/github.com/startdusk/strnaming)](https://goreportcard.com/report/github.com/startdusk/strnaming)&nbsp;[![](https://img.shields.io/github/license/startdusk/strnaming)](https://github.com/startdusk/strnaming/blob/main/LICENSE)&nbsp;[![Goproxy.cn](https://goproxy.cn/stats/github.com/startdusk/strnaming/badges/download-count.svg)](https://goproxy.cn/stats/github.com/startdusk/strnaming/badges/download-count.svg)

Reference from [https://github.com/iancoleman/strcase](https://github.com/iancoleman/strcase) and changed a lot.

## Contents

- [strnaming](#strnaming)
  - [Contents](#contents)
  - [API Examples](#api-examples)
    - [Install](#install)
    - [Quick start](#quick-start)
      - [camelCase](#camel)
      - [snake_case](#snake)
      - [kebab-case](#kebab)
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

#### camel

```go
package main

import (
	"fmt"

	"github.com/startdusk/strnaming"
)

func main() {
	// camel case
	camel := strnaming.NewCamel()
	fmt.Println(camel.Convert("camelcase_key")) // camelcaseKey

	fmt.Println(camel.Convert("user_id")) // userId

	camel.WithDelimiter('-')
	fmt.Println(camel.Convert("user-id")) // userId

	camel.WithUpperFirst(true)
	fmt.Println(camel.Convert("user_id")) // UserId

	camel.WithCache("user_id", "UserID")
	fmt.Println(camel.Convert("user_id")) // UserID

	camel.WithPrefix("My")
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
	// snake case
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
	// kebab case
	kebab := strnaming.NewKebab()
	fmt.Println(kebab.Convert("KebabKey")) // kebab-key

	kebab.WithIgnore('@', '.')
	fmt.Println(kebab.Convert("ben_love@gmail.com")) // ben-love@gmail.com

	kebab.WithScreaming(true)
	fmt.Println(kebab.Convert("KebabKey")) // KEBAB-KEY

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

convert json keys to camelcase keys, eg:

```json
// ./testdata/test.json

{
    "test_url": "http://json-schema.org/draft-04/schema",
    "another_case": [
        {
	    "sub_case": [
		{
		    "for_ready": 1234,
		    "bba_media": "hahahaha"
		}
	    ]
	},
	{
	    "sub_url_two": [
		{
		    "for_ready_two": "ben",
		    "bba_media_two": 2021,
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
    "anotherCase": [
        {
            "subCase": [
                {
                    "bbaMedia": "hahahaha",
                    "forReady": 1234       
                }
            ]
        },
        {
            "subUrlTwo": [
                {
                    "bbaMediaTwo": 2021,
                    "forReadyTwo": "ben",
                    "keySpace": [
                        [
                            [
                                {
                                    "lowCode": true
                                }
                            ]
                        ]
                    ]
                }
            ]
        }
    ],
    "testUrl": "http://json-schema.org/draft-04/schema"
}
```

## TODO

- [x] Add prefix for string
- [x] Cli for command line access
- [ ] Support Golang language naming style
