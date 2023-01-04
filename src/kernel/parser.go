package kernel

import (
	"shiza.io/src"
	"shiza.io/src/kernel/ast_tree"
)

type Parser struct {
	Tokens []Token
	Pos    int64
	Scope  map[string]string
}

func NewParser(tokenList []Token) Parser {
	return Parser{
		Tokens: tokenList,
		Pos:    0,
		Scope:  map[string]string{},
	}
}

func (parser *Parser) match(tokens []TokenQuery) *Token {
	if parser.Pos < int64(len(tokens)) {
		currentToken := parser.Tokens[parser.Pos]
		for _, v := range tokens {
			if v.Name == currentToken.Type.Name {
				parser.Pos += 1
				return &currentToken
			}
		}
	}
	return nil
}

func (parser *Parser) checkIsExpected(tokens ...TokenQuery) *Token {
	token := parser.match(tokens)
	if token == nil {
		src.LogParseError(parser.Pos, tokens[0].Name, true)
	}
	return token
}

// func (parser *Parser) RunParseCode() ast_tree.ExpressionNode {
// 	root := ast_tree.StatementsNode
// 	for parser.Pos < len(parser.Tokens) {
// 		codeStringNode := parser.parseExpression()
// 	}
// 	return root
// }

func (parser *Parser) parseExpression() ast_tree.ExpressionNode {
	return ast_tree.ExpressionNode{}
}

func get(item TokenQuery) []TokenQuery {
	return []TokenQuery{item}
}
