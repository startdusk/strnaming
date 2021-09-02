package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/startdusk/strnaming"
)

var (
	typ = flag.String("-c", "camel", "type to other string name")
)

func main() {
	flag.Parse()
	args := os.Args
	if len(args) < 2 {
		return
	}
	screen := bufio.NewWriter(os.Stdout)
	defer screen.Flush()

	switch *typ {
	case "camel":
		{
			camel := strnaming.NewCamel()
			for _, v := range args[1:] {
				fmt.Fprintln(screen, camel.Do(v))
			}
		}
	}
}
