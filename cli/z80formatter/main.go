package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/jeromelesaux/z80format"
)

var (
	format  = flag.String("format", "", "assembly file to format")
	help    = flag.Bool("help", false, "display help message")
	rasm    = flag.Bool("rasm", false, "enable rasm syntaxe substitution")
	version = "0.1.1"
)

func main() {
	flag.Parse()
	var in *os.File
	if *help {
		fmt.Println("Version: " + version)
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
	// in.Seek(0, io.SeekStart)
	// s, _ := ioutil.ReadAll(in)
	// if compare(string(s), result) {
	// 	fmt.Fprintf(os.Stderr, "this convertion may not result a good convertion.\n")
	// }
	if *rasm {
		result = z80format.RasmSyntaxe(result)
	}
	fmt.Printf("%s", result)

	os.Exit(0)
}

func compare(s0, s1 string) bool {
	replacer := strings.NewReplacer(" ", "", "\t", "", "\r", "", "\n", "")
	s0Replaced := strings.ToUpper(replacer.Replace(s0))
	s1Replaced := strings.ToUpper(replacer.Replace(s1))
	md0 := md5.Sum([]byte(s0Replaced))
	md1 := md5.Sum([]byte(s1Replaced))
	fmt.Fprintf(os.Stderr, "%x %x\n", md0, md1)
	return reflect.DeepEqual(md0, md1)
}
