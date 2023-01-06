package main

import (
	"fmt"
	"sync"

	"shiza.io/src"
	"shiza.io/src/kernel"
	"shiza.io/src/kernel/compilation"
)

var wg sync.WaitGroup

func main() {
	fs := src.NewFS()
	content := fs.GetFileContent("./src/main.bmx")
	lexer := kernel.NewLexer(content)
	lexer.RunLexicalAnalysis()
	tokenList := lexer.TokenList
	parser := kernel.NewParser(tokenList)
	compilator := compilation.NewPompilator()

	rootNode := parser.RunParseCode()
	wg.Add(1)
	go compilator.TryToCompile(&wg, &rootNode)
	wg.Wait()

	defer fmt.Println(lexer.TokenList)
}
