package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jeromelesaux/z80format"
)

var (
	format = flag.String("format", "", "assembly file to format")
	help   = flag.Bool("help", false, "display help message")
)

func main() {
	flag.Parse()
	var in *os.File
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	if *format == "" {
		in = os.Stdin
	} else {
		var err error
		in, err = os.Open(*format)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while reading file (%s) error [%s]\n", *format, err.Error())
			os.Exit(-1)
		}
		defer in.Close()
	}

	content, err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while reading content file (%s) error [%s]\n", *format, err.Error())
		os.Exit(-1)
	}
	result, _ := z80format.Format(string(content))
	fmt.Printf("%s", result)
	os.Exit(0)
}
