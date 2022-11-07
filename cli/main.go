package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jeromelesaux/z80format"
)

var (
	format = flag.String("format", "", "assembly file to format")
	help   = flag.Bool("help", false, "display help message")
	rasm   = flag.Bool("rasm", true, "enable rasm syntaxe substitution")
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

	result, _ := z80format.Format(in)
	if *rasm {
		result = z80format.RasmSyntaxe(result)
	}
	fmt.Printf("%s", result)
	os.Exit(0)
}
