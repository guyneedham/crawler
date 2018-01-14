package main

import (
	"crawler/sub"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("usage: ./crawler start-page output-file")
		os.Exit(1)
	}
	crawler.Run(args[0], args[1])
}
