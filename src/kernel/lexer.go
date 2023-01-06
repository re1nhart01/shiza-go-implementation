package kernel

import (
	"regexp"
	"strings"

	"shiza.io/src"
)

type Lexer struct {
	Code      []string
	Pos       int64
	Line      int64
	TokenList []Token
}

func NewLexer(codeContent string) *Lexer {
	return &Lexer{
		TokenList: []Token{},
		Code:      strings.Split(codeContent, "\n"),
		Pos:       0,
		Line:      0,
	}
}

func (lexer *Lexer) RunLexicalAnalysis() {
	for lexer.getNextToken() {
	}
}

func (lexer *Lexer) getNextToken() bool {
	if lexer.Line >= int64(len(lexer.Code)) || len(lexer.Code[lexer.Line]) <= 0 {
		return false
	}
	for _, tokenType := range TokenQuerySlice {
		regexp := regexp.MustCompile("^" + tokenType.Reg)
		result := regexp.FindAllString(lexer.Code[int(lexer.Line)][int(lexer.Pos):], -1)
		if len(result) > 0 && len(result[0]) > 0 {
			if !tokenType.IsNotRequired {
				token := NewToken(tokenType, result[0], lexer.Pos, lexer.Line)
				lexer.TokenList = append(lexer.TokenList, token)
			}
			lexer.Pos += int64(len(result[0]))
			if lexer.Pos >= int64(len(lexer.Code[int(lexer.Line)])) {
				lexer.Line++
				lexer.Pos = 0
			}
			return true
		}
	}
	src.LogLexerError(lexer.Pos, lexer.Line+1, lexer.Code[lexer.Line], true)
	return false
}
