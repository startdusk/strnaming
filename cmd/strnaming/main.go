package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/startdusk/strnaming"
	"github.com/urfave/cli/v2"
)

var screen = bufio.NewWriter(os.Stdout)

var commands = []*cli.Command{
	{
		Name:    "camelcase",
		Aliases: []string{"c"},
		Usage:   "convert to camelcase string",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:    "file",
				Usage:   "input a file path",
				Aliases: []string{"f"},
			},
			&cli.StringFlag{
				Name:    "object",
				Usage:   "input a object",
				Aliases: []string{"obj"},
			},
			&cli.StringFlag{
				Name:        "delimiter",
				Usage:       "using custom delimiter",
				Aliases:     []string{"d"},
				DefaultText: "_",
			},
			&cli.BoolFlag{
				Name:        "lower-first",
				Usage:       "using first char lower",
				Aliases:     []string{"lf"},
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
				Usage:   "using cache",
				Aliases: []string{"c"},
			},
		},
		Action: func(c *cli.Context) error {
			delimiter := strings.TrimSpace(c.String("delimiter"))
			prefix := strings.TrimSpace(c.String("prefix"))
			lowerFirst := c.Bool("lower-first")
			cacheSlice := c.StringSlice("cache")
			// filePath := strings.TrimSpace(c.Path("file"))
			obj := strings.TrimSpace(c.String("object"))
			if len(cacheSlice)%2 != 0 {
				return fmt.Errorf("cache should be a pair")
			}
			camel := strnaming.NewCamel().
				WithLowerFirst(lowerFirst).
				WithPrefix(prefix)

			if delimiter != "" {
				camel.WithDelimiter(delimiter[0])
			}
			for i, cl := 0, len(cacheSlice); i < cl; i += 2 {
				camel.WithCache(cacheSlice[i], cacheSlice[i+1])
			}
			if obj != "" {
				srcObj := make(map[string]interface{})
				err := json.Unmarshal([]byte(obj), &srcObj)
				if err != nil {
					return err
				}

				destObj := make(map[string]interface{})
				for k, v := range srcObj {
					destObj[camel.Convert(k)] = v
				}

				bytes, err := json.MarshalIndent(destObj, "", "    ")
				if err != nil {
					return err
				}
				fmt.Fprintln(screen, string(bytes))
			}
			return nil
		},
	},
	{
		Name:    "snake",
		Aliases: []string{"s"},
		Usage:   "convert to snake string",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "ignore",
				Usage:       "using custom delimiter",
				Aliases:     []string{"d"},
				DefaultText: "_",
			},
		},
		Action: func(c *cli.Context) error {
			delimiter := c.String("delimiter")
			fmt.Println("get key", delimiter)
			return nil
		},
	},
	{
		Name:    "kebab",
		Aliases: []string{"k"},
		Usage:   "convert to kabab string",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "delimiter",
				Usage:       "using custom delimiter",
				Aliases:     []string{"d"},
				DefaultText: "_",
			},
		},
		Action: func(c *cli.Context) error {
			delimiter := c.String("delimiter")
			fmt.Println("get key", delimiter)
			return nil
		},
	},
}

func main() {
	app := cli.NewApp()
	app.Usage = "a cli tool to convert string name"
	app.Version = fmt.Sprintf("%s %s/%s", "v0.2.0", runtime.GOOS, runtime.GOARCH)
	app.Commands = commands

	defer screen.Flush()

	// cli already print error messages
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(screen, err)
		os.Exit(1)
	}
}
