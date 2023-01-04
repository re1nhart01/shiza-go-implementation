package main

import (
	"fmt"

	"shiza.io/src"
	"shiza.io/src/kernel"
)

func main() {
	fs := src.NewFS()
	content := fs.GetFileContent("./src/main.bmx")
	lexer := kernel.NewLexer(content)
	lexer.RunLexicalAnalysis()
	defer fmt.Println(lexer.TokenList)
}
