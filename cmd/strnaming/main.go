// Copyright (c) 2021 startdusk. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/startdusk/strnaming"
	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	{
		Name:    "camel",
		Aliases: []string{"c"},
		Usage:   "convert to camel string",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:    "file",
				Usage:   "input a json file path (eg: /path/to/strnaming.json)",
				Aliases: []string{"f"},
			},
			&cli.StringFlag{
				Name:    "json",
				Usage:   "input a json",
				Aliases: []string{"j"},
			},
			&cli.StringFlag{
				Name:        "delimiter",
				Usage:       "using custom delimiter",
				Aliases:     []string{"d"},
				DefaultText: "_",
			},
			&cli.BoolFlag{
				Name:        "upperFirst",
				Usage:       "using first char upper",
				Aliases:     []string{"uf"},
				DefaultText: "false",
			},
			&cli.StringFlag{
				Name:        "prefix",
				Usage:       "using prefix",
				Aliases:     []string{"p"},
				DefaultText: "",
			},
			&cli.StringSliceFlag{
				Name:    "cache",
				Usage:   `using cache (eg: -c="user_id" -c="UserID")`,
				Aliases: []string{"c"},
			},
			&cli.StringFlag{
				Name:    "style",
				Usage:   "using code style (eg: go)",
				Aliases: []string{"s"},
			},
		},
		// strnaming c -d="._"  -j='{"abv.bbc":123, "as_ds":123, "user_id": "1233", "sub":{"tu_bi":123}}' -c="user_id" -c="UserID" -f=./testdata/test.json
		Action: func(c *cli.Context) error {
			delimiter := strings.TrimSpace(c.String("delimiter"))
			prefix := strings.TrimSpace(c.String("prefix"))
			upperFirst := c.Bool("upperFirst")
			cacheSlice := c.StringSlice("cache")
			jsonObj := strings.TrimSpace(c.String("json"))
			filePath := strings.TrimSpace(c.String("file"))
			// style := strings.TrimSpace(c.String("style"))

			if len(cacheSlice)%2 != 0 {
				return fmt.Errorf(`cache key value pairs should appear in pairs, eg: -c="user_id" -c="UserID"`)
			}
			camel := strnaming.NewCamel().
				WithUpperFirst(upperFirst).
				WithPrefix(prefix).
				WithDelimiter([]byte(delimiter)...)
				// WithStyle(codes.Golang)

			for i, cl := 0, len(cacheSlice); i < cl; i += 2 {
				camel.WithCache(cacheSlice[i], cacheSlice[i+1])
			}

			if filePath != "" {
				if _, err := os.Stat(filePath); err != nil {
					return err
				}

				file, err := ioutil.ReadFile(filePath)
				if err != nil {
					return err
				}

				res, err := convert(camel, string(file))
				if err != nil {
					return err
				}
				fmt.Println(string(res))
			}

			if jsonObj != "" {
				res, err := convert(camel, jsonObj)
				if err != nil {
					return err
				}
				fmt.Println(string(res))
			}
			return nil
		},
	},
	{
		Name:    "snake",
		Aliases: []string{"s"},
		Usage:   "convert to snake string",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:    "file",
				Usage:   "input a json file path (eg: /path/to/strnaming.json)",
				Aliases: []string{"f"},
			},
			&cli.StringFlag{
				Name:    "json",
				Usage:   "input a json",
				Aliases: []string{"j"},
			},
			&cli.BoolFlag{
				Name:        "screaming",
				Usage:       "using screaming",
				DefaultText: "false",
				Aliases:     []string{"s"},
			},
			&cli.StringFlag{
				Name:    "ignores",
				Usage:   "using ignore chars",
				Aliases: []string{"i"},
			},
			&cli.StringSliceFlag{
				Name:    "cache",
				Usage:   "using cache",
				Aliases: []string{"c"},
			},
		},
		Action: func(c *cli.Context) error {
			screaming := c.Bool("screaming")
			ignores := strings.TrimSpace(c.String("ignore"))
			prefix := strings.TrimSpace(c.String("prefix"))
			cacheSlice := c.StringSlice("cache")
			jsonObj := strings.TrimSpace(c.String("json"))
			filePath := strings.TrimSpace(c.String("file"))

			if len(cacheSlice)%2 != 0 {
				return fmt.Errorf(`cache key value pairs should appear in pairs, eg: -c="user_id" -c="UserID"`)
			}
			snake := strnaming.NewSnake().
				WithScreaming(screaming).
				WithPrefix(prefix).
				WithIgnore([]byte(ignores)...)

			for i, cl := 0, len(cacheSlice); i < cl; i += 2 {
				snake.WithCache(cacheSlice[i], cacheSlice[i+1])
			}

			if filePath != "" {
				if _, err := os.Stat(filePath); err != nil {
					return err
				}

				file, err := ioutil.ReadFile(filePath)
				if err != nil {
					return err
				}

				res, err := convert(snake, string(file))
				if err != nil {
					return err
				}
				fmt.Println(string(res))
			}

			if jsonObj != "" {
				res, err := convert(snake, jsonObj)
				if err != nil {
					return err
				}
				fmt.Println(string(res))
			}
			return nil
		},
	},
	{
		Name:    "kebab",
		Aliases: []string{"k"},
		Usage:   "convert to kabab string",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:    "file",
				Usage:   "input a json file path (eg: /path/to/strnaming.json)",
				Aliases: []string{"f"},
			},
			&cli.StringFlag{
				Name:    "json",
				Usage:   "input a json",
				Aliases: []string{"j"},
			},
			&cli.BoolFlag{
				Name:        "screaming",
				Usage:       "using screaming",
				DefaultText: "false",
				Aliases:     []string{"s"},
			},
			&cli.StringFlag{
				Name:    "ignore",
				Usage:   "using ignore char",
				Aliases: []string{"i"},
			},
			&cli.StringSliceFlag{
				Name:    "cache",
				Usage:   "using cache",
				Aliases: []string{"c"},
			},
		},
		Action: func(c *cli.Context) error {
			screaming := c.Bool("screaming")
			ignores := strings.TrimSpace(c.String("ignore"))
			prefix := strings.TrimSpace(c.String("prefix"))
			cacheSlice := c.StringSlice("cache")
			jsonObj := strings.TrimSpace(c.String("json"))
			filePath := strings.TrimSpace(c.String("file"))

			if len(cacheSlice)%2 != 0 {
				return fmt.Errorf(`cache key value pairs should appear in pairs, eg: -c="user_id" -c="UserID"`)
			}
			kebab := strnaming.NewKebab().
				WithScreaming(screaming).
				WithPrefix(prefix).
				WithIgnore([]byte(ignores)...)

			for i, cl := 0, len(cacheSlice); i < cl; i += 2 {
				kebab.WithCache(cacheSlice[i], cacheSlice[i+1])
			}

			if filePath != "" {
				if _, err := os.Stat(filePath); err != nil {
					return err
				}

				file, err := ioutil.ReadFile(filePath)
				if err != nil {
					return err
				}

				res, err := convert(kebab, string(file))
				if err != nil {
					return err
				}
				fmt.Println(string(res))
			}

			if jsonObj != "" {
				res, err := convert(kebab, jsonObj)
				if err != nil {
					return err
				}
				fmt.Println(string(res))
			}
			return nil
		},
	},
}

func main() {
	app := cli.NewApp()
	app.Usage = "a cli tool to convert string name"
	app.Version = fmt.Sprintf("%s %s/%s", "v0.4.0", runtime.GOOS, runtime.GOARCH)
	app.Commands = commands

	// cli already print error messages
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
