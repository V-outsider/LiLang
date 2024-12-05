package main

import (
	"github.com/antlr4-go/antlr/v4"

	"fmt"
	"os"

	"github.com/V-outsider/LiLang/compiler"
	generated "github.com/V-outsider/LiLang/generated"
)

func main() {
	content, err := antlr.NewFileStream("test/test.lil")
	if err != nil {
		panic("Error while reading file")
	}

	lexer := generated.NewLiLangLexer(content)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := generated.NewLiLangParser(tokens)
	tree := parser.Program()

	// Create and use the IR listener
	irListener := compiler.NewIRListener()
	antlr.ParseTreeWalkerDefault.Walk(irListener, tree)
	file, err := os.Create("output.ll")
	if err != nil {
		panic("Error creating output file")
	}
	defer file.Close()

	_, err = irListener.State.Module.WriteTo(file)
	if err != nil {
		panic("Error writing LLVM IR to file")
	}

	fmt.Println("LLVM IR code has been written to output.ll")
}
