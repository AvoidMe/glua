package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AvoidMe/glua/pkg/compiler"
	"github.com/AvoidMe/glua/pkg/interpreter"
	"github.com/AvoidMe/glua/pkg/lexer"
	"github.com/AvoidMe/glua/pkg/parser"
)

var (
	filename = flag.String("file", "", "File to run")
)

func main() {
	flag.Parse()
	if *filename == "" {
		fmt.Println("Need -file to run")
		return
	}
	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lexer := lexer.New()
	err = lexer.Consume(file, *filename)
	if err != nil {
		panic(err)
	}
	parser := parser.New(lexer.Tokens)
	ast, err := parser.Consume()
	if err != nil {
		panic(err)
	}
	compiler := compiler.New()
	bytecode, err := compiler.Compile(ast)
	if err != nil {
		panic(err)
	}
	interpreter := interpreter.New()
	err = interpreter.Run(bytecode)
	if err != nil {
		panic(err)
	}
}
