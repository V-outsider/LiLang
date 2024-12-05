package main

import (
	"github.com/V-outsider/LiLang/analyzers"
	"github.com/antlr4-go/antlr/v4"

	"github.com/V-outsider/LiLang/compiler"
	generated "github.com/V-outsider/LiLang/generated"
)

func main() {
	content, err := antlr.NewFileStream("test/test.lil")
	if err != nil {
		panic("Well, sucks while getting file {filename}")
	}

	lexer := generated.NewLiLangLexer(content)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := generated.NewLiLangParser(tokens)

	tree := parser.Program()

	analyzer := analyzers.NewStaticAnalyzer()

	antlr.ParseTreeWalkerDefault.Walk(analyzer, tree)

	visitor := &compiler.LiLangVisitor{}
	// visitor.
	// fmt.Println(tree.ToStringTree(parser.LiteralNames, parser))
}
