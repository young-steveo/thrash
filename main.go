package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/young-steveo/thrash/debugger"
	"github.com/young-steveo/thrash/lexer"
	"github.com/young-steveo/thrash/result"
)

var f string
var r bool
var d bool

func main() {

	flag.StringVar(&f, `f`, `main.rock`, `rockstar script to execute`)
	flag.BoolVar(&r, `r`, false, `start a repl`)
	flag.BoolVar(&d, `d`, false, `debug mode`)
	flag.Parse()

	if r {
		repl()
	} else {
		file(f)
	}
	os.Exit(int(result.ExOK))
}

func repl() {
	// todo
}

func file(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("Source file does not exist: %s\n", path)
		os.Exit(int(result.ExNoInput))
	}
	source, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Source file could not be read: %s\n", path)
		os.Exit(int(result.ExNoInput))
	}
	tokens := lexer.Scan(lexer.FromBytes(source))
	if d {
		debugger.PrintTokens(tokens)
	}
	os.Exit(int(0))
}
